package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func getPortFromEnv() int {
	ensureEnvLoaded()

	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		portEnv = "5000"
	}

	port, err := strconv.Atoi(portEnv)
	if err != nil {
		log.Fatalf("PORT must be number, but got \"%s\"", portEnv)
	}

	return port
}

func getMongoURIFromEnv() string {
	ensureEnvLoaded()
	return os.Getenv("MONGO_URI")
}

func getMongoDBNameFromEnv() string {
	ensureEnvLoaded()
	return os.Getenv("MONGO_DB_NAME")
}

var hasLoadedEnv = false

func ensureEnvLoaded() {
	if !hasLoadedEnv {
		hasLoadedEnv = true
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file", "\n\t", err)
		}
	}
}
