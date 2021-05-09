//+build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"github.com/kerwinan/kwina-demo/client"
	"github.com/kerwinan/kwina-demo/option"
	"github.com/kerwinan/kwina-demo/provider"
	"github.com/kerwinan/kwina-demo/service"
)

func setupProvider() (*provider.Provider, error) {
	wire.Build(provider.NewProvider, service.NewService, client.NewClient, option.NewOptions)
	return nil, nil
}
