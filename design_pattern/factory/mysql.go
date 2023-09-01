package factory

type MysqlCollectMeta struct {
}

func (mc *MysqlCollectMeta) GetMetaInfo(psm, psmType string) (bool, error, interface{}) {
	return false, nil, nil
}
func (mc *MysqlCollectMeta) InterPsmSchemaInfo(psm, name string, metaInfo interface{}) error {
	return nil
}

func (mc *MysqlCollectMeta) UpdatePmsInfo(psm string, psmInfo interface{}) error {
	return nil
}
