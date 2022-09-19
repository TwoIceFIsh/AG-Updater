package io

import (
	"io"
	"log"
	"net"
)

func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 4096)
	for {
		n, err := conn.Read(recvBuf)
		if nil != err {
			if io.EOF == err {
				log.Println(err)
				return
			}
			log.Println(err)
			return
		}
		if 0 < n {
			data := recvBuf[:n]
			log.Println(string(data))
			_, err = conn.Write(data[:n])
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
