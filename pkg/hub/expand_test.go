package hub

import (
	"testing"

	"github.com/lomik/nooLiteHub/pkg/mtrf"
	"github.com/stretchr/testify/assert"
)

func TestExpandResponse(t *testing.T) {
	assert := assert.New(t)

	table := map[string](map[string]string){
		"[173,2,0,0,7,130,0,2,0,1,255,0,0,203,182,187,174]": {
			"tx-f/7/0000CBB6/state/bind":       "off",
			"tx-f/7/0000CBB6/state/brightness": "255",
			"tx-f/7/0000CBB6/state/power":      "on",
		},
	}

	for body, expected := range table {
		response, err := mtrf.JSONResponse([]byte(body))
		assert.NoError(err)
		assert.Equal(expected, expandResponse(response))
	}
}
