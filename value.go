package go_vm

type Value struct {
	data      interface{}
	valueType ValueType
}

type ValueType int

const (
	_ ValueType = iota
	ValueType_STRING
	ValueType_INT64
	ValueType_FLOAT64
)
