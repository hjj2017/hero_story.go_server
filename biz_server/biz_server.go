package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"hero_story.go_server/biz_server/handler"
)

func main() {
	ctx := context.Background();

	rdb := redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})

	defer rdb.Close()
	rdb.Get(ctx, "User_0")

	handler.Handle()
	handler.CreateHandler()
	fmt.Print("start bizServer")
}
