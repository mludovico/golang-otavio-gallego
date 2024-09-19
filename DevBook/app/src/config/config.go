package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APIURL   string
	Port     int
	HashKey  []byte
	BlockKey []byte
)

func LoadConfig() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file")
	}

	APIURL = os.Getenv("API_URL")
	Port, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
