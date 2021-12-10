package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mister-turtle/golang-webapp-structure/domain/evidence"
	memory "github.com/mister-turtle/golang-webapp-structure/storage/database_memory"
)

func main() {
	var memoryRepository = memory.NewRepository()
	var serviceIOC = evidence.NewIOCService(&memoryRepository.IOC)
	var newIOCs = []evidence.IOC{
		{"ipv4", "127.0.0.1", time.Now(), "test-service"},
		{"ipv4", "8.8.8.8", time.Now(), "test-service"},
		{"domain", "exmaple.org", time.Now(), "test-service"},
		{"file-md5", "ace3ea1ab3ae555ccf3125c134b6ab2f", time.Now(), "test-service"},
	}

	for _, ioc := range newIOCs {
		err := serviceIOC.Create(context.TODO(), ioc)
		if err != nil {
			panic(err)
		}
	}

	iocs, err := memoryRepository.IOC.FindAll(context.TODO())
	if err != nil {
		panic(err)
	}
	for i, ioc := range iocs {
		fmt.Printf("%02d: [%s] - %s (%s)\n", i, ioc.Discovered, ioc.Type, ioc.Value)
	}

}
