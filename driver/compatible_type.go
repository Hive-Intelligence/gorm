package driver

type DefaultCompatibleType struct {
}

func (d DefaultCompatibleType) CompatibleType(dataType string) string {
	return dataType
}
