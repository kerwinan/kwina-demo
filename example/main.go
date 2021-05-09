package main

import (
	"context"
	"fmt"
	conf "github.com/basharal/config"
	"github.com/kerwinan/kwina-demo/option"
	pb "github.com/kerwinan/kwina-demo/pb/HelloWorld"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/transport/grpc"
)

var ServiceName = "demo.HelloServer"

func main() {
	TestRpcDemo()
}

type options struct {
	ClientOpt option.ClientOpt `yaml:"client_opt"`
}

func TestRpcDemo() {
	opts := new(options)
	err := conf.Parse("./config/client.yaml", &opts.ClientOpt)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = opts.ClientOpt.DemoClientOpt.ETCDHost //地址
	})
	srv := micro.NewService(
		micro.Name(ServiceName),
		micro.Registry(reg),
		micro.Transport(grpc.NewTransport()))
	srv.Init()
	rpcCli := pb.NewGreeterClient(ServiceName, srv.Client())
	input := pb.HelloRequest{
		Name: "haha",
	}
	resp, err := rpcCli.SayHello(context.Background(), &input)
	if err != nil {
		fmt.Println("rpc call err: ", err)
		return
	}
	fmt.Printf("resp: %+v", resp.Message)
}
