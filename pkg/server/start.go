package server

import (
	genericapiserver "github.com/ForbiddenR/apiserver/pkg/server"
	genericoptions "github.com/ForbiddenR/apiserver/pkg/server/options"
	"github.com/ForbiddenR/jxserver/pkg/apiserver"
)

type ServerOptions struct {
	RecommendedOptions *genericoptions.RecommendedOptions
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{}
}

func StartServer(defaults *ServerOptions, stopCh <-chan struct{}) error {
	return nil
}

func (o ServerOptions) Validate() error {
	return nil
}

func (o ServerOptions) Complete() error {
	return nil
}

func (o *ServerOptions) Config() (*apiserver.Config, error) {
	serverConfig := genericapiserver.NewRecommendedConfig()

	if err := o.RecommendedOptions.ApplyTo(serverConfig); err != nil {
		return nil, err
	}

	config := &apiserver.Config{
		GenericConfig: serverConfig,
	}
	
	return config, nil
}

func (o ServerOptions) RunServer(stopCh <-chan struct{}) error {
	config, err := o.Config()
	if err != nil {
		return err
	}

	server, err := config.Complete().New()
	if err != nil {
		return err
	}
	if server == nil {
		return nil
	}
	return nil
}
