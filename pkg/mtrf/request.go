package mtrf

import "fmt"

func (r *Request) String() string {
	return fmt.Sprintf("[%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d]",
		r.St, r.Mode, r.Ctr, r.Res,
		r.Ch, r.Cmd, r.Fmt, r.D0,
		r.D1, r.D2, r.D3, r.ID0,
		r.ID1, r.ID2, r.ID3, r.Crc,
		r.Sp,
	)
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
