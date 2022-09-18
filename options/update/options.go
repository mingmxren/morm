package update

import (
	"github.com/mingmxren/morm/condition"
	"github.com/mingmxren/morm/limit"
)

type Options struct {
	Fields       []string
	IgnoreFields []string
	Condition    condition.Condition
	Limit        limit.Limit
}

type Option func(o *Options)

func WithCondition(c condition.Condition) Option {
	return func(o *Options) {
		o.Condition = c
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
