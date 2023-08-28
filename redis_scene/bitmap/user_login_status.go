package bitmap

import (
	"fmt"
	"golang-base/redis_scene"
)

//利用redis bitmap来判断用户登录状态

/**
选型为什么选redis bitmap 来判断用户登录状态呢？ 为什么不选择redis sds来做判断呢

首先如果选sds动态字符串作为用户登录判断可以 set username 1  这种方式来实现证明用户登录状态
但是会存在如下缺点：
1.简单动态字符串有TLV T 类型 L 长度 V 数据  其中都是占据内存的，，还有一个 RedisObject 结构的开销，因为 Redis 的数据类型有很多，而且，不同数据类型都有些相同的元数据要记录（比如最后一次访问的时间、被引用的次数等）。
所以，Redis 会用一个 RedisObject 结构体来统一记录这些元数据，同时指向实际数据。
为什么会选择bitmap，因为bitmap底层是由sds数据结构来表示的：利用底层byte数组 每个byte用8位 0 或 1 来表示  //但是有最大长度限制 2^32-1=512MB
*/

// SETBIT  <key> <offset> <value>
//例如     login_status user_id 1   设置key=login_status  user_id=10086 status=1 代表user_id登录状态信息

//GETBIT <key> <offset>            GETBIT login_status 10086   获取登录状态 offset=10086的登录状态

const (
	LoginStatus = "login_status"
)

//SetUserLogin 设置用户登录
//uint32 最大值 4294967295
//uint64 最大值 18446744073709551615
//             241079929031545845   字节user_id 基本上都是以2开头
func SetUserLogin(userID int64) {
	redis_scene.RedisConnectInit()
	bit := redis_scene.RedisClient.GetBit(LoginStatus, userID)
	if bit.Val() == '1' {
		fmt.Println("用户已登录，无需再设置登录")
		return
	}
	setBit := redis_scene.RedisClient.SetBit(LoginStatus, userID, 1)
	fmt.Println(setBit.Val())
}

//SetUserLoginOut 设置用户退出
func SetUserLoginOut(userID int64) {
	redis_scene.RedisConnectInit()
	bit := redis_scene.RedisClient.GetBit(LoginStatus, userID)
	if bit.Err() == nil && bit.Val() == '0' {
		return
	}
	result := redis_scene.RedisClient.SetBit(LoginStatus, userID, 1)
	if result.Err() != nil {
		fmt.Println(result.Err().Error())
	}
}

//IsUserLogin 判断用户是否在登录当中
func IsUserLogin(userID int64) bool {
	bit := redis_scene.RedisClient.GetBit(LoginStatus, userID)
	if bit.Val() == '1' {
		return true
	}
	return false
}
