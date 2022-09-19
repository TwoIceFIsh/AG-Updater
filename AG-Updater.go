package main

import (
	"AG-Updater/io"
	"fmt"
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
	var ServerStatus io.ServerStatus
	for {
		conn, err := net.Dial("tcp", ":8000")
		if nil != err {
			log.Println(err)
			ServerStatus.Ping = 1
		} else {

			// 통신이 수립 되었을때 하고 싶은 일
			fmt.Println("Connected")
			go io.ReadData(conn, &ServerStatus)
			go io.WriteData(conn)

			// goroutine 종료 방지 무한 루프
			for {
				fmt.Println("종료 방지 무한 루프", ServerStatus.Ping)
				if ServerStatus.Ping == 1 {
					break
				}
				time.Sleep(3 * time.Second)
			}
		}
		// 통신 재게 무한 루프
		ServerStatus.Ping = 0
		fmt.Println("통신 탐지", ServerStatus.Ping)
		time.Sleep(15 * time.Second)
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
