package jxserver

import (
	"fmt"
	"net"

	"github.com/gofiber/fiber/v2"
)

type apiServer struct {
	// Fiber instance
	server *fiber.App
}

// New returns a new Fiber instance
func New() *apiServer {
	return &apiServer{
		server: fiber.New(),
	}
}

func (s *apiServer) Start(port int) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	s.server.Listener(ln)
}

func (s *apiServer) Stop() {
	s.server.Shutdown()
}

func (s *apiServer) Group(prefix string) fiber.Router {
	return s.server.Group(prefix)
}

func AddRoute(router fiber.Router, method string, path string, handler fiber.Handler) {
	router.Add(method, path, handler)
}

func (s *apiServer) Install(options Options) {
	router := s.Group(options.Version)
	for _, handlers := range options.Handler {
		for method, handler := range handlers {
			AddRoute(router, fiber.MethodPost, method, handler)
		}
	}
}
