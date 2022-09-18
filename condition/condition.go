package condition

type Condition interface {
	Sql() string
	Args() []interface{}
}
