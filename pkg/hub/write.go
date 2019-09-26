package hub

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/lomik/nooLiteHub/pkg/mtrf"
)

type writeContext struct {
	ch      uint8
	d0      uint8
	d1      uint8
	d2      uint8
	d3      uint8
	payload string
}

func (h *Hub) init() {
	h.writeRouter.AddParam("ch", func(value string, ctx interface{}) error {
		i, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		if i < 0 || i > 63 {
			return fmt.Errorf("ch value %d out of range [0, 63]", i)
		}
		ctx.(*writeContext).ch = uint8(i)
		return nil
	})

	h.writeRouter.AddParam("device", func(value string, ctx interface{}) error {
		if len(value) != 8 {
			return fmt.Errorf("invalid length of device id, expected 8")
		}

		v, err := strconv.ParseInt(value, 16, 64)
		if err != nil {
			return err
		}

		ctx.(*writeContext).d0 = uint8((v >> 24) % 256)
		ctx.(*writeContext).d1 = uint8((v >> 16) % 256)
		ctx.(*writeContext).d2 = uint8((v >> 8) % 256)
		ctx.(*writeContext).d3 = uint8(v % 256)
		return nil
	})

	h.write("raw", func(ctx *writeContext) {
		r, err := mtrf.JSONRequest([]byte(ctx.payload))
		if err != nil {
			h.onError(err)
			return
		}
		h.sendRequest(r)
	})

	// TX topics
	h.write("tx/:ch/on", func(ctx *writeContext) {
		h.sendRequest(&mtrf.Request{Mode: mtrf.ModeTX, Ch: ctx.ch, Cmd: mtrf.CmdOn})
	})

	h.write("tx/:ch/off", func(ctx *writeContext) {
		h.sendRequest(&mtrf.Request{Mode: mtrf.ModeTX, Ch: ctx.ch, Cmd: mtrf.CmdOff})
	})

	h.write("tx/:ch/switch", func(ctx *writeContext) {
		h.sendRequest(&mtrf.Request{Mode: mtrf.ModeTX, Ch: ctx.ch, Cmd: mtrf.CmdSwitch})
	})

	h.write("tx/:ch/bind", func(ctx *writeContext) {
		h.sendRequest(&mtrf.Request{Mode: mtrf.ModeTX, Ch: ctx.ch, Cmd: mtrf.CmdBind})
	})

	// TX-F topics
	h.write("tx-f/:ch/on", func(ctx *writeContext) {
		h.sendRequest(&mtrf.Request{Mode: mtrf.ModeTXF, Ch: ctx.ch, Cmd: mtrf.CmdOn})
	})

	h.write("tx-f/:ch/off", func(ctx *writeContext) {
		h.sendRequest(&mtrf.Request{Mode: mtrf.ModeTXF, Ch: ctx.ch, Cmd: mtrf.CmdOff})
	})

	h.write("tx-f/:ch/switch", func(ctx *writeContext) {
		h.sendRequest(&mtrf.Request{Mode: mtrf.ModeTXF, Ch: ctx.ch, Cmd: mtrf.CmdSwitch})
	})

	h.write("tx-f/:ch/bind", func(ctx *writeContext) {
		h.sendRequest(&mtrf.Request{Mode: mtrf.ModeTXF, Ch: ctx.ch, Cmd: mtrf.CmdBind})
	})

	h.write("tx-f/:ch/read_state", func(ctx *writeContext) {
		h.sendRequest(&mtrf.Request{Mode: mtrf.ModeTXF, Ch: ctx.ch, Cmd: mtrf.CmdReadState})
	})

	h.write("tx-f/:ch/:device/read_state", func(ctx *writeContext) {
		h.sendRequest(&mtrf.Request{Mode: mtrf.ModeTXF, Ch: ctx.ch, D0: ctx.d0, D1: ctx.d1, D2: ctx.d2, D3: ctx.d3, Cmd: mtrf.CmdReadState})
	})
}

// регистрирует callback на входящее mqtt сообщение
func (h *Hub) write(path string, callback func(ctx *writeContext)) {
	h.writeRouter.AddPath(path, func(ctx interface{}) {
		callback(ctx.(*writeContext))
	})
}

// ждет новые события из mqtt
func (h *Hub) mqttWorker() {
	for m := range h.mqttClient.Incoming {
		b := new(bytes.Buffer)
		m.Payload.WritePayload(b)
		log.Printf("[mqtt] -> %s: %s", m.TopicName, b.String())

		topicName := m.TopicName
		topicName = strings.TrimPrefix(topicName, h.options.Topic+"/write/")

		ctx := &writeContext{payload: b.String()}
		if err := h.writeRouter.Route(topicName, ctx); err != nil {
			log.Println(err)
		}
	}
}
