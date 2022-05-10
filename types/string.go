package t

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

const (
	// 字母 + 数字，移除 il o0
	CHARS S = `abcdefghjkmnpqrstuvwxyz123456789`

	// 纯字母，移除 il
	PURE_LETTER_CHARS S = `abcdefghjkmnopqrstuvwxyz`

	// 纯数字字符
	PURE_NUMBER_CHARS S = `0123456789`
)

// string 转 int, 需设置转换错误时的默认值
func (s *StrType) Int(str S, errValue I) I {
	value, err := strconv.Atoi(str)
	if err != nil {
		return errValue
	}
	return value
}

// string 转 bool, 需设置转换错误时的默认值
func (s *StrType) Bool(str S, errBool B) B {
	value, err := strconv.ParseBool(str)
	if err != nil {
		return errBool
	}
	return value
}

// Covert string to unsigned int using strconv.Atoi(), return errValue if err != nil or value < 0
func (s *StrType) UInt(str S, errValue I) I {
	value := s.Int(str, errValue)
	if value < 0 {
		return errValue
	}

	return value
}

// Covert string to positive int using strconv.Atoi(), return errValue if err != nil or value <= 0
func (s *StrType) Pint(str S, errValue I) I {
	value := s.Int(str, errValue)
	if value <= 0 {
		return errValue
	}

	return value
}

//
func (s *StrType) IsInt(str S) B {
	_, err := strconv.Atoi(str)
	return err == nil
}

func (s *StrType) IsNum(str S) B {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func (s *StrType) Empty(value S) B {
	return value == ""
}

// utf8(6 bytes at most) substring
func (s *StrType) Sub(str S, begin, length I) S {
	rs := []rune(str)
	lth := len(rs)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}

	end := begin + length

	if end > lth {
		end = lth
	}

	return string(rs[begin:end])
}

// Return string param trueValue if boolValue=true, return string param falseValue otherwise
func (s *StrType) If(boolValue B, trueValue, falseValue S) S {
	if boolValue {
		return trueValue
	}

	return falseValue
}

// Replace baseStr with regexpPattern to replacement
func (s *StrType) RegReplace(baseStr, regexpPattern, replacement S) S {
	reg := regexp.MustCompile(regexpPattern)
	return reg.ReplaceAllString(baseStr, replacement)
}

// Generate random str base on letter&number mixed chars
func (s *StrType) Rand(length I) S {
	return s.RandChars(CHARS, length)
}

// Generate random str base on letter chars
func (s *StrType) RandLetters(length I) S {
	return s.RandChars(PURE_LETTER_CHARS, length)
}

// Generate random str base on number chars
func (s *StrType) RandNumbers(length I) S {
	return s.RandChars(PURE_NUMBER_CHARS, length)
}

// Generate random str base on baseChars
func (s *StrType) RandChars(chars S, length I) (str S) {
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		randPos := rand.Intn(len(chars))
		str += chars[randPos : randPos+1]
	}
	return
}
