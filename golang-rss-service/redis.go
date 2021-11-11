package main

import (
	"github.com/go-redis/redis/v7"
	"os"
)

var client *redis.Client

//init initialise a redis instance when the service is started
func init() {

	dsn := os.Getenv("REDIS_DSN")

	if len(dsn) == 0 {

		dsn = "localhost:6379"
	}

	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})

	_, err := client.Ping().Result()

	if err != nil {

		panic(err)
	}
}
