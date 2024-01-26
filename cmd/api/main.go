package main

import (
	"context"
	"net/http"
	"time"

	"github.com/bsm/redislock"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "some-redis:6379",
	})
	defer client.Close()

	// Create a new lock client.
	locker := redislock.New(client)

	r := gin.Default()
	r.GET("/process-job", func(c *gin.Context) {
		key := "process-lock"
		ctx := context.Background()
		ttl := 5 * time.Second
		lock, err := locker.Obtain(ctx, key, ttl, nil)
		if err == redislock.ErrNotObtained {
			c.JSON(http.StatusOK, gin.H{"message": "cannot process because still locking"})
			return

		} else if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "something went wrong"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "working job"})
		go func() {
			defer lock.Release(ctx)
			time.Sleep(120 * time.Second)
		}()

	})
	r.Run(":8080")

}
