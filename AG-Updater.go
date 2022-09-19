package main

import (
	"github.com/kardianos/service"
	"log"
	"net"
	"time"
)

type GoWindowsService struct{}

func (goWindowsService *GoWindowsService) Start(windowsService service.Service) error {
	go goWindowsService.run()
	return nil
}

func (goWindowsService *GoWindowsService) run() {

	for {
		conn, err := net.Dial("tcp", ":8000")
		if nil != err {
			log.Println(err)
		} else {
			go func() {
				data := make([]byte, 4096)

				for {
					n, err := conn.Read(data)
					if err != nil {
						log.Println(err)
						return
					}

					log.Println("Server send : " + string(data[:n]))
					time.Sleep(time.Duration(3) * time.Second)
				}
			}()

			go func() {
				for {
					_, _ = conn.Write([]byte("PING"))
					time.Sleep(time.Duration(3) * time.Second)
				}
			}()

			for {

				time.Sleep(time.Duration(3) * time.Second)
			}
		}
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
