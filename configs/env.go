package configs

import (
	"os"
)

func EnvAMQPURL() string {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	return os.Getenv("AMQPURL")
}
func EnvQueueName() string {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	return os.Getenv("QUEUENAME")
}
