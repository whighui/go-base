package factory

type HttpCollectMeta struct {
}

func (mc *HttpCollectMeta) GetMetaInfo(psm, psmType string) (bool, error, interface{}) {
	return false, nil, nil
}
func (mc *HttpCollectMeta) InterPsmSchemaInfo(psm, name string, metaInfo interface{}) error {
	return nil
}

func (mc *HttpCollectMeta) UpdatePmsInfo(psm string, psmInfo interface{}) error {
	return nil
}

func (mc *HttpCollectMeta) CollectMeta(t collectType) ICollectMeta {
	switch t {
	case "mysql":
		return new(MysqlCollectMeta)
	}
	return nil
}
