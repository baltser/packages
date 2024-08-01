package main

import (
	"fmt"
	"log"
	"packages/config"
)

func main() {
	var cfg config.Config
	if err := cfg.Getting(); err != nil {
		log.Printf("Error loading configuration: %v", err)
		//return
	}

	fmt.Println("Postgres Host:", cfg.PostgresHost)
	fmt.Println("Postgres Port:", cfg.PostgresPort)
	// и т.д.
}
