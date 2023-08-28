package bitmap

import (
	"fmt"
	"golang-base/redis_scene"
	"strconv"
)

/**
场景题：每月签到情况统计
1.签到时间分布
2.首次签到时间
3.连续三天签到时间
4.判断某一天是否打卡
*/

const SingInPrefix = "sing_in"

//bitcount key  返回offset==1的个数
//bitpos key

func SetUserSingByDate(month, day int, userID int64) {
	redis_scene.RedisConnectInit()
	key := SingInPrefix + ":" + strconv.Itoa(month) + ":" + strconv.Itoa(int(userID))
	bit := redis_scene.RedisClient.SetBit(key, int64(day), 1)
	fmt.Println(bit.Val())
}

//JudgeSingByDate 判断某一天是否签到
func JudgeSingByDate(month int, day int64, userID int64) bool {
	redis_scene.RedisConnectInit()
	key := SingInPrefix + ":" + strconv.Itoa(month) + ":" + strconv.Itoa(int(userID))
	intCmd := redis_scene.RedisClient.GetBit(key, day)
	if intCmd.Err() != nil {
		return false
	}
	if intCmd.Val() == 1 {
		return true
	}
	return false
}

//SingTimeDistribute 签到时间分布
func SingTimeDistribute(month string, userID int64) {

}

//FirstSingTime 首次签到时间
func FirstSingTime(month string, userID int64) {

}

//SingContinuesThreeDay 连续三天签到的首次时间
func SingContinuesThreeDay(month string, userID int64) {

	//连续三天签到的首次时间
}
