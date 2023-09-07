package apiserver

import (
	genericapiserver "github.com/ForbiddenR/apiserver/pkg/server"
	"github.com/ForbiddenR/jxserver/pkg/registry/manage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/valyala/fasthttp"
)

type Config struct {
	GenericConfig   *genericapiserver.RecommendedConfig
	ManageInterface manage.Interface
}

type Server struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
	Manage           manage.Interface
}

type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ManageInterface  manage.Interface
}

type CompletedConfig struct {
	*completedConfig
}

func (cfg *Config) Complete() CompletedConfig {
	c := completedConfig{
		cfg.GenericConfig.Complete(),
		cfg.ManageInterface,
	}
	return CompletedConfig{&c}
}

// New returns a new instance of Server from the given config.
func (c completedConfig) New() (*Server, error) {
	genericServer, err := c.GenericConfig.New("jx-apiserver")
	if err != nil {
		return nil, err
	}

	s := &Server{
		GenericAPIServer: genericServer,
		Manage:           c.ManageInterface,
	}

	v1 := s.GenericAPIServer.Handler.GoRestfulApp.Group("/manage")
	v1.Post("/setLoggingSwitch", func(c *fiber.Ctx) error {
		request := &manage.SetLoggingSwitchRequest{}
		if err := c.BodyParser(request); err != nil {
			log.Error(err)
			return c.Status(fasthttp.StatusOK).JSON(manage.NewResponse(manage.Failed, err.Error()))
		}

		err = s.Manage.SwitchLogging(request.Feature, request.Switch)
		if err != nil {
			log.Error(err)
			return c.Status(fasthttp.StatusOK).JSON(manage.NewResponse(manage.Failed, err.Error()))
		}
		return c.Status(fasthttp.StatusOK).JSON(manage.NewResponse(manage.Succeeded, "success"))
	})

	v1.Post("/getConnections", func(c *fiber.Ctx) error {
		request := &manage.GetConnectionsRequest{}
		if err := c.BodyParser(request); err != nil {
			log.Error(err)
			return c.Status(fasthttp.StatusOK).JSON(manage.NewGetConnectionsResponse(manage.NewResponse(manage.Failed, err.Error()), nil))
		}
		var count uint64
		var err error
		if count, err = s.Manage.GetConnections(request.Type); err != nil {
			log.Error(err)
			return c.Status(fasthttp.StatusOK).JSON(manage.NewGetConnectionsResponse(manage.NewResponse(manage.Failed, err.Error()), nil))
		}
		return c.Status(fasthttp.StatusOK).JSON(manage.NewGetConnectionsResponse(manage.NewResponse(manage.Succeeded, "success"), &manage.GetConnectionsResponseData{Count: count}))
	})

	return s, nil
}
