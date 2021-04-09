package data

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
)

// List function lists all the tasks in the list
func List(ctx context.Context, rdb *redis.Client) {
	list := rdb.LRange(ctx, "tasks", 0, -1)
	val := list.Val()

	for _, t := range val {
		s := strings.TrimSpace(t)
		fmt.Println("☑️", s)
	}
}

// Add function adds a new task into the list
func Add(ctx context.Context, rdb *redis.Client, t string) {
	rdb.LPushX(ctx, "tasks", t)
	fmt.Println("☑️ : The task was added")
}

// Add function adds a new task into the list
// func Done(ctx context.Context, rdb *redis.Client, t string) {
// 	rdb.LPushX(ctx, "tasks", t)
// 	fmt.Println("☑️ : The task was added")
// }
