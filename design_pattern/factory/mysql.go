package factory

type MysqlCollectMeta struct {
}

func (mc *MysqlCollectMeta) CollectPmsInfo(psm string) (error, interface{}) {
	return nil, nil
}
func (mc *MysqlCollectMeta) CollectSchemaInfo(psm string, psmType collectType) (error, interface{}) {
	return nil, nil
}
