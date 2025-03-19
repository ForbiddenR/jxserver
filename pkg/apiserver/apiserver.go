package apiserver

import (
	"net/http"
	"slices"

	genericapiserver "github.com/ForbiddenR/apiserver/pkg/server"
	"github.com/ForbiddenR/jxserver/pkg/registry/manage"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
)

const permKey = "Perms"

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

	validate := validator.New()

	app := s.GenericAPIServer.Handler.GoRestfulApp

	handler := func(w http.ResponseWriter, r *http.Request) {
		s.Manage.BeforeGetMetrics()
		promhttp.Handler().ServeHTTP(w, r)
	}
	app.Get("/metrics", adaptor.HTTPHandlerFunc(handler))

	v1 := s.GenericAPIServer.Handler.GoRestfulApp.Group("/manage")
	v1.Use(recover.New())
	v1.Use(func(c *fiber.Ctx) error {
		if c.Get("Content-Type", "") != "application/json" {
			return c.Status(fasthttp.StatusOK).JSON(manage.NewResponse(manage.Failed, "Invalid Content-Type"))
		}
		return c.Next()
	})

	v1.Post("/setLoggingSwitch", func(c *fiber.Ctx) error {
		if ok := exist(c.GetReqHeaders(), permKey, "manage:logging:switch"); !ok {
			return permissionError(c)
		}
		request := &manage.SetLoggingSwitchRequest{}
		if err := c.BodyParser(request); err != nil {
			return wrapError(c, err)
		}
		err = s.Manage.SwitchLogging(request.Feature, request.Switch)
		if err != nil {
			return wrapError(c, err)
		}
		return c.Status(fasthttp.StatusOK).JSON(manage.NewResponse(manage.Succeeded, "success"))
	})

	v1.Post("/getConnections", func(c *fiber.Ctx) error {
		if ok := exist(c.GetReqHeaders(), permKey, "manage:get:connections"); !ok {
			return permissionError(c)
		}
		request := &manage.GetConnectionsRequest{}
		if err := c.BodyParser(request); err != nil {
			return c.Status(fasthttp.StatusOK).JSON(manage.NewGetConnectionsResponse(manage.NewResponse(manage.Failed, err.Error()), nil))
		}
		var count uint64
		var host string
		var err error
		if count, host, err = s.Manage.GetConnections(request.Type); err != nil {
			return c.Status(fasthttp.StatusOK).JSON(manage.NewGetConnectionsResponse(manage.NewResponse(manage.Failed, err.Error()), nil))
		}
		return c.Status(fasthttp.StatusOK).JSON(manage.NewGetConnectionsResponse(manage.NewResponse(manage.Succeeded, "success"), &manage.GetConnectionsResponseData{Count: count, Host: host}))
	})

	v1.Post("/disconnectConnection", func(c *fiber.Ctx) error {
		if ok := exist(c.GetReqHeaders(), permKey, "manage:connection:disconnect"); !ok {
			return permissionError(c)
		}
		request := &manage.DisconnectConnectionRequest{}
		if err := c.BodyParser(request); err != nil {
			return wrapError(c, err)
		}
		if err = s.Manage.CloseConnection(request.Sn); err != nil {
			return wrapError(c, err)
		}
		return wrapSuccess(c, "success")
	})

	v1.Post("/getConnectionStatus", func(c *fiber.Ctx) error {
		if ok := exist(c.GetReqHeaders(), permKey, "manage:connection:status:get"); !ok {
			return permissionError(c)
		}
		request := &manage.GetConnectionStatusRequest{}
		if err := c.BodyParser(request); err != nil {
			return wrapError(c, err)
		}
		var ct, lht, local, remote string
		var err error
		if ct, lht, local, remote, err = s.Manage.GetConnectionStatus(request.Sn); err != nil {
			return wrapError(c, err)
		}
		return c.Status(fasthttp.StatusOK).JSON(&manage.GetConnectionStatusResponse{
			Response: *manage.NewResponse(manage.Succeeded, "success"),
			Data: &manage.GetConnectionStatusResponseData{
				HandlerCreateTime: ct,
				LastHearbeatTime:  lht,
				LocalAddress:      local,
				RemoteAddress:     remote,
			},
		})
	})

	v1.Post("/updateLogLevel", func(c *fiber.Ctx) error {
		if ok := exist(c.GetReqHeaders(), permKey, "manage:update:log:level"); !ok {
			return permissionError(c)
		}
		request := &manage.UpdateLogLevelRequest{}
		if err := c.BodyParser(request); err != nil {
			return wrapError(c, err)
		}
		if err := validate.Struct(request); err != nil {
			return wrapError(c, err)
		}
		if err := s.Manage.UpdateLogLevel(request.Location, request.Level); err != nil {
			return wrapError(c, err)
		}
		return wrapSuccess(c, "success")
	})
	return s, nil
}

func exist[S ~map[D][]D, D ~string](m S, key, target D) bool {
	strSlice, ok := m[key]
	return ok && slices.Contains(strSlice, target)
}

func permissionError(c *fiber.Ctx) error {
	return c.Status(fasthttp.StatusOK).JSON(manage.NewResponse(manage.Failed, "Permission denied"))
}

func wrapError(c *fiber.Ctx, err error) error {
	return c.Status(fasthttp.StatusOK).JSON(manage.NewResponse(manage.Failed, err.Error()))
}

func wrapSuccess(c *fiber.Ctx, msg string) error {
	return c.Status(fasthttp.StatusOK).JSON(manage.NewResponse(manage.Succeeded, msg))
}
