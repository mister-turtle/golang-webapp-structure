package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mister-turtle/golang-webapp-structure/evidence"
	"github.com/mister-turtle/golang-webapp-structure/httpd"
	"github.com/mister-turtle/golang-webapp-structure/storage/memory"
)

func main() {
	argListen := flag.String("l", "0.0.0.0", "IP address to listen on")
	argPort := flag.Int("p", 9000, "Port to listen on")
	flag.Parse()

	listenAddress := fmt.Sprintf("%s:%d", *argListen, *argPort)

	memoryRepository := memory.NewRepository()
	iocService := evidence.NewIOCService(&memoryRepository.IOC)
	webServer, err := httpd.NewServer(listenAddress, iocService)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting web server on %s:%d\n", *argListen, *argPort)
	webServer.Start()

}
