package t

func (itf *ItfType) If(isTrue B, trueValue, falseValue Itf) Itf {
	if isTrue {
		return trueValue
	}

	return falseValue
}
