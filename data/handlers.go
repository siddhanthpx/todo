package data

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
)

// List function lists all the tasks in the list
func List(ctx context.Context, rdb *redis.Client, filterDone bool) {
	if filterDone {
		list := rdb.LRange(ctx, "done", 0, -1)
		val := list.Val()

		for _, t := range val {
			s := strings.TrimSpace(t)
			fmt.Println("☑️", s)
		}

	} else {
		list := rdb.LRange(ctx, "tasks", 0, -1)
		val := list.Val()

		for _, t := range val {
			s := strings.TrimSpace(t)
			fmt.Println("☑️", s)
		}
	}
}

// Add function adds a new task into the list
func Add(ctx context.Context, rdb *redis.Client, t string) {
	rdb.LPushX(ctx, "tasks", t)
	fmt.Println("☑️ : The task was added")
}

// Done function marks the given task for completion and moves it into a new list called "done"
func Done(ctx context.Context, rdb *redis.Client, t string) {
	rdb.LRem(ctx, "tasks", 0, t)
	rdb.LPush(ctx, "done", t)
	fmt.Println("☑️ : The task was completed")
}
