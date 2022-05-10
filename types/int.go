package t

import "strconv"

func (i *IntType) Str(value I) S {
	return strconv.Itoa(value)
}

func (i *IntType) If(isTrue B, trueValue, falseValue I) I {
	if isTrue {
		return trueValue
	}

	return falseValue
}
