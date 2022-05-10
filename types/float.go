package t

func (f *FloatType) If(isTrue B, trueValue, falseValue F) F {
	if isTrue {
		return trueValue
	}

	return falseValue
}
