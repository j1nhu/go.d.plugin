package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"runtime/debug"
	"strings"
)

type Model struct {
	Str1    string   `redis:"str1"`
	Str2    string   `redis:"str2"`
	Int     int      `redis:"int"`
	Bool    bool     `redis:"bool"`
	Ignored struct{} `redis:"-"`
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	var cursor uint64
	var n int

	for {
		var keys []string
		var err error
		keys, cursor, err = rdb.Scan(ctx, cursor, "queues:*", 2).Result()
		if err != nil {
			panic(err)
		}

		if len(keys) > 0 {
			n += len(keys)

			for _, key := range keys {
				fmt.Println("==== looping ====", key)
				colonCount := strings.Count(key, ":")
				if colonCount > 1 {
					lastInd := strings.LastIndex(key, ":")
					key = key[:lastInd]
				}
				queueLength, queueLengthErr := rdb.LLen(ctx, key).Result()
				if queueLengthErr != nil {
					//log.Println(err)
					panic(queueLengthErr)
				}
				fmt.Println(key, "key length", queueLength)
				//fmt.Println(key, queueLength)
			}
			//fmt.Printf("found keys %s \n", keys)
		}

		if cursor == 0 {
			break
		}
	}

	fmt.Printf("found %d keys\n", n)
}
