package main

import (
	genericapiserver "github.com/ForbiddenR/apiserver/pkg/server"
	"github.com/ForbiddenR/jxserver/pkg/server"
)

func main() {
	stopCh := genericapiserver.SetupSignalHandler()
	// stopCh := make(chan struct{})
	options := server.NewServerOptions(nil)
	err := server.NewStartServer(options, stopCh)
	if err != nil {
		panic(err)
	}
}
