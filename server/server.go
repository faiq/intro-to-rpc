package main

import (
	"fmt"
	"os"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/faiq/intro-to-rpc/gen-go/service"
	"github.com/faiq/intro-to-rpc/server/handler"
)

func main() {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTServerSocket("localhost:8090")
	if err != nil {
		fmt.Printf("There was an error creating your socket! Here it is %v", err)
	}
	transportFactory := thrift.NewTTransportFactory()
	processor := service.NewMakeTagsProcessor(handler.NewTagsHandler(os.Getenv("ACCESS_TOKEN")))
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	fmt.Printf("server listening on %s\n", "localhost:8090")
	server.Serve()
}
