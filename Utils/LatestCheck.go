package Utils

import (
	"log"
	"time"
)

type AgentInfo struct {
	StatusCode     int    `json:"StatusCode"`
	AgentVersion   string `json:"AgentVersion"`
	UpdaterVersion string `json:"UpdaterVersion"`
}

type ReponseJson struct {
	StatusCode     int    `json:"StatusCode"`
	AgentVersion   string `json:"AgentVersion"`
	UpdaterVersion string `json:"UpdaterVersion"`
}

func LatestCheck(ip string, port string) {

	for {

		log.Println("최신버전 확인...")

		time.Sleep(15 * time.Second)
	}

}
