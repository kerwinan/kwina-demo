package Demo1ServerPvtService

import "context"

type Demo1ServerPvtService interface {
	SayHello(ctx context.Context, input *SayHelloInput, output *SayHelloOutput) error
}

type SayHelloInput struct {
	Name string
}

type SayHelloOutput struct {
	Msg string
}