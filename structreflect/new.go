package structreflect

import "reflect"

func NewSameStruct(s interface{}) interface{} {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = reflect.Indirect(v)
	}
	return reflect.New(v.Type()).Interface()
}
