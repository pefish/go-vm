package go_vm

import (
	"errors"
	"fmt"
)

type StackFrame struct {
	values     []*Value // 存放栈帧中所有数据
	retAddress int64    // 存放返回地址
}

func (s StackFrame) String() string {
	valuesStr := ""
	for _, value := range s.values {
		valuesStr += fmt.Sprintf("%s", value) + "; "
	}
	return fmt.Sprintf("values: <%s>, retAddress: <%d>", valuesStr, s.retAddress)
}

func NewStackFrame() *StackFrame {
	return &StackFrame{
		values: make([]*Value, 0),
	}
}

// 返回栈帧的大小
func (s *StackFrame) Size() int {
	return len(s.values)
}

// pop出一个值
func (s *StackFrame) Pop() *Value {
	if len(s.values) <= 0 {
		panic(errors.New("insufficient stack space"))
	}
	popValue := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return popValue
}

// push一个值
func (s *StackFrame) Push(value *Value) {
	s.values = append(s.values, value)
}

// 获取栈帧顶部的值
func (s *StackFrame) GetTopValue() *Value {
	return s.values[len(s.values)-1]
}
