package main

// Option nolint:golint
type Option func(*Options)

// Options nolint:golint
type Options struct {
	Name string
	Age  int
}

//nolint:golint
var (
	DefaultName = "defaultName"
	DefaultAge  = 10
)

// NewOptions default constructor
func NewOptions(opts ...Option) *Options {

	// Init
	options := &Options{
		Name: DefaultName,
		Age:  DefaultAge,
	}

	for _, opt := range opts {
		opt(options)
	}

	return options
}

// Name Init name
func Name(name string) Option {
	return func(opt *Options) {
		opt.Name = name
	}
}

// Age Init age
func Age(age int) Option {
	return func(opt *Options) {
		opt.Age = age
	}
}
