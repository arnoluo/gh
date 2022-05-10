package t

func (b *BoolType) If(isTrue, trueValue, falseValue B) B {
	if isTrue {
		return trueValue
	}

	return falseValue
}
