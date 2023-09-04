package factory

type HttpCollectMeta struct {
}

func (mc *HttpCollectMeta) CollectPmsInfo(psm string) (error, interface{}) {
	return nil, nil
}

func (mc *HttpCollectMeta) CollectSchemaInfo(psm string, psmType collectType) (error, interface{}) {
	return nil, nil

}
