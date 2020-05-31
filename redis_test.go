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

func TestRedisHSet(t *testing.T) {
	client := redis.NewClient(&redis.Options{})
	client.HSet(backgroundContext, "user1", "email", "aa@aa.com")
	client.HSet(backgroundContext, "user1", "lang", "ko")

	m, err := client.HGetAll(backgroundContext, "user1").Result()

	if err != nil {
		t.Error(err)
	}

	fmt.Println(m)

	email, err := client.HGet(backgroundContext, "user1", "email").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(email)
}

func TestRedisZRange(t *testing.T) {
	client := redis.NewClient(&redis.Options{})
	client.ZAdd(backgroundContext, "myvalue", &redis.Z{
		Member: "냐옹",
		Score:  1,
	})
	client.ZAdd(backgroundContext, "myvalue", &redis.Z{
		Member: 3,
		Score:  2,
	}, &redis.Z{
		Member: "꿀꿀",
		Score:  0,
	})

	r, err := client.ZRangeWithScores(backgroundContext, "myvalue", 0, -1).Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
}

func TestSlic(t *testing.T) {
	a := make([]int, 10)
	a[0] = 1
	a[1] = 2
	b := a[2:]
	b[0] = 3
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(cap(a))
	fmt.Println(cap(b))

}
