package server

import (
	"context"
	"errors"
	"log"

	"github.com/Edilberto-Vazquez/website-services/database"
	"github.com/Edilberto-Vazquez/website-services/repository"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Port        string
	DataBaseUrl string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *gin.Engine
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.DataBaseUrl == "" {
		return nil, errors.New("db url is required")
	}
	broker := &Broker{
		config: config,
		router: gin.Default(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *gin.Engine)) {
	log.Println("Connecting to the DB")
	db, err := database.NewMongoDBRepository(b.config.DataBaseUrl)
	if err != nil {
		log.Fatalf("Could not connect to DB %v", err)
	}
	repository.SetImplementedDB(db)
	b.router.SetTrustedProxies([]string{"127.0.0.1"})
	binder(b, b.router)
	log.Println("Starting server on port", b.config.Port)
	if err := b.router.Run(b.config.Port); err != nil {
		log.Println("Error starting server:", err)
	} else {
		log.Fatalf("Server stopped")
	}
}
