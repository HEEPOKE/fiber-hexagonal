package main

import (
	"fmt"
	"log"

	"github.com/HEEPOKE/fiber-hexagonal/internals/server"
	"github.com/HEEPOKE/fiber-hexagonal/pkg/configs"
	"github.com/HEEPOKE/fiber-hexagonal/pkg/database"
)

func main() {
	_, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	address := fmt.Sprintf(":%s", configs.Cfg.PORT)
	route := server.NewServer()
	route.Init(address)
}
