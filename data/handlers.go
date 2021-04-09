package data

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
)

func List(ctx context.Context, rdb *redis.Client) {
	list := rdb.LRange(ctx, "tasks", 0, -1)
	val := list.Val()

	for idx, t := range val {
		strings.TrimSpace(t)
	}
	fmt.Println()
}
