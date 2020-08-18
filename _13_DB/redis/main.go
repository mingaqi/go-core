package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var rClient *redis.Client

func main() {
	rClient = redis.NewClient(&redis.Options{
		Addr:        "127.0.0.1:6379",
		Password:    "", // no password set
		DB:          0,  // use default DB
		PoolSize:    10,
		MaxRetries:  3,
		IdleTimeout: 10 * time.Second,
	})

	pong, err := rClient.Ping().Result()
	fmt.Println(pong, err)

	err = rClient.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rClient.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rClient.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
