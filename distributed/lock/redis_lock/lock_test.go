package redis_lock

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func TestConn(t *testing.T) {
	// 建立Redis连接
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		return
	}
	defer conn.Close()

	// 执行Redis命令
	_, err = conn.Do("SET", "key", "value")
	if err != nil {
		fmt.Println("Failed to set key:", err)
		return
	}

	value, err := redis.String(conn.Do("GET", "key"))
	if err != nil {
		fmt.Println("Failed to get value:", err)
		return
	}

	fmt.Println("Value:", value)
}
