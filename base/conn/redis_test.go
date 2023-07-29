package conn

import (
	"fmt"
	"testing"
)

func TestStringMaxLen(t *testing.T) {

	err := InitRedis()

	fmt.Println(err)

	bytes := make([]byte, 0, 1024*512)
	//这里边设置512MB
	for i := 0; i < 1024*1024*512-1; i++ {
		bytes = append(bytes, 'b')
	}
	//将字符数组转化为strig
	str := string(bytes)
	set := redisclient.Append("key", str)
	fmt.Println(set)
}
