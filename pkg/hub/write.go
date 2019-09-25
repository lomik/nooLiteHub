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
	mode    uint8
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

	modes := map[string]uint8{
		"tx":   mtrf.ModeTX,
		"rx":   mtrf.ModeRX,
		"tx-f": mtrf.ModeTXF,
		"rx-f": mtrf.ModeRXF,
	}

	h.writeRouter.AddParam("mode", func(value string, ctx interface{}) error {
		m, ok := modes[value]
		if !ok {
			return fmt.Errorf("unknown mode %#v", value)
		}
		ctx.(*writeContext).mode = m
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

	h.write(":mode/:ch/on", func(ctx *writeContext) {
		h.sendRequest(&mtrf.Request{Mode: ctx.mode, Ch: ctx.ch, Cmd: mtrf.CmdOn})
	})

	h.write(":mode/:ch/off", func(ctx *writeContext) {
		h.sendRequest(&mtrf.Request{Mode: ctx.mode, Ch: ctx.ch, Cmd: mtrf.CmdOff})
	})

	h.write(":mode/:ch/switch", func(ctx *writeContext) {
		h.sendRequest(&mtrf.Request{Mode: ctx.mode, Ch: ctx.ch, Cmd: mtrf.CmdSwitch})
	})

	h.write(":mode/:ch/bind", func(ctx *writeContext) {
		h.sendRequest(&mtrf.Request{Mode: ctx.mode, Ch: ctx.ch, Cmd: mtrf.CmdBind})
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
