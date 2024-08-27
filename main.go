package main

import (
	"example/config"
	"example/internal/diregistry"
	"example/internal/router"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	log.Println("Init application...")
}

func main() {
	diregistry.BuildDIContainer()
	cfg := diregistry.GetDependency(diregistry.ConfigDIName).(*config.Config)
	gin.SetMode(cfg.Env)

	routersInit := router.InitRouter()

	port := fmt.Sprintf(":%d", cfg.HttpAddress)
	server := &http.Server{
		Addr:    port,
		Handler: routersInit,
	}

	log.Printf("[INFO] Start http server listening %s", port)
	server.ListenAndServe()
}
