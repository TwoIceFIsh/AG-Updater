package Utils

import (
	"log"
	"time"
)

func Infinity() {
	for {

		log.Println("Service is Running")
		time.Sleep(1 * time.Hour)
	}
}
