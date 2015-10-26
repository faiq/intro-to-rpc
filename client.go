package main

import (
	"fmt"
	"gen-go/service"
	"git.apache.org/thrift.git/lib/go/thrift"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	socket, err := thrift.NewTSocket("localhost:8000")
	if err != nil {
		fmt.Printf("There was an error creating your socket! Here it is %v", err)
	}
	protocolFactory := thrift.NewTBinaryProtocolFactory()
	client := service.NewMakeTagsClientFactory(socket, protocolFactory)
	pwd, _ := os.Getwd()
	fileName := pwd + "/img.png"
	fileName, _ = filepath.Abs(fileName)
	imgBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("There was an err reading the file! Here it is %v", err)
	}
	tags, err := client.Generate(service.Image(imgBytes))
	fmt.Printf("These are the tags for your image %v", tags)
}
