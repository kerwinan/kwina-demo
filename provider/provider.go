package provider

import (
	"github.com/kerwinan/kwina-demo/client"
	"github.com/kerwinan/kwina-demo/option"
	"github.com/kerwinan/kwina-demo/service"
)

//Provider Provider
type Provider struct {
	opts          *option.Options
	c             chan error
	svc           *service.Service
	cli           *client.Client
	HelloProvider *HelloProvider
}

//NewProvider NewProvider
func NewProvider(opts *option.Options, svc *service.Service, cli *client.Client) (*Provider, error) {
	p := Provider{
		opts: opts,
		c:    make(chan error),
		svc:  svc,
		cli:  cli,
	}

	if opts.ProviderOpt.HelloProviderOpt.Enable {
		HTTPProviderOpt, err := NewHelloProvider(opts, svc.HelloService)
		if err != nil {
			return nil, err
		}
		p.HelloProvider = HTTPProviderOpt
	}

	return &p, nil
}

//Run Run
func (p *Provider) Run() chan error {

	if p.opts.ProviderOpt.HelloProviderOpt.Enable {
		go func() {
			if err := p.HelloProvider.Start(); err != nil {
				p.c <- err
			}
		}()
	}

	return p.c
}

//Close Close
func (p *Provider) Close() {
	if p.opts.ProviderOpt.HelloProviderOpt.Enable {
		p.HelloProvider.Stop()
	}
	close(p.c)
}
