package dbtag

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/mingmxren/morm/structreflect"
)

type dbTag struct {
	name string
}

var (
	dbTagRe = regexp.MustCompile("[a-zA-Z0-9_]+")
)

func parseDbTag(tag string) (t *dbTag, err error) {
	if !dbTagRe.Match([]byte(tag)) {
		return nil, fmt.Errorf("parseDbTag tag:%s format error", tag)
	}
	return &dbTag{
		name: tag,
	}, nil
}

func getDbTags(rt reflect.Type) ([]*dbTag, error) {
	tags, err := structreflect.CachedStructTags(rt, "db",
		func(tag string) (structreflect.Tag, error) { return parseDbTag(tag) })
	if err != nil {
		return nil, err
	}
	dbTags := make([]*dbTag, rt.NumField())
	for i, t := range tags {
		if t != nil {
			dbTags[i] = t.(*dbTag)
		}
	}
	return dbTags, nil
}
