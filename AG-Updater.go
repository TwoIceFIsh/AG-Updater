package main

import (
	"AG-Updater/Utils"
	"fmt"
	"github.com/kardianos/service"
	"log"
	"os"
	"time"
)

type GoWindowsService struct{}

func (goWindowsService *GoWindowsService) Start(windowsService service.Service) error {
	go goWindowsService.run()
	return nil
}

func (goWindowsService *GoWindowsService) run() {
	// Do your work here
	var server Utils.Server
	server.Ip = os.Args[1]
	server.Port = os.Args[2]

	for {
		fmt.Println("arg check", server.Ip, server.Port)
		time.Sleep(5 * time.Second)
	}

}

func (goWindowsService *GoWindowsService) Stop(windowsService service.Service) error {
	return nil
}

func main() {
	serviceConfig := &service.Config{
		Name:        "GoWindowsService",
		DisplayName: "Go Windows service",
		Description: "Go Windows service",
	}

	goWindowsService := &GoWindowsService{}
	windowsService, err := service.New(goWindowsService, serviceConfig)
	if err != nil {
		log.Println(err)
	}

	err = windowsService.Run()
	if err != nil {
		log.Println(err)
	}
}
