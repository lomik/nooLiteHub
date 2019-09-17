package mtrf

import (
	"fmt"
	"strings"
)

type Message struct {
	body [17]byte
}

func (rq *Message) SendPack() []byte {
	rq.body[0] = 171
	rq.body[16] = 172
	x := uint8(0)
	for i := 0; i < 15; i++ {
		x += rq.body[i]
	}
	rq.body[15] = x
	return rq.body[:]
}

func (rs *Message) RecvUnpack(p []byte) error {
	if len(p) != 17 {
		return fmt.Errorf("invalid message length")
	}

	for i := 0; i < 17; i++ {
		rs.body[i] = p[i]
	}
	if rs.body[0] != 173 || rs.body[16] != 174 {
		return fmt.Errorf("invalid first or last byte")
	}

	x := uint8(0)
	for i := 0; i < 15; i++ {
		x += rs.body[i]
	}
	if rs.body[15] != x {
		return fmt.Errorf("invalid checksum")
	}

	return nil
}

func (rq *Message) copy() *Message {
	var b [17]byte
	copy(b[:], rq.body[:])
	return &Message{body: b}
}

func (rq *Message) Mode(m Mode) *Message {
	r := rq.copy()
	r.body[1] = uint8(m)
	return r
}

func Bind(channel uint8) *Message {
	return Raw([17]byte{0, 0, 0, 0, channel, 15, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

func Unbind(channel uint8) *Message {
	return Raw([17]byte{0, 0, 0, 0, channel, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

func PowerOn(channel uint8) *Message {
	return Raw([17]byte{0, 0, 0, 0, channel, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

func PowerOff(channel uint8) *Message {
	return Raw([17]byte{0, 0, 0, 0, channel, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

func PowerSwitch(channel uint8) *Message {
	return Raw([17]byte{0, 0, 0, 0, channel, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

func RxBind(channel uint8) *Message {
	return Raw([17]byte{0, 1, 3, 0, channel, 15, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

func RxClear(channel uint8) *Message {
	return Raw([17]byte{0, 1, 5, 0, channel, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

func Raw(p [17]byte) *Message {
	m := &Message{}
	copy(m.body[:], p[:])
	return m
}

func (m *Message) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "raw: %s", bodyString(m.body))
	return b.String()
}

func bodyString(p [17]byte) string {
	return fmt.Sprintf("{%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d}",
		p[0], p[1], p[2], p[3], p[4], p[5], p[6], p[7], p[8], p[9], p[10], p[11],
		p[12], p[13], p[14], p[15], p[16],
	)
}

func (m *Message) Channel() uint8 {
	return m.body[4]
}
