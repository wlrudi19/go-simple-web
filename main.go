package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wlrudi19/go-simple-web/app/user/api"
	"github.com/wlrudi19/go-simple-web/app/user/repository"
	"github.com/wlrudi19/go-simple-web/app/user/service"
	"github.com/wlrudi19/go-simple-web/config"
)

func main() {
	loadConfig := config.LoanConfig()
	connDB, connRedis, err := config.ConnectConfig(loadConfig.Database, loadConfig.Redis)

	if err != nil {
		log.Fatalf("error connecting to postgres :%v", err)
		return
	}
	defer connDB.Close()
	defer connRedis.Close()

	fmt.Println("GO SIMPLE WEB PROJECT")
	log.Printf("connected to postgres successfulyy")
	log.Printf("connected to redis successfulyy")

	userRepository := repository.NewUserRepository(connDB, connRedis)
	userLogic := service.NewUserLogic(userRepository)
	userHandler := api.NewUserHandler(userLogic)
	userRouter := api.NewUserRouter(userHandler)

	server := http.Server{
		Addr:    "localhost:3012",
		Handler: userRouter,
	}

	fmt.Println("starting server on port 3012...")

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
