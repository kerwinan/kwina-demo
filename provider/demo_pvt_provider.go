package provider

import (
	"github.com/kerwinan/kwina-demo/option"
	pb "github.com/kerwinan/kwina-demo/pb/HelloWorld"
	"github.com/kerwinan/kwina-demo/service"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/transport/grpc"
	"time"
)

//HelloProvider HelloProvider
type HelloProvider struct {
	opts         *option.Options
	helloService *service.HelloService
	rpcSrv       micro.Service
}

//NewHelloProvider NewHelloProvider
func NewHelloProvider(opts *option.Options, helloService *service.HelloService) (*HelloProvider, error) {
	f := HelloProvider{
		opts:         opts,
		helloService: helloService,
	}

	return &f, nil
}

//Start Start
func (h *HelloProvider) Start() error {
	//var err error
	/* TODO 自己实现一套etcd的服务注册与发现框架 */
	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = h.opts.ServiceOpt.HelloServiceOpt.ETCDHost //地址
	})

	h.rpcSrv = micro.NewService(
		micro.Name(h.opts.ServiceOpt.HelloServiceOpt.Service),
		micro.RegisterTTL(time.Second*30),
		micro.Version(h.opts.ServiceOpt.HelloServiceOpt.Version),
		micro.Registry(reg),
		micro.Transport(grpc.NewTransport()))

	h.rpcSrv.Init()
	pb.RegisterGreeterHandler(h.rpcSrv.Server(), h.helloService)
	return h.rpcSrv.Run()
}

//Stop Stop
func (h *HelloProvider) Stop() error {
	return nil
}
