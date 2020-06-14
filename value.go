package go_vm

import "fmt"

type Value struct {
	data      interface{}
	valueType ValueType
}

func (v Value) String() string {
	return fmt.Sprintf("data: <%v>, valueType: <%d>", v.data, v.valueType)
}

type ValueType int

const (
	_ ValueType = iota
	ValueType_STRING
	ValueType_NUMBER // 对应Golang中的float64
)

func (v *Value) GetString() (string, error) {
	if v.valueType != ValueType_STRING {
		return "", fmt.Errorf("type error - %s", v)
	}
	return v.data.(string), nil
}

func (v *Value) GetNumber() (float64, error) {
	if v.valueType != ValueType_NUMBER {
		return 0, fmt.Errorf("type error - %s", v)
	}
	return v.data.(float64), nil
}
