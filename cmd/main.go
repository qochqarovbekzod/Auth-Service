package main

import (
	"log"
	"sync"
	router "users/api"
	"users/api/handlers"
	"users/cmd/server"
	"users/config"
	l "users/pkg/logger"
	"users/storage/postgres"

	"go.uber.org/zap"
)

var logger *zap.Logger

func initLog() {
	log, err := l.NewLogger()
	if err != nil {
		panic(err)
	}
	logger = log
}

func main() {
	initLog()
	db, err := postgres.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cfg := config.Load()

	router := router.RouterApi(handlers.NewHandler(postgres.NewUserRepo(db), logger))

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := router.Run(cfg.HTTP_PORT)
		if err != nil {
			log.Fatal(err)
		}
	}()

	server.ServerRun(postgres.NewUserRepo(db), &cfg)
	wg.Wait()
}
