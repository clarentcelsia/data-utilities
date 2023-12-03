package controllers

import (
	"encoding/json"
	"redis/common"

	"github.com/gin-gonic/gin"
)

func SetDataToRedisCache(c *gin.Context) {
	val := map[string]interface{}{
		"Data":   1,
		"Object": "Hello World",
	}
	val_json, _ := json.Marshal(val)
	if err := common.RedisClient.SetDataToRedis(c, "data1", val_json); err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, "succeed")
}

func GetDataRedis(c *gin.Context) {
	result, err := common.RedisClient.GetDataFromRedis(c, "data1")
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, result)
}
