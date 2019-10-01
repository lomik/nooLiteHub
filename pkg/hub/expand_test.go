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
			"txf/7/0000CBB6/state/bind":       "off",
			"txf/7/0000CBB6/state/brightness": "255",
			"txf/7/0000CBB6/state/power":      "on",
		},
		"[173,1,0,16,42,21,7,205,32,48,255,0,0,0,0,32,174]": {
			"rx/42/sensor/temperature": "20.5",
			"rx/42/sensor/humidity":    "48",
			"rx/42/sensor/low_battery": "false",
			"rx/42/sensor/device":      "PT111",
		},
	}

	for body, expected := range table {
		response, err := mtrf.JSONResponse([]byte(body))
		assert.NoError(err)
		assert.Equal(expected, expandResponse(response))
	}
}
