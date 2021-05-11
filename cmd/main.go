package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pratheeshpcplpta/ecs-task-definition-ci-cd/config"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.GetConfig()
	if err != nil {
		logger.Fatal("error parsing config", zap.Error(err))
	}

	if cfg == nil {
		logger.Fatal("unable to load configurations")
		return
	}

	r := gin.New()
	r.RedirectTrailingSlash = true
	r.NoRoute(noRouteFunc)

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	var routeString string
	if cfg.Env == config.Production || cfg.Env == config.Local {
		routeString = "v1"
	} else {
		routeString = fmt.Sprintf("%s/v1", cfg.Env)
	}

	v1 := r.Group(routeString)
	v1.GET("/hi", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello")
	})
	v1.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Pong")
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r,
	}
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func noRouteFunc(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, fmt.Sprintf("unable to find a valid http route for %s", ctx.Request.RequestURI))
}
