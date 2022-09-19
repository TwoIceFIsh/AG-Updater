package io

import (
	"log"
	"net"
	"time"
)

type ServerStatus struct {
	Ping int
}

type Command struct {
	Method int8
}

func ReadData(conn net.Conn, server *ServerStatus) {

	data := make([]byte, 4096)

	for {
		n, err := conn.Read(data)
		if err != nil {
			log.Println(err)
			server.Ping = 1
			return
		}
		log.Println("Server send : " + string(data[:n]))
		time.Sleep(3 * time.Second)
	}

}

func WriteData(conn net.Conn) {

	for {
		_, _ = conn.Write([]byte("PING"))
		time.Sleep(3 * time.Second)
	}

}
