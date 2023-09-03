package factory

import (
	"strings"
	"sync"
)

const psmInfoChannelLength = 100

var collectMetaProcess *CollectMetaProcess

type collectType string

const (
	MysqlType   collectType = "mysql"
	HttpType    collectType = "http"
	ThriftType  collectType = "thrift"
	UnknownType collectType = "unknown"
)

func init() {
	collectMetaProcess = &CollectMetaProcess{
		mutex:          &sync.Mutex{},
		wg:             &sync.WaitGroup{},
		psmInfoChannel: make(chan string, psmInfoChannelLength),
	}
	collectMetaProcess.start()
}

type ICollectMetaFactory interface {
	CollectMeta(t collectType) ICollectMeta
}

type ICollectMeta interface {
	CollectPmsInfo(psm string) (error, interface{}) //补充Psm信息
	CollectSchemaInfo(psm string, psmType collectType) (error, interface{})
}

type CollectMetaProcess struct {
	mutex          *sync.Mutex
	wg             *sync.WaitGroup
	psmInfoChannel chan string
}

func (cmp *CollectMetaProcess) CollectMeta(t collectType) ICollectMeta {
	switch t {
	case MysqlType:
		return new(MysqlCollectMeta)
	case HttpType:
		return new(HttpCollectMeta)
	case ThriftType:
		return new(ThriftCollectMeta)
	default: //走混合刷新Meta信息逻辑
		return new(MultiCollectMeta)
	}
}

// 工厂对外统一处理消息来源呗  psmInfo=psm+psmType
func (cmp *CollectMetaProcess) receiveTask(psmInfo string) {
	cmp.psmInfoChannel <- psmInfo
}

func (cmp *CollectMetaProcess) start() {
	go func() {
		for {
			psmInfo := <-cmp.psmInfoChannel
			psm, psmType := strings.Split(psmInfo, "::")[0], strings.Split(psmInfo, "::")[1]
			cmp.wg.Add(1)
			go cmp.process(psm, collectType(psmType)) //处理转成这种数据类型呗
		}
	}()
}

func (cmp *CollectMetaProcess) process(psm string, psmType collectType) {
	defer cmp.wg.Done() //这里无论如何一定要进行释放wg.Done 否则会造成内存溢出呗

	cmp.CollectMeta(psmType).CollectPmsInfo(psm) // 是否需要进行重试机制 RPC获取overpass有限流机制 每个类型是否需要进行限流 到时候讨论一下

	cmp.CollectMeta(psmType).CollectSchemaInfo(psm, psmType) //补充 schema 信息
}
