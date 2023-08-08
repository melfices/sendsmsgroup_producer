package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvAMQPURL() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("AMQPURL")
}
func EnvQueueName() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("QUEUENAME")
}
