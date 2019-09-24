package mtrf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonRequest(t *testing.T) {
	assert := assert.New(t)

	func() {
		r, err := JSONRequest([]byte("[171,1,0,26,44,2,0,0,0,0,0,0,0,0,0,246,172]"))
		assert.NoError(err)
		assert.Equal(&Request{St: 171, Mode: 1, Ctr: 0, Res: 26, Ch: 44, Cmd: 2, Fmt: 0, D0: 0, D1: 0, D2: 0, D3: 0, ID0: 0, ID1: 0, ID2: 0, ID3: 0, Crc: 246, Sp: 172}, r)
	}()

	badMessage := []string{
		"[171,1,26,44,2,0,0,0,0,0,0,0,0,0,246,172]",    // короткое
		"[171,1,0,-1,44,2,0,0,0,0,0,0,0,0,0,246,172]",  // отрицательное число
		"[171,1,0,256,44,2,0,0,0,0,0,0,0,0,0,246,172]", // большое число
	}

	for _, m := range badMessage {
		func() {
			r, err := JSONRequest([]byte(m))
			assert.Error(err)
			assert.Nil(r)
		}()
	}
}
