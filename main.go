package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"milkteapool.com/pool-server/api"
	"milkteapool.com/pool-server/pool"
	"milkteapool.com/pool-server/utils"
)

func main() {

	config, err := NewConfig("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Setup Redis
	var (
		host     = utils.GetEnv("REDIS_HOST", "localhost")
		port     = string(utils.GetEnv("REDIS_PORT", "6379"))
		password = utils.GetEnv("REDIS_PASSWORD", "")
	)

	cache := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})

	_, err = cache.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	// Init pool
	pool := pool.NewPool(config)
	pool.Start()

	// Setup router
	r := gin.Default()
	// api := &api.Controller{pool: pool}
	api := &api.Controller{Pool: pool}

	r.GET("/pool_info", api.PoolInfo)
	r.POST("/submit_partial", api.SubmitPartial)

	// Start server
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
