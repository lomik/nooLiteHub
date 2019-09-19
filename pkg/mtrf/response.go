package mtrf

import "fmt"

// NewResponse создает инстанс Response из байт, полученных от модуля
func NewResponse(b []byte) (*Response, error) {
	if len(b) != MessageLen {
		return nil, fmt.Errorf("wrong message length %d", len(b))
	}
	r := &Response{
		St: b[0], Mode: b[1], Ctr: b[2], Togl: b[3],
		Ch: b[4], Cmd: b[5], Fmt: b[6], D0: b[7],
		D1: b[8], D2: b[9], D3: b[10], ID0: b[11],
		ID1: b[12], ID2: b[13], ID3: b[14], Crc: b[15],
		Sp: b[16],
	}
	return r, nil
}

// MustResponse делает тоже самое что NewResponse. Но паникует вместо возврата ошибки
func MustResponse(b []byte) *Response {
	r, err := NewResponse(b)
	if err != nil {
		panic(err)
	}
	return r
}

func (r *Response) String() string {
	return fmt.Sprintf("[%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d]",
		r.St, r.Mode, r.Ctr, r.Togl,
		r.Ch, r.Cmd, r.Fmt, r.D0,
		r.D1, r.D2, r.D3, r.ID0,
		r.ID1, r.ID2, r.ID3, r.Crc,
		r.Sp,
	)
}
