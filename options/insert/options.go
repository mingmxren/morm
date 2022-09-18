package insert

type Options struct {
	Fields       []string
	IgnoreFields []string
}

type Option func(o *Options)

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
