package gh

import (
	"errors"
	"os"
	"reflect"

	t "github.com/arnoluo/gh/types"
)

var (
	Str   t.StrType
	Int   t.IntType
	Bool  t.BoolType
	Float t.FloatType
	Itf   t.ItfType
)

// 获取环境变量
func Env(name t.S) t.S {
	return os.Getenv(name)
}

// 获取环境变量，返回int, 发生错误时返回指定默认值
func EnvInt(name t.S, defaultValue t.I) t.I {
	return Str.Int(Env(name), defaultValue)
}

func Int64(i t.Itf, defaultValue t.LI) t.LI {
	switch i.(type) {
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64:
		return i.(t.LI)
	case string:
		return t.LI(Str.Int(i.(string), -1))
	default:
		return defaultValue
	}
}

func IntValue(itf t.Itf, defaultValue t.I) t.I {
	switch itf.(type) {
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64:
		return itf.(t.I)
	case string:
		return Str.Int(itf.(string), -1)
	default:
		return defaultValue
	}
}

func Uint64(i t.Itf) (uint64, error) {
	var v t.LI
	switch i.(type) {
	case uint, uint8, uint16, uint32, int, int8, int16, int32, int64, float32, float64:
		v = i.(t.LI)
	case uint64:
		return i.(uint64), nil
	case string:
		v = t.LI(Str.Int(i.(string), -1))
	default:
	}

	var err error
	if v < 0 {
		err = errors.New("wrong value")
	}

	return uint64(v), err
}

func Float64(i t.Itf, defaultValue t.F) t.F {
	switch i.(type) {
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64, float32, float64:
		return i.(t.F)
	case string:
		return t.F(Str.Int(i.(string), -1))
	default:
		return defaultValue
	}
}

// 结构同名称转换
func StructTo(src, dst t.Itf) error {
	srcType := reflect.TypeOf(src)
	srcValue := reflect.ValueOf(src)
	dstValue := reflect.ValueOf(dst)

	if srcType.Kind() != reflect.Struct {
		return errors.New("src must be struct")
	}
	if dstValue.Kind() != reflect.Ptr {
		return errors.New("dst must be pointer")
	}

	for i := 0; i < srcType.NumField(); i++ {
		dstField := dstValue.Elem().FieldByName(srcType.Field(i).Name)
		if dstField.CanSet() {
			dstField.Set(srcValue.Field(i))
		}
	}

	return nil
}

// Return true if stack has the element item, return false otherwise
func InArray(item, stack t.Itf) t.B {
	arrType := reflect.TypeOf(stack)
	kd := arrType.Kind()
	if kd == reflect.Slice || kd == reflect.Array {
		v := reflect.ValueOf(stack)
		for i := 0; i < v.Len(); i++ {
			if v.Index(i).Interface() == item {
				return true
			}
		}
	}

	return false
}
