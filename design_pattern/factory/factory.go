package factory

import (
	"strings"
	"sync"
)

var psmInfoChannelLength = 100
var psmInfoChannel chan string
var collectMetaProcess *CollectMetaProcess

type collectType string

func init() {
	psmInfoChannel = make(chan string, psmInfoChannelLength)
	collectMetaProcess = &CollectMetaProcess{
		mutex: &sync.Mutex{},
		wg:    &sync.WaitGroup{},
	}
}

// 处理已知类型信息
type ICollectMetaFactory interface {
	CollectMeta(t collectType) ICollectMeta
}

type ICollectMeta interface {
	GetMetaInfo(psm, psmType string) (bool, error, interface{}) //判断获取http是否存在  判断获取rpc是否存在 判断获取mysql是否存在被
	InterPsmSchemaInfo(psm, name string, metaInfo interface{}) error
	UpdatePmsInfo(psm string, psmInfo interface{}) error
}

type CollectMetaProcess struct {
	mutex *sync.Mutex
	wg    *sync.WaitGroup
}

func (cmp *CollectMetaProcess) CollectMeta(t collectType) ICollectMeta {
	switch t {
	case "mysql":
		return new(MysqlCollectMeta)
	case "http":
		return new(HttpCollectMeta)
	case "thrift":
		return new(ThriftCollectMeta)
	case "unknown":
		return nil
	case "":
		return new(MultiCollectMeta)
	}
	return nil
}

// 工厂对外统一处理消息来源呗  psmInfo=psm+psmType
func (cmp *CollectMetaProcess) metaProcess(psmInfoChannel chan string) error {
	for {
		psmInfo := <-psmInfoChannel
		psm, psmType := strings.Split(psmInfo, "::")[0], strings.Split(psmInfo, "::")[1]
		collectHandler := cmp.CollectMeta(collectType(psmType))
		cmp.wg.Add(1)
		//这里边可以改进使用线程池来并发处理呗
		go func() {
			defer cmp.wg.Done()
			collectHandler.GetMetaInfo(psm, psmType)
		}()
	}
}
