package redis_scene

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func RedisConnectInit() {
	// 创建Redis客户端
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址和端口
		Password: "",               // Redis密码，如果没有设置密码则为空字符串
	})
	// 使用Ping方法测试与Redis的连接
	_, err := RedisClient.Ping().Result()
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		return
	}
}
