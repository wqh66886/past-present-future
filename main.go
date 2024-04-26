package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/wqh66886/past-present-future/define"
	"github.com/wqh66886/past-present-future/initial"
	"github.com/wqh66886/past-present-future/middleware"
	"github.com/wqh66886/past-present-future/route"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CrosHandler())
	route.Setup(r)

	server := &http.Server{
		Addr:    ":8888",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown Server.....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func init() {
	initial.InitConfig()
	initial.InitMysql(define.Cfg)
}
