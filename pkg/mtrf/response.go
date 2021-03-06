package mtrf

import (
	"encoding/json"
	"fmt"
)

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

	if b[0] != 173 {
		return nil, fmt.Errorf("wrong start byte")
	}

	if b[16] != 174 {
		return nil, fmt.Errorf("wrong end byte")
	}

	x := uint8(0)
	for i := 0; i < 15; i++ {
		x += b[i]
	}

	if x != b[15] {
		return nil, fmt.Errorf("wrong crc")
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

// JSON возвращает сериализованное представление Response
func (r *Response) JSON() string {
	return fmt.Sprintf("[%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d]",
		r.St, r.Mode, r.Ctr, r.Togl,
		r.Ch, r.Cmd, r.Fmt, r.D0,
		r.D1, r.D2, r.D3, r.ID0,
		r.ID1, r.ID2, r.ID3, r.Crc,
		r.Sp,
	)
}

func (r *Response) String() string {
	return fmt.Sprintf("Response{St: %d, Mode: %d, Ctr: %d, Togl: %d, Ch: %d, Cmd: %d, Fmt: %d, D0: %d, D1: %d, D2: %d, D3: %d, ID0: %d, ID1: %d, ID2: %d, ID3: %d, Crc: %d, Sp: %d}",
		r.St, r.Mode, r.Ctr, r.Togl,
		r.Ch, r.Cmd, r.Fmt, r.D0,
		r.D1, r.D2, r.D3, r.ID0,
		r.ID1, r.ID2, r.ID3, r.Crc,
		r.Sp,
	)
}

// Device returns device id with "base 16, upper-case, two characters per byte" encoding
func (r *Response) Device() string {
	return fmt.Sprintf("%X", []byte{r.ID0, r.ID1, r.ID2, r.ID3})
}

// JSONResponse парсит json-массив из 17 чисел
func JSONResponse(payload []byte) (*Response, error) {
	p := make([]byte, 17)

	err := json.Unmarshal(payload, &p)
	if err != nil {
		return nil, err
	}

	if len(p) != 17 {
		return nil, fmt.Errorf("message length != 17")
	}

	return NewResponse(p)
}
