package factory

import "testing"

func TestMysql(t *testing.T) {
	var psmType string
	switch psmType {
	case "mysql":
		new(MysqlCollectMeta).
			CollectMeta("mysql").
			GetMetaInfo("", "")
	case "http":
		new(MysqlCollectMeta).CollectMeta("mysql").GetMetaInfo("", "")
	case "thrift":
		new(MysqlCollectMeta).CollectMeta("mysql").GetMetaInfo("", "")
	case "unknown": //已经标记为unknown类型  如果后续该Psm又存在类型呢
		return
	default: //针对剩余类型
		new(MultiCollectMeta).CollectMeta("").GetMetaInfo("", "")
		//先判断 mysql
		//在判断 http
		//在判断 thrift
	}
}
