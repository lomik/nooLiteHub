package router

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookup(t *testing.T) {
	assert := assert.New(t)
	r := New()

	addPath := func(path string) {
		r.AddPath(path, func(ctx interface{}) { *(ctx.(*string)) = path })
	}

	test := func(path string, expectedRoute string, expectedParams map[string]string) {
		e, m := r.lookup(r.root, strings.Split(path, "/"))
		if expectedRoute == "" {
			assert.Nil(e, path)
			assert.Nil(m, path)
		} else {
			var p string
			if assert.NotNil(e, path) {
				e.callback(&p)
				assert.Equal(expectedRoute, p, path)
				assert.Equal(expectedParams, m, path)
			}
		}
	}

	addPath("rx/:ch/on")
	addPath("raw")

	test("rx/42/on", "rx/:ch/on", map[string]string{"ch": "42"})
	test("rx/42/off", "", nil)
	test("raw", "raw", map[string]string{})
}
