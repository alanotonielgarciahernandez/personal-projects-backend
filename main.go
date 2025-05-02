// main.go
// Checks the necessary variables and connections to start the HTTP Server.

package main

import (
	"alan-otoniel-garcia-hernandez-personal-projects/backend/db"
	"alan-otoniel-garcia-hernandez-personal-projects/backend/handlers"
	"log"
	"os"
)

func main() {
	if os.Getenv("DB_URI") == "" {
		log.Fatal("Missing DB_URI Environment Variable")
	}

	if os.Getenv("REPOSITORIES_DIR") == "" {
		log.Fatal("Missing REPOSITORIES_DIR Environment Variable")
	}

	if db.CheckDBConnection() != nil {
		log.Fatal("Could not connect with Database")
	}

	log.Println("DB connection established")

	handlers.Managers()
}
