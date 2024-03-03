package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stage|prod)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &api{
		config: cfg,
		logger: logger,
	}

	addr := fmt.Sprintf(":%d", cfg.port)
	srv := http.Server{
		Addr:         addr,
		Handler:      app.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

type config struct {
	port int
	env  string
}

type api struct {
	config config
	logger *log.Logger
}
