package main

import (
	"context"
	"redis/common"
	"redis/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	// Connect to your redis server
	err := common.RedisClient.ConnectRedis(context.TODO())
	if err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()
	r.POST("/setdata", controllers.SetDataToRedisCache)
	r.GET("/getdata", controllers.GetDataRedis)
	r.Run(":9080")
}
