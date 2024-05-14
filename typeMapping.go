package goject

type TypeMapping struct {
	Index      int
	TypeName   string
	ServiceKey string
}

func NewTypeMapping(index int, typeName, serviceKey string) *TypeMapping {
	return &TypeMapping{
		Index:      index,
		TypeName:   typeName,
		ServiceKey: serviceKey,
	}
}
