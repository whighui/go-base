package factory

import "fmt"

type MultiCollectMeta struct {
	mysqlCollect  *MysqlCollectMeta
	httpCollect   *HttpCollectMeta
	thriftCollect *ThriftCollectMeta
}

func (mc *MultiCollectMeta) GetMetaInfo(psm, psmType string) (bool, error, interface{}) {

	if isSuccess, err, metaInfo := mc.mysqlCollect.GetMetaInfo(psm, psmType); err != nil {
		//这里边就是证明http收集到了呢
		fmt.Println(isSuccess)
		fmt.Println(metaInfo)
	}
	//这里边应该实现三种方式呗
	return false, nil, nil
}
func (mc *MultiCollectMeta) InterPsmSchemaInfo(psm, name string, metaInfo interface{}) error {
	return nil
}

func (mc *MultiCollectMeta) UpdatePmsInfo(psm string, psmInfo interface{}) error {
	return nil
}
