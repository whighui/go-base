package redis_lock

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

//https://mp.weixin.qq.com/s/KYiZvFRX0CddJVCwyfkLfQ

/**
第一版分布式锁主要实现方式： 分布式锁是对多进程下对临界共享资源访问的一种保护措施

实现分布式锁主要有 redis、mysql、etcd、zookper 前两种类似于cas+自旋方式来实现  后两种带有回调通知机制

1.利用redis set key value ex 过期时间 NX  (ex表示秒过期时间、px表示毫秒过期时间)


*/

//GetCurrentProcessID 获取当前进程ID
func GetCurrentProcessID() string {
	return strconv.Itoa(os.Getpid())
}

// GetCurrentGoroutineID 获取当前的协程ID
func GetCurrentGoroutineID() string {
	buf := make([]byte, 128)
	buf = buf[:runtime.Stack(buf, false)]
	stackInfo := string(buf)
	return strings.TrimSpace(strings.Split(strings.Split(stackInfo, "[running]")[0], "goroutine")[1])
}

func GetProcessAndGoroutineIDStr() string {
	return fmt.Sprintf("%s_%s", GetCurrentProcessID(), GetCurrentGoroutineID())
}
