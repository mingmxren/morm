package condition

import (
	"fmt"
	"strings"
)

func newCompareCondition(compare string, field string, value interface{}) Condition {
	return Manual(fmt.Sprintf("%s %s ?", field, compare), value)
}

// Equal field = value
func Equal(field string, value interface{}) Condition {
	return newCompareCondition("=", field, value)
}

// NotEqual field != value
func NotEqual(field string, value interface{}) Condition {
	return newCompareCondition("!=", field, value)
}

// Less field < value
func Less(field string, value interface{}) Condition {
	return newCompareCondition("<", field, value)
}

// LessEqual field <= value
func LessEqual(field string, value interface{}) Condition {
	return newCompareCondition("<=", field, value)
}

// Greater field > value
func Greater(field string, value interface{}) Condition {
	return newCompareCondition(">", field, value)
}

// GreaterEqual field >= value
func GreaterEqual(field string, value interface{}) Condition {
	return newCompareCondition(">=", field, value)
}

// NotNull field IS NOT NULL
func NotNull(field string) Condition {
	return Manual(fmt.Sprintf("%s IS NOT NULL", field))
}

// Null field IS NULL
func Null(field string) Condition {
	return Manual(fmt.Sprintf("%s IS NULL", field))
}

// Like field LIKE value
func Like(field string, value string) Condition {
	return newCompareCondition("LIKE", field, value)
}

// In field in (args...)
func In(field string, args ...interface{}) Condition {
	s := strings.Repeat(",?", len(args))
	if len(s) > 0 {
		s = s[1:]
	}
	return Manual(fmt.Sprintf("%s IN (%s)", field, s), args...)
}

// Between field BETWEEN begin AND end
func Between(field string, begin, end interface{}) Condition {
	return Manual(fmt.Sprintf("%s BETWEEN ? AND ?", field), begin, end)
}

// Range field >= begin AND field < end
func Range(field string, begin, end interface{}) Condition {
	return And(GreaterEqual(field, begin), Less(field, end))
}
