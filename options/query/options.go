package query

import (
	"github.com/mingmxren/morm/condition"
	"github.com/mingmxren/morm/limit"
	"github.com/mingmxren/morm/orderby"
)

type Options struct {
	Fields       []string
	IgnoreFields []string
	Condition    condition.Condition
	OrderBy      orderby.OrderBy
	Limit        limit.Limit
	ForUpdate    bool
}

type Option func(o *Options)

func WithCondition(c condition.Condition) Option {
	return func(o *Options) {
		o.Condition = c
	}
}

func WithOrderBy(orderby orderby.OrderBy) Option {
	return func(o *Options) {
		o.OrderBy = orderby
	}
}

func WithLimit(limit limit.Limit) Option {
	return func(o *Options) {
		o.Limit = limit
	}
}

func WithFields(fields ...string) Option {
	return func(o *Options) {
		o.Fields = fields
	}
}

func WithIgnoreFields(fields ...string) Option {
	return func(o *Options) {
		o.IgnoreFields = fields
	}
}

func WithForUpdate(forUpdate bool) Option {
	return func(o *Options) {
		o.ForUpdate = forUpdate
	}
}
