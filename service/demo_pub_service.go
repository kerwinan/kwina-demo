package service

import (
	"context"
	"github.com/kerwinan/go/kit/log"
	"github.com/kerwinan/kwina-demo/client"
	"github.com/kerwinan/kwina-demo/option"
	pbhw "github.com/kerwinan/kwina-demo/pb/HelloWorld"
)

type HelloService struct {
	opts *option.Options
	c    *client.Client
}

func NewHelloService(opts *option.Options, c *client.Client) (*HelloService, error) {
	h := HelloService{
		opts: opts,
		c:    c,
	}
	return &h, nil
}

func (h *HelloService) SayHello(ctx context.Context, input *pbhw.HelloRequest, output *pbhw.HelloReply) (err error) {
	log.X.Infof("Received HelloService.Call request, %v", input.Name)
	output.Message = "Hello " + input.Name
	return nil
}

// TODO
//func (h *HelloService) SayHello(ctx context.Context, input *idl.SayHelloInput, output *idl.SayHelloOutput) error {
//	log.X.Infof("Received HelloService.Call request, %+v", input.Name)
//	output.Msg = "haha"
//	return nil
//}
