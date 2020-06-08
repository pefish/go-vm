package go_vm

import (
	"github.com/pefish/go-test-assert"
	"testing"
)

func TestStackFrame_Pop(t *testing.T) {
	stackFrame := NewStackFrame()
	stackFrame.Push(&Value{data: 1, valueType: ValueType_INT64})
	test.Equal(t, 1, stackFrame.Size())
	result := stackFrame.Pop()
	test.Equal(t, 1, result.data)
	test.Equal(t, 0, stackFrame.Size())
}
