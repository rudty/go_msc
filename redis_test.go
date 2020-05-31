package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-redis/redis"
)

// redis-cli config get dir
// redis db 저장 경로 확인
// config.conf 복사 save ""
// redis-server inmemory.conf
var backgroundContext = context.Background()

func TestRedis(t *testing.T) {
	client := redis.NewClient(&redis.Options{})
	fmt.Println(client)
	pong, err := client.Ping(context.Background()).Result()
	fmt.Println(pong, err)

	r := client.Set(backgroundContext, "key", "value", 0)
	client.HSet(backgroundContext, "myuser", "key1", "value1", 0)
	fmt.Println(r.Args())
	fmt.Println(r)
}

func TestRedisSubscribe(t *testing.T) {
	client := redis.NewClient(&redis.Options{})
	pubSub := client.PSubscribe(backgroundContext, "first")
	fmt.Println(pubSub)
	for {
		msg := <-pubSub.Channel()
		fmt.Println(msg)
	}
}
