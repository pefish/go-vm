package go_vm

import (
	"github.com/pefish/go-test-assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := NewStack()
	stackFrame := NewStackFrame()
	stackFrame.Push(&Value{data: "haha", valueType:ValueType_STRING})
	stackFrame.Push(&Value{data: "haha1", valueType:ValueType_STRING})
	stack.Push(stackFrame)
	currentStackFrame := stack.GetTopStackFrame()
	test.Equal(t, "haha1", currentStackFrame.GetTopValue().data)
	popStackFrame := stack.Pop()
	test.Equal(t, "haha1", popStackFrame.Pop().data)
	test.Equal(t, "haha", popStackFrame.Pop().data)
}