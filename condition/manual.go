package condition

type manualCondition struct {
	sql  string
	args []interface{}
}

func (c *manualCondition) Sql() string {
	return c.sql
}

func (c *manualCondition) Args() []interface{} {
	return c.args
}

func Manual(sql string, args ...interface{}) Condition {
	return &manualCondition{
		sql:  sql,
		args: args,
	}
}
