package structreflect

import "reflect"

func Copy(dest, src interface{}) {
	v1 := reflect.ValueOf(dest)
	if v1.Kind() == reflect.Ptr {
		v1 = reflect.Indirect(v1)
	}
	v2 := reflect.ValueOf(src)
	if v2.Kind() == reflect.Ptr {
		v2 = reflect.Indirect(v2)
	}
	v1.Set(v2)
}
