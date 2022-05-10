package gh

import (
	t "gh/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIf(t *testing.T) {
	assert.Equal(t, "trueValue", Str.If(true, "trueValue", "falseValue"))
	assert.Equal(t, "falseValue", Str.If(false, "trueValue", "falseValue"))
	assert.Equal(t, 1, Int.If(true, 1, 0))
	assert.Equal(t, 0, Int.If(false, 1, 0))
	assert.Equal(t, true, Bool.If(true, true, false))
	assert.Equal(t, false, Bool.If(false, true, false))
	assert.Equal(t, 1.11, Float.If(true, 1.11, 0.00))
	assert.Equal(t, 0.00, Float.If(false, 1.11, 0.00))
	assert.Equal(t, 1, Itf.If(true, 1, 0))
	assert.Equal(t, 0, Itf.If(false, 1, 0))
}

// response type
type Rsp struct {
	Code t.I `json:"code"`
}

// response type with ts(timestamp)
type Rspt struct {
	Code t.I  `json:"code"`
	Ts   t.LI `json:"ts"`
}

func TestStructTo(t *testing.T) {
	jsonRspt := Rspt{
		Code: 1,
		Ts:   123,
	}

	var jsonRsp Rsp
	StructTo(jsonRspt, &jsonRsp)
	assert.Equal(t, jsonRsp, Rsp{
		Code: 1,
	})

	var jRspt Rspt
	StructTo(jsonRsp, &jRspt)
	assert.Equal(t, jRspt, Rspt{
		Code: 1,
		Ts:   0,
	})
}

func TestInArray(t *testing.T) {
	assert.True(t, InArray(1, []int{1, 2, 3}))
	var a float64 = 1
	assert.True(t, InArray(a, []float64{1, 2}))
	// not recommanded
	assert.False(t, InArray(1, []float64{1, 2}))
	assert.True(t, InArray("a", []string{"b", "a"}))
	assert.True(t, InArray("1", []string{"1", "2"}))
	assert.False(t, InArray("1", []int{1, 2}))
}
