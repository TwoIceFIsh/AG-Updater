package Utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func Downloads(downUrl string, aName string) (string, error) {

	fmt.Println("[-] Connecting to Server...(10s)")
	fmt.Println("[-] Downloading...", downUrl)

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(downUrl)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("[!] Agent Download Fail!", downUrl)
		fmt.Println("Exit AG-Installer 10s...!")
		time.Sleep(10 * time.Second)
		os.Exit(3)
	}

	file, err := os.Create(aName)

	if err != nil {
		fmt.Println("[=================================]")
		fmt.Println("[= Please, Run as Administrator  =]")
		fmt.Println("[=================================]")
		log.Println(err)
		fmt.Println("Exit AG-Installer 10s...!")
		time.Sleep(10 * time.Second)
		os.Exit(3)
	}

	_, err = io.Copy(file, resp.Body)
	fmt.Println("[O] Downloaded a file", aName)
	defer func(file *os.File) {
		err2 := file.Close()
		if err2 != nil {

		}
	}(file)
	return aName, err
}
