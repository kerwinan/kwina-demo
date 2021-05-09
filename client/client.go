package client

import "github.com/kerwinan/kwina-demo/option"

type Client struct {
	opts       *option.Options
	//DemoClient *client.Client
}

func NewClient(opts *option.Options) (*Client, error) {
	//DemoClient, err := client.NewClientByConf(opts.ClientOpt.DemoClientOpt)
	//if err != nil {
	//	return nil, err
	//}
	//c := Client{
	//	opts:       opts,
	//	DemoClient: DemoClient,
	//}
	//return &c, nil
	return nil, nil
}
