package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

var ctx = context.Background()

func makeKey(ns string, id int64, i int) string {
	return fmt.Sprintf("%s-%d-%d", ns, id, i)
}

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Errorf("expected key name in cmd argument e.g: %s key_name", os.Args[0]))
	}
	ns := os.Args[1]
	restArgs := os.Args[2:]

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	incrRes := rdb.Incr(ctx, ns)
	if err := incrRes.Err(); err != nil {
		panic(err)
	}
	id := incrRes.Val()

	for idx, val := range restArgs {
		err := rdb.Set(ctx, makeKey(ns, id, idx), val, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}
