package main

import "fmt"

//Service server
type Service interface {
	Output()
}

type service struct {
	opts *Options
}

// NewService default constructor
func NewService(opts ...Option) Service {

	// Init parameters
	options := NewOptions(opts...)
	return &service{
		opts: options,
	}
}

func (p *service) Output() {
	fmt.Printf("name:%s\tage:%d", p.opts.Name, p.opts.Age)
}
