package main

import (
	"log"
	"os"
)

func main() {
	// Load Confguration for application
	env := os.Args[1]
	appConfig, err := LoadConfig(env)
	if err != nil {
		panic(err)
	}
	log.Printf("config: %+v\n", appConfig)

	// Invoke DB migration and router methods
	DataMigration(appConfig.DatabaseUrl)
	HandlerRouting()
}
