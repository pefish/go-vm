package go_vm

type Stack struct {
	stackFrames []*StackFrame // 存放所有栈帧
}

func NewStack() *Stack {
	stackFrames := make([]*StackFrame, 0)
	return &Stack{
		stackFrames: append(stackFrames, NewStackFrame()),  // 有一个初始栈帧。后面每个新增一个栈帧
	}
}

// 栈帧入栈
func (s *Stack) Push(stackFrame *StackFrame) {
	s.stackFrames = append(s.stackFrames, stackFrame)
}

// 栈帧出栈
func (s *Stack) Pop() *StackFrame {
	stackFrame := s.stackFrames[len(s.stackFrames) - 1]
	s.stackFrames = s.stackFrames[:len(s.stackFrames) - 1]
	return stackFrame
}

func (s *Stack) GetTopStackFrame() *StackFrame {
	return s.stackFrames[len(s.stackFrames) - 1]
}
