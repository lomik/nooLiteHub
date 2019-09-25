package hub

import (
	"fmt"

	"github.com/lomik/nooLiteHub/pkg/mtrf"
)

func (h *Hub) expandResponse(r *mtrf.Response) {
	if r.Ctr != 0 {
		return
	}

	switch r.Cmd {
	case mtrf.CmdSendState:
		fmt.Println(r.String(), r.Device())
	}
}
