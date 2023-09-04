package apiserver

import genericapiserver "github.com/ForbiddenR/apiserver/pkg/server"

type Config struct {
	GenericConfig *genericapiserver.RecommendedConfig
}

type Server struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
}

type CompletedConfig struct {
	*completedConfig
}

func (cfg *Config) Complete() CompletedConfig {
	c := completedConfig{}
	return CompletedConfig{&c}
}

func (c completedConfig) New() (*Server, error) {
	genericServer, err := c.GenericConfig.New("jx-apiserver")
	if err != nil {
		return nil, err
	}

	s := &Server{
		GenericAPIServer: genericServer,
	}



	return s, nil
}
