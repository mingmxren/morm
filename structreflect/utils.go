package structreflect

import (
	"fmt"
	"reflect"
)

func AssertStruct(val reflect.Value, needCanAddr bool) (reflect.Value, error) {
	if val.Type().Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	if val.Type().Kind() != reflect.Struct {
		return val, fmt.Errorf("AssertStruct(needCanAddr:%v) val:%v is not Struct", needCanAddr, val)
	}

	if needCanAddr && !val.CanAddr() {
		return val, fmt.Errorf("AssertStruct(needCanAddr:%v) val:%v is not CanAddr", needCanAddr, val)

	}
	
	return val, nil
}

func AssertPtr(val reflect.Value) (reflect.Value, error) {
	if val.Type().Kind() != reflect.Ptr {
		return val, fmt.Errorf("AssertStruct val:%v is not Ptr", val)
	}
	return val, nil
}
