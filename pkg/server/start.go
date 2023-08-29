package server

import "github.com/ForbiddenR/apiserver/pkg/apiserver"

type ServerOptions struct {
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
	return nil, nil
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