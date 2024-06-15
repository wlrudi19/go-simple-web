package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	prodApi "github.com/wlrudi19/go-simple-web/app/product/api"
	prodRepo "github.com/wlrudi19/go-simple-web/app/product/repository"
	prodSvc "github.com/wlrudi19/go-simple-web/app/product/service"
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

	//init user
	userRepository := repository.NewUserRepository(connDB, connRedis)
	userLogic := service.NewUserLogic(userRepository)
	userHandler := api.NewUserHandler(userLogic)
	userRouter := api.NewUserRouter(userHandler)

	//init product
	productRepository := prodRepo.NewProductRepository(connDB)
	productLogic := prodSvc.NewProductLogic(productRepository)
	productHanlder := prodApi.NewProductHandler(productLogic)
	productRouter := prodApi.NewProductRouter(productHanlder)

	r := chi.NewRouter()
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3014"}, // Adjust this to your frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(corsMiddleware.Handler)

	// Serve static files from the frontend directory
	r.Handle("/", http.FileServer(http.Dir("./frontend/")))

	//r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Mount("/users", userRouter)
		r.Mount("/products", productRouter)
	})

	server := http.Server{
		Addr:    "localhost:3012",
		Handler: r,
	}

	fmt.Println("starting server on port 3012...")

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
