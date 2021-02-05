package router

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	ginRouter *gin.Engine
	port      string
)

// StartApp :
func StartApp() {
	flag.StringVar(&port, "PORT", "8085", "Port to start the server")
	flag.Parse()
	ginRouter = gin.Default()
	address := ":" + port
	srv := http.Server{
		Addr:    address,
		Handler: ginRouter,
	}
	mapURL()
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicln(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGABRT)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Panicln(err)
	}
	select {
	case <-ctx.Done():
		fmt.Println("Timeout of 2 sec")
	}
	log.Println("Shutdown gracefully")
}
