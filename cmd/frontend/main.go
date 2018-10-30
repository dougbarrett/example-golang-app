package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dougbarrett/example-golang-app/config"
	"github.com/dougbarrett/example-golang-app/handler"
	"github.com/dougbarrett/example-golang-app/server"
	"github.com/dougbarrett/example-golang-app/storage/mysql"
)

const defaultConfigPath = "./config/config.json"
const defaultPortVariable = "PORT"
const defaultPort = "3000"

/* Setup is the function `main` calls in order to start the web server bootstrapped with the handler */
func Setup(config *config.Config) {
	ds, err := mysql.NewFrontend(
		config.MySQL.User,
		config.MySQL.Password,
		config.MySQL.Host,
		config.MySQL.DB,
	)

	if err != nil {
		log.Fatalf("Cannot set up mysql: %v", err)
	}

	// cache := cache.New(ds)

	frontend, err := handler.NewFrontend(
		ds,
	)

	if err != nil {
		panic(err)
	}

	e, err := server.NewFrontend(
		frontend,
	)

	if err != nil {
		panic(err)
	}

	port := defaultPort

	if os.Getenv(defaultPortVariable) != "" {
		port = os.Getenv(defaultPortVariable)
	}

	addr := fmt.Sprintf(":%s", port)

	e.Start(addr)
}

func main() {
	configPath := flag.String("config", defaultConfigPath, "path of th config file")

	flag.Parse()

	config, err := config.FromFile(*configPath)

	if err != nil {
		panic(err)
	}
	Setup(config)
}
