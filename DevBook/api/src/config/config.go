package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectionString = ""
	DbPort           = 0
	ApiPort          = 0
	SecretKey        = []byte{}
)

func LoadConfig() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DbPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		DbPort = 3306
	}

	ApiPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		ApiPort = 9000
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
