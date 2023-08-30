package main

import "github.com/ForbiddenR/jxserver"

func main() {
	server := jxserver.New()
	server.Start(3000)
}