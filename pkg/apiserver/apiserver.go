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
	GenericConfig   genericapiserver.CompletedConfig
	ManageInterface manage.Interface
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
		if perm, ok := c.GetReqHeaders()["Perms"]; !ok || perm != "manage:logging:switch" {
			return c.Status(fasthttp.StatusForbidden).JSON(manage.NewResponse(manage.Failed, "permission denied"))
		}
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
		if perm, ok := c.GetReqHeaders()["Perms"]; !ok || perm != "manage:get:connections" {
			return c.Status(fasthttp.StatusForbidden).JSON(manage.NewResponse(manage.Failed, "permission denied"))
		}
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
	v1.Post("disconnectConnection", func(c *fiber.Ctx) error {
		if perm, ok := c.GetReqHeaders()["Perms"]; !ok || perm != "manage:connection:disconnect" {
			return c.Status(fasthttp.StatusForbidden).JSON(manage.NewResponse(manage.Failed, "permission denied"))
		}
		request := &manage.DisconnectConnectionRequest{}
		if err := c.BodyParser(request); err != nil {
			log.Error(err)
			return c.Status(fasthttp.StatusOK).JSON(manage.NewResponse(manage.Failed, err.Error()))
		}
		if err = s.Manage.CloseConnection(request.Sn); err != nil {
			log.Error(err)
			return c.Status(fasthttp.StatusOK).JSON(manage.NewResponse(manage.Failed, err.Error()))
		}
		return c.Status(fasthttp.StatusOK).JSON(manage.NewResponse(manage.Succeeded, "success"))
	})
	v1.Post("getConnectionStatus", func(c *fiber.Ctx) error {
		if perm, ok := c.GetReqHeaders()["Perms"]; !ok || perm != "manage:connection:status:get" {
			return c.Status(fasthttp.StatusForbidden).JSON(manage.NewResponse(manage.Failed, "permission denied"))
		}
		request := &manage.GetConnectionStatusRequest{}
		if err := c.BodyParser(request); err != nil {
			log.Error(err)
			return c.Status(fasthttp.StatusOK).JSON(manage.NewResponse(manage.Failed, err.Error()))
		}
		var local, remote string
		var err error
		if local, remote, err = s.Manage.GetConnectionStatus(request.Sn); err != nil {
			log.Error(err)
			return c.Status(fasthttp.StatusOK).JSON(manage.NewResponse(manage.Failed, err.Error()))
		}
		return c.Status(fasthttp.StatusOK).JSON(&manage.GetConnectionStatusResponse{
			Response: *manage.NewResponse(manage.Succeeded, "success"),
			Data: &manage.GetConnectionStatusResponseData{
				LocalAddress:  local,
				RemoteAddress: remote,
			},
		})
	})
	v1.Post("getConnectionAlarmRules", func(c *fiber.Ctx) error {
		if perm, ok := c.GetReqHeaders()["Perms"]; !ok || perm != "manage:connection:rules:get" {
			return c.Status(fasthttp.StatusForbidden).JSON(manage.NewResponse(manage.Failed, "permission denied"))
		}
		rule, limit, err := s.Manage.GetConnectionAlarmRule()
		if err != nil {
			log.Error(err)
			return c.Status(fasthttp.StatusOK).JSON(manage.NewResponse(manage.Failed, err.Error()))
		}
		return c.Status(fasthttp.StatusOK).JSON(&manage.GetConnectionAlarmRulesResponse{
			Response: *manage.NewResponse(manage.Succeeded, "success"),
			Data: &manage.GetConnectionAlarmRulesResponseData{
				Rule:  rule,
				Limit: limit,
			},
		})
	})
	v1.Post("setConnectionAlarmRules", func(c *fiber.Ctx) error {
		if perm, ok := c.GetReqHeaders()["Perms"]; !ok || perm != "manage:connection:rules:set" {
			return c.Status(fasthttp.StatusForbidden).JSON(manage.NewResponse(manage.Failed, "permission denied"))
		}
		request := &manage.SetConnectionAlarmRulesRequest{}
		if err := c.BodyParser(request); err != nil {
			log.Error(err)
			return c.Status(fasthttp.StatusOK).JSON(manage.NewResponse(manage.Failed, err.Error()))
		}
		err := s.Manage.SetConnectionAlarmRules(request.Rule, request.Limit)
		if err != nil {
			log.Error(err)
			return c.Status(fasthttp.StatusOK).JSON(manage.NewResponse(manage.Failed, err.Error()))
		}
		return c.Status(fasthttp.StatusOK).JSON(manage.NewResponse(manage.Succeeded, "success"))
	})
	return s, nil
}
