package server

import (
	genericapiserver "github.com/ForbiddenR/apiserver/pkg/server"
	genericoptions "github.com/ForbiddenR/apiserver/pkg/server/options"
	"github.com/ForbiddenR/jxserver/pkg/apiserver"
	"github.com/ForbiddenR/jxserver/pkg/registry/manage"
)

type ServerOptions struct {
	RecommendedOptions *genericoptions.RecommendedOptions
	ManageOptions      manage.Interface
}

func NewServerOptions(ifs manage.Interface) *ServerOptions {
	if ifs == nil {
		ifs = &manage.NoopInterface{}
	}
	return &ServerOptions{
		RecommendedOptions: genericoptions.NewRecommendedOptions(),
		ManageOptions:      ifs,
	}
}

func NewStartServer(defaults *ServerOptions, stopCh <-chan struct{}) error {
	o := *defaults

	if err := o.Complete(); err != nil {
		return err
	}

	if err := o.Validate(); err != nil {
		return err
	}

	if err := o.RunServer(stopCh); err != nil {
		return err
	}
	return nil
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
		ManageInterface: o.ManageOptions,
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

	return server.GenericAPIServer.PrepareRun().Run(stopCh)
}
