package mtrf

import (
	"encoding/json"
	"fmt"
)

// JSON возвращает сериализованное представление Request
func (r *Request) JSON() string {
	r.crc(0)
	return fmt.Sprintf("[%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d]",
		r.St, r.Mode, r.Ctr, r.Res,
		r.Ch, r.Cmd, r.Fmt, r.D0,
		r.D1, r.D2, r.D3, r.ID0,
		r.ID1, r.ID2, r.ID3, r.Crc,
		r.Sp,
	)
}

func (r *Request) String() string {
	r.crc()
	return fmt.Sprintf("Request{St: %d, Mode: %d, Ctr: %d, Res: %d, Ch: %d, Cmd: %d, Fmt: %d, D0: %d, D1: %d, D2: %d, D3: %d, ID0: %d, ID1: %d, ID2: %d, ID3: %d, Crc: %d, Sp: %d}",
		r.St, r.Mode, r.Ctr, r.Res,
		r.Ch, r.Cmd, r.Fmt, r.D0,
		r.D1, r.D2, r.D3, r.ID0,
		r.ID1, r.ID2, r.ID3, r.Crc,
		r.Sp,
	)
}

func (r *Request) crc() {
	b := r.Bytes()
	r.St = b[0]
	r.Crc = b[15]
	r.Sp = b[16]
}

// Bytes возращает массив байт для отправки модулю.
// Первый байт всегда 171, последний 172.
// CRC вычисляется автоматически
func (r *Request) Bytes() []byte {
	b := []byte{171, r.Mode, r.Ctr, r.Res,
		r.Ch, r.Cmd, r.Fmt, r.D0,
		r.D1, r.D2, r.D3, r.ID0,
		r.ID1, r.ID2, r.ID3, r.Crc,
		172}
	x := uint8(0)
	for i := 0; i < 15; i++ {
		x += b[i]
	}
	b[15] = x
	return b
}

func makeRequest(p []uint8) *Request {
	return &Request{
		St:   p[0],
		Mode: p[1], Ctr: p[2], Res: p[3],
		Ch: p[4], Cmd: p[5], Fmt: p[6], D0: p[7],
		D1: p[8], D2: p[9], D3: p[10], ID0: p[11],
		ID1: p[12], ID2: p[13], ID3: p[14], Crc: p[15],
		Sp: p[16],
	}
}

// JSONRequest парсит json-массив из 17 чисел
func JSONRequest(payload []byte) (*Request, error) {
	p := make([]byte, 17)

	err := json.Unmarshal(payload, &p)
	if err != nil {
		return nil, err
	}

	if len(p) != 17 {
		return nil, fmt.Errorf("message length != 17")
	}

	return makeRequest(p), nil
}
