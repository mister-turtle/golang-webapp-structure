package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/mister-turtle/golang-webapp-structure/evidence"
	memory "github.com/mister-turtle/golang-webapp-structure/storage/memory"
	"github.com/mister-turtle/golang-webapp-structure/webserver"
)

func main() {
	argListen := flag.String("l", "0.0.0.0", "IP address to listen on")
	argPort := flag.Int("p", 9000, "Port to listen on")
	flag.Parse()

	listenAddress := fmt.Sprintf("%s:%d", *argListen, *argPort)

	memoryRepository := memory.NewRepository()
	iocService := evidence.NewIOCService(&memoryRepository.IOC)
	webServer, err := webserver.NewServer(listenAddress, iocService)
	if err != nil {
		log.Fatal(err)
	}

	err = memoryRepository.IOC.Create(context.Background(), evidence.IOC{
		Type:   "file-md5",
		Value:  "12345",
		Date:   time.Now(),
		Source: "preseed",
	})
	log.Printf("Starting web server on %s:%d\n", *argListen, *argPort)
	webServer.Start()

}
