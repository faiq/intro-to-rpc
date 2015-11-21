package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/faiq/intro-to-rpc/gen-go/service"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	socket, err := thrift.NewTSocket("localhost:8090")
	if err != nil {
		fmt.Printf("There was an error creating your socket! Here it is %v", err)
	}
	transport := thrift.NewTBufferedTransport(socket, 1024)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := service.NewMakeTagsClientFactory(transport, protocolFactory)
	pwd, _ := os.Getwd()
	fileName := pwd + "/img.png"
	fileName, _ = filepath.Abs(fileName)
	imgBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("There was an err reading the file! Here it is %v", err)
	}
	tags, err := client.Generate(service.Image(imgBytes))
	fmt.Printf("These are the tags for your image %v", tags)
	socket.Close()
}
