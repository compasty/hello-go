package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	val, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(val)
	}
}
