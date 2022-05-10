package t

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var str StrType

func TestStrIf(t *testing.T) {
	assert.Equal(t, "trueValue", str.If(true, "trueValue", "falseValue"), "should be string: trueValue")
	assert.Equal(t, "falseValue", str.If(false, "trueValue", "falseValue"), "should be string: falseValue")
}

func TestSub(t *testing.T) {
	assert.Equal(t, str.Sub("abcdef", 0, 2), "ab")
	assert.Equal(t, str.Sub("测试中", 0, 2), "测试")
}

func TestRandom(t *testing.T) {
	a := str.Rand(6)
	b := str.RandLetters(6)
	c := str.RandNumbers(6)
	d := str.RandChars("c", 3)
	fmt.Println(a, b, c, d)
	assert.Len(t, a, 6)
	assert.Len(t, b, 6)
	assert.Len(t, c, 6)
	assert.Equal(t, "ccc", d)
}

func TestRegReplace(t *testing.T) {
	c := str.RegReplace("acacac", `a+`, "c")
	assert.Equal(t, "cccccc", c)
}

func TestConvert(t *testing.T) {
	var a, b, c, d string = "1", "-1", "a", "0"
	assert.Equal(t, true, str.Bool(a, false))
	assert.Equal(t, 1, str.Int(a, -1))
	assert.Equal(t, 1, str.Pint(b, 1))
	assert.Equal(t, 0, str.UInt(d, 1))
	assert.Equal(t, false, str.Empty(c))
	assert.Equal(t, false, str.IsInt(c))
	assert.Equal(t, true, str.IsInt(a))
	assert.Equal(t, true, str.IsNum(b))
	assert.Equal(t, false, str.IsNum(c))

}
