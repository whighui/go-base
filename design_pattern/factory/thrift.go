package factory

type ThriftCollectMeta struct {
}

func (mc *ThriftCollectMeta) GetMetaInfo(psm, psmType string) (bool, error, interface{}) {
	return false, nil, nil
}
func (mc *ThriftCollectMeta) InterPsmSchemaInfo(psm, name string, metaInfo interface{}) error {
	return nil
}

func (mc *ThriftCollectMeta) UpdatePmsInfo(psm string, psmInfo interface{}) error {
	return nil
}

func (mc *ThriftCollectMeta) CollectMeta(t collectType) ICollectMeta {
	switch t {
	case "mysql":
		return new(MysqlCollectMeta)
	}
	return nil
}
