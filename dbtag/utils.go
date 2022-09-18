package dbtag

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/mingmxren/morm/structreflect"
)

func NameValuesWithIgnore(s interface{}, target []string) (names []string, values []interface{}, err error) {
	return NameValuesWithTargetOrIgnore(s, target, nil)
}

func NameValuesWithTarget(s interface{}, target []string) (names []string, values []interface{}, err error) {
	return NameValuesWithTargetOrIgnore(s, target, nil)
}

func NameValues(s interface{}) (names []string, values []interface{}, err error) {
	return NameValuesWithTargetOrIgnore(s, nil, nil)
}

func NameValuesWithTargetOrIgnore(s interface{}, target []string, ignore []string) (names []string,
	values []interface{}, err error) {
	if len(target) > 0 && len(ignore) > 0 {
		return nil, nil, errors.New("target and ignore both true")
	}

	val, err := structreflect.AssertStruct(reflect.ValueOf(s), false)
	if err != nil {
		return nil, nil, err
	}

	structDbTags, err := getDbTags(val.Type())
	if err != nil {
		return nil, nil, err
	}

	structFields, err := structreflect.CachedStructFields(val.Type())
	if err != nil {
		return nil, nil, err
	}

	if target == nil {
		ignores := make(map[string]bool)
		for _, f := range ignore {
			ignores[f] = true
		}
		for _, t := range structDbTags {
			if t == nil {
				continue
			}
			if _, ok := ignores[t.name]; !ok {
				target = append(target, t.name)
			}
		}
	}

	nameToValue := make(map[string]reflect.Value)
	for i, f := range structFields {
		if structDbTags[i] == nil {
			continue
		}
		nameToValue[structDbTags[i].name] = val.FieldByIndex(f.Index)
	}

	for _, f := range target {
		if v, ok := nameToValue[f]; !ok {
			return nil, nil, fmt.Errorf("field[%s] not found", f)
		} else {
			names = append(names, f)
			values = append(values, v.Interface())
		}
	}

	return names, values, nil
}

func NameValuesByFieldPtr(s interface{}, fields ...interface{}) (names []string, values []interface{}, err error) {
	val, err := structreflect.AssertStruct(reflect.ValueOf(s), true)
	if err != nil {
		return nil, nil, err
	}

	structDbTags, err := getDbTags(val.Type())
	if err != nil {
		return nil, nil, err
	}

	structFields, err := structreflect.CachedStructFields(val.Type())
	if err != nil {
		return nil, nil, err
	}

	valueToName := make(map[uintptr]string)
	for i, f := range structFields {
		if structDbTags[i] == nil {
			continue
		}
		valueToName[val.FieldByIndex(f.Index).Addr().Pointer()] = structDbTags[i].name
	}

	for _, f := range fields {
		fv, err := structreflect.AssertPtr(reflect.ValueOf(f))
		if err != nil {
			return nil, nil, err
		}
		if name, ok := valueToName[fv.Pointer()]; !ok {
			return nil, nil, fmt.Errorf("NamesByFieldPtr fv:%v not found", fv)
		} else {
			names = append(names, name)
		}
		values = append(values, reflect.Indirect(fv).Interface())
	}
	return names, values, nil

}
