package conn

import (
	"github.com/go-redis/redis"
)

// 定义一个全局变量
var redisclient *redis.Client

func InitRedis() (err error) {
	redisclient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // 指定
		Password: "",
		DB:       0, // redis一共16个库，指定其中一个库即可
	})
	_, err = redisclient.Ping().Result()
	return
}

//下边就是做一些事物呗

//multi  开启事物 多组名列的打包功能

//exec   执行事务

func RedisTransaction() {
	//开启事务 类似于multi
	pipe := redisclient.TxPipeline()
	pipe.Set("name", "whig", 0)
	pipe.Set("age", "18", 0)
	pipe.Set("grade", "750", 0)

	//执行事务，类似于exec
	pipe.Exec()
}
