package service

import (
	"github.com/kerwinan/kwina-demo/client"
	"github.com/kerwinan/kwina-demo/option"
)

type Service struct {
	HelloService *HelloService
}

func NewService(opts *option.Options, c *client.Client) (*Service, error) {
	helloService, err := NewHelloService(opts, c)
	if err != nil {
		return nil, err
	}

	s := Service{
		HelloService: helloService,
	}
	return &s, nil
}