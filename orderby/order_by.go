package orderby

import "fmt"

type OrderBy string

func Desc(fields string) OrderBy {
	return OrderBy(fmt.Sprintf(" ORDER BY %s DESC ", fields))
}

func Asc(fields string) OrderBy {
	return OrderBy(fmt.Sprintf(" ORDER BY %s ASC ", fields))
}
