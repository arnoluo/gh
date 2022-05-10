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
