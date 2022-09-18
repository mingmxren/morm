package condition

import (
	"fmt"
	"strings"
)

type mergeStrategy string

const (
	and mergeStrategy = "AND"
	or  mergeStrategy = "OR"
)

func merge(strategy mergeStrategy, conditions ...Condition) Condition {
	s := make([]string, len(conditions))
	args := make([]interface{}, 0)
	for i, c := range conditions {
		s[i] = c.Sql()
		args = append(args, c.Args()...)
	}
	return Manual(fmt.Sprintf("(%s)", strings.Join(s, string(strategy))), args...)
}

func And(conditions ...Condition) Condition {
	return merge(and, conditions...)
}

func Or(conditions ...Condition) Condition {
	return merge(or, conditions...)
}
