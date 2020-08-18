package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
	"time"
)

var rdb *redis.Client

func TestRedis(t *testing.T) {

	err := initRedis()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("链接成功")
	err = rdb.Set("key", "value", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	key := "zset"
	ite := []redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	rdb.ZAdd(key, ite...)

	/////////////////////读取//////////////////////////
	// 把Golang的分数加10
	// 取分数最高的3个
	ret, err := rdb.ZRevRangeWithScores(key, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
	// 取95~100分的

	fmt.Println(rdb.Get("key").Result())

}

func initRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:        "127.0.0.1:6379",
		Password:    "", // no password set
		DB:          0,  // use default DB
		PoolSize:    10,
		MaxRetries:  3,
		IdleTimeout: 10 * time.Second,
	})
	pong, err := rdb.Ping().Result()
	if err != nil {
		return err
	} else {
		fmt.Println(pong)
	}
	return
}
