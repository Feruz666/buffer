package main

import (
	"context"
	"github.com/Feruz666/buffer"
	"github.com/Feruz666/buffer/handler"
	"github.com/Feruz666/buffer/internal/store"
	"github.com/Feruz666/buffer/internal/worker"
	"github.com/Feruz666/buffer/pkg"
	"github.com/go-resty/resty/v2"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	port         = "3000"
	host         = "localhost"
	redisAddress = "localhost:6379"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	rst := resty.New()

	pkgs := pkg.New(rst)
	wrkr := worker.New(rdb, pkgs)

	buffStore := store.New(rdb)
	buffHandler := handler.New(buffStore)

	srv := &buffer.Server{}
	go func() {
		if err := srv.Run(host, port, buffHandler.RouterFunc()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	go func() {
		for {
			time.Sleep(5 * time.Second)
			wrkr.SaveFacts(context.Background())
		}
	}()

	log.Println("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s\n", err.Error())
	}

}
