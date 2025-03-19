package jxserver

import "github.com/gofiber/fiber/v2"

type Options struct {
	Version string
	Handler []map[string]Func
}

type Func fiber.Handler
