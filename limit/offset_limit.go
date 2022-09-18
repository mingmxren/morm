package limit

import "fmt"

type Limit string

func LimitBy(limit int) Limit {
	return Limit(fmt.Sprintf("LIMIT %d", limit))
}

func LimitOffset(limit int, offset int) Limit {
	return Limit(fmt.Sprintf("LIMIT %d OFFSET %d", limit, offset))
}
