package t

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItfIf(t *testing.T) {
	var itf ItfType
	assert.Equal(t, 1, itf.If(true, 1, 0))
	assert.Equal(t, 0, itf.If(false, 1, 0))
	assert.Equal(t, "trueValue", itf.If(true, "trueValue", "falseValue"))
	assert.Equal(t, "falseValue", itf.If(false, "trueValue", "falseValue"))
}
