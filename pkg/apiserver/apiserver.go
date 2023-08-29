package apiserver

type Config struct {

}

type Server struct {
}

type completedConfig struct {
}

type CompletedConfig struct {
	*completedConfig
}

func (cfg *Config) Complete() CompletedConfig {
	c := completedConfig{}
	return CompletedConfig{&c}
}

func (c completedConfig) New() (*Server, error) {
	return nil, nil
}