package main

import (
	"example/config"
	"example/internal/diregistry"
	"example/internal/router"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func init() {
	log.Println("Init application...")
}

func main() {
	diregistry.BuildDIContainer()
	cfg := diregistry.GetDependency(diregistry.ConfigDIName).(*config.Config)
	gin.SetMode(cfg.Env)

	c := diregistry.GetDependency(diregistry.CronSchedulerDIName).(*cron.Cron)
	defer c.Stop()

	routersInit := router.InitRouter()

	port := fmt.Sprintf(":%d", cfg.HttpAddress)
	server := &http.Server{
		Addr:    port,
		Handler: routersInit,
	}

	log.Printf("[INFO] Start http server listening %s", port)
	server.ListenAndServe()
}
