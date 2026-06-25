package helper

import (
	"github.com/joho/godotenv"
	"log"
)

var Env = map[string]string{}

func SetupConfig() {
	var err error
	Env, err = godotenv.Read(".env")
	if err != nil {
		log.Fatal("failed to read env file: ", err)
	}
}

func GetEnv(key string, val string) string {
	// val use as default value
	result := Env[key]
	if result == "" {
		result = val
	}
	return result
}
