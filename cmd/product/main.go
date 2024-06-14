package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wlrudi19/go-simple-web/app/product/api"
	"github.com/wlrudi19/go-simple-web/app/product/repository"
	"github.com/wlrudi19/go-simple-web/app/product/service"
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

	fmt.Println("ELASTIC ENGINE PROJECT")
	log.Printf("connected to postgres successfulyy")

	productRepository := repository.NewProductRepository(connDB)
	productLogic := service.NewProductLogic(productRepository)
	productHanlder := api.NewProductHandler(productLogic)
	productRouter := api.NewProductRouter(productHanlder)

	server := http.Server{
		Addr:    "localhost:3013",
		Handler: productRouter,
	}

	fmt.Println("starting server on port 3013...")

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
