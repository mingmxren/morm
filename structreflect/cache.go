package structreflect

import (
	"fmt"
	"reflect"
	"sync"
)

var fieldsCache = sync.Map{}

func CachedStructFields(t reflect.Type) (fields []reflect.StructField, err error) {
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("CachedStructFields t:%v is not Struct", t)
	}

	if value, ok := fieldsCache.Load(t); ok {
		return value.([]reflect.StructField), nil
	}

	fields = make([]reflect.StructField, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if fields != nil {
			fields[i] = f
		}
	}
	if value, loaded := fieldsCache.LoadOrStore(t, fields); loaded {
		return value.([]reflect.StructField), nil
	}
	return fields, nil
}

var tagsCache = sync.Map{}

type Tag interface{}
type ParseTag func(tag string) (Tag, error)

func CachedStructTags(t reflect.Type, name string, parse ParseTag) (tags []Tag, err error) {
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("CachedStructTags t:%v is not Struct", t)
	}
	var tagsMap *sync.Map
	if value, ok := tagsCache.Load(t); ok {
		tagsMap = value.(*sync.Map)
	} else {
		tagsMap = &sync.Map{}
		if value, loaded := tagsCache.LoadOrStore(t, tagsMap); loaded {
			tagsMap = value.(*sync.Map)
		}
	}

	if value, ok := tagsMap.Load(name); ok {
		return value.([]Tag), nil
	}

	tags = make([]Tag, t.NumField())
	fields, err := CachedStructFields(t)
	if err != nil {
		return nil, err
	}
	for i, f := range fields {
		tagStr := f.Tag.Get(name)
		if len(tagStr) == 0 {
			continue
		}
		if t, err := parse(tagStr); err != nil {
			return nil, err
		} else if tags != nil {
			tags[i] = t
		}
	}

	if value, loaded := tagsMap.LoadOrStore(name, tags); loaded {
		tags = value.([]Tag)
	}
	return tags, nil
}
