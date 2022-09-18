package main

import (
	"context"
	"log"

	"github.com/Edilberto-Vazquez/website-services/src/config"
	"github.com/Edilberto-Vazquez/website-services/src/routes"
	"github.com/Edilberto-Vazquez/website-services/src/server"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal("Can't load environment variables")
	}
	s, err := server.NewServer(context.Background(), &server.Config{Port: config.Port, DataBaseUrl: config.DataBaseUrl})
	if err != nil {
		log.Fatal(err)
	}
	s.Start(routes.GetRoutes)
}
