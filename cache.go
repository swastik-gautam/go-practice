package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx := context.Background()

	err := rdb.Set(ctx, "name", "swastik", 0).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	value, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(value)
}
