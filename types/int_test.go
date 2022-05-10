package t

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntIf(t *testing.T) {
	var i IntType
	assert.Equal(t, 1, i.If(true, 1, 0))
	assert.Equal(t, 0, i.If(false, 1, 0))
}

func TestIntStr(t *testing.T) {
	var i IntType
	assert.Equal(t, "1", i.Str(1))
}
