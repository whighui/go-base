package factory

type MultiCollectMeta struct {
	mysqlCollect  *MysqlCollectMeta
	httpCollect   *HttpCollectMeta
	thriftCollect *ThriftCollectMeta
}

func (mc *MultiCollectMeta) CollectPmsInfo(psm string) (error, interface{}) {
	return nil, nil
}

func (mc *MultiCollectMeta) CollectSchemaInfo(psm string, psmType collectType) (error, interface{}) {
	return nil, nil
}
