package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"handling-transactions/employee/config"
	"handling-transactions/employee/internal/app"
)

func main() {
	var wait time.Duration
	var confPath string
	flag.StringVar(&confPath, "config_path", "./config/config.yaml", "the config path yaml file")
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	conf, err := config.ReadConfig(confPath)
	if err != nil {
		panic(err)
	}
	// Run our server in a goroutine so that it doesn't block.
	server := app.New(conf)
	go func() {
		if err := server.Start(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.StopGRPC()
	server.Stop(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
