package hub

import (
	"fmt"

	"github.com/lomik/nooLiteHub/pkg/mtrf"
)

var statePower = map[uint8]string{
	0: "off",
	1: "on",
	2: "temporary_on",
}

var stateOnOff = map[uint8]string{
	0: "off",
	1: "on",
}

func expandResponse(r *mtrf.Response) (topicPayload map[string]string) {
	topicPayload = make(map[string]string)

	if r.Ctr != 0 {
		return
	}

	switch r.Cmd {
	case mtrf.CmdSendState:
		switch r.Fmt {
		case 0:
			topicPayload[fmt.Sprintf("tx-f/%d/%s/state/power", r.Ch, r.Device())] = statePower[r.D2&0xf]
			topicPayload[fmt.Sprintf("tx-f/%d/%s/state/bind", r.Ch, r.Device())] = stateOnOff[r.D2>>7]
			topicPayload[fmt.Sprintf("tx-f/%d/%s/state/brightness", r.Ch, r.Device())] = fmt.Sprintf("%d", r.D3)
		case 1:
			topicPayload[fmt.Sprintf("tx-f/%d/%s/state/input", r.Ch, r.Device())] = stateOnOff[r.D2]
			topicPayload[fmt.Sprintf("tx-f/%d/%s/state/temporary_disable_noolite", r.Ch, r.Device())] = stateOnOff[(r.D3>>1)&0x1]
			topicPayload[fmt.Sprintf("tx-f/%d/%s/state/disable_noolite", r.Ch, r.Device())] = stateOnOff[r.D3&0x1]
		case 2:
			topicPayload[fmt.Sprintf("tx-f/%d/%s/state/noolite_free_slots", r.Ch, r.Device())] = fmt.Sprintf("%d", r.D2)
			topicPayload[fmt.Sprintf("tx-f/%d/%s/state/noolite_f_free_slots", r.Ch, r.Device())] = fmt.Sprintf("%d", r.D3)
		}
	}

	return
}
