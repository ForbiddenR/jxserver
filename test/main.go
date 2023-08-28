package main

import "github.com/ForbiddenR/apiserver"

func main() {
	server := apiserver.New()
	server.Start(3000)
}