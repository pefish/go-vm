package go_vm

type Value struct {
	data      interface{}
	valueType ValueType
}

type ValueType int

const (
	_ ValueType = iota
	ValueType_STRING
	ValueType_NUMBER // 对应Golang中的float64
)

func (v *Value) GetString() (string, error) {
	if v.valueType != ValueType_STRING {
		return "", nil
	}
	return v.data.(string), nil
}

func (v *Value) GetNumber() (float64, error) {
	if v.valueType != ValueType_NUMBER {
		return 0, nil
	}
	return v.data.(float64), nil
}
