package services

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv" // package used to read the .env file
	"log"
	"os" // used to read the environment variable
)

func InitRedis() *redis.Client {

	e:= godotenv.Load()
	if e != nil {
		log.Fatalf("Error loading .env file")
	}

	redis_host := os.Getenv("REDIS_HOST")
	redis_port := os.Getenv("REDIS_PORT")
	redis_password := os.Getenv("REDIS_PASSWORD")
	rdb := redis.NewClient(&redis.Options{
		Addr:     redis_host + ":" + redis_port,
		Password: redis_password, // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("Successfully connected to Redis!")
	return rdb
}
