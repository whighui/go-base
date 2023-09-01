package factory

type ThriftCollectMeta struct {
}

func (mc *ThriftCollectMeta) CollectPmsInfo(psm string) (error, interface{}) {
	return nil, nil
}
func (mc *ThriftCollectMeta) CollectSchemaInfo(psm string, psmType collectType) (error, interface{}) {
	return nil, nil
}
