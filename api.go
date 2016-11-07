package main

import (
	"fmt"
	"github.com/gauravtayal0/sysutil/controller"
)

const (
	AppName = "sysutil"
)

func main() {
	toWS := make(chan string)
	webServer := controller.NewWebServer(toWS, 8090)
	err := webServer.Start()
	if err != nil {
		fmt.Print(err)
	}
}
