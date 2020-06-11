package go_vm

import (
	"errors"
	"fmt"
	"strconv"
)

type Vm struct {
	instructions       []*Instruction // 输入的指令集。代码段
	stack              *Stack         // 栈
	instructionPointer int64          // 指令运行位置（instructions数组的索引）
	heap               []interface{}  // 堆区
}

func NewVm(instructions []*Instruction) *Vm {
	return &Vm{
		instructions:       instructions,
		stack:              NewStack(),
		instructionPointer: 0,
		heap:               make([]interface{}, 0),
	}
}

func NewVmFromText(str string) *Vm {
	lexer := NewLexer(str)
	tokens := lexer.ParseTokens()
	return &Vm{
		instructions:       GetInstructionsFromTokens(tokens),
		stack:              NewStack(),
		instructionPointer: 0,
		heap:               make([]interface{}, 0),
	}
}

func GetInstructionsFromTokens(tokens []Token) []*Instruction {
	instructions := make([]*Instruction, 0)
	for _, token := range tokens {
		if token.Type == TokenType_STRING {
			lastInstruction := instructions[len(instructions)-1]
			dataType, ok := TokenTypeToDataType[token.Type]
			if !ok {
				panic(fmt.Errorf("value type not found - token type: %d", token.Type))
			}
			lastInstruction.args = append(lastInstruction.args, &Value{data: token.Literal, valueType: dataType})
			continue
		}
		if token.Type == TokenType_NUMBER {
			lastInstruction := instructions[len(instructions)-1]
			dataType, ok := TokenTypeToDataType[token.Type]
			if !ok {
				panic(fmt.Errorf("value type not found - token type: %d", token.Type))
			}
			float64_, err := strconv.ParseFloat(token.Literal, 64)
			if err != nil {
				panic(err)
			}
			lastInstruction.args = append(lastInstruction.args, &Value{data: float64_, valueType: dataType})
			continue
		}
		opCode, ok := TokenTypeToOpCode[token.Type]
		if !ok {
			continue
		}
		instruction := Instruction{
			opCode: opCode,
			args:   make([]*Value, 0),
		}
		instructions = append(instructions, &instruction)
	}
	return instructions
}

func (vm *Vm) DecompileText() (string, error) {
	if len(vm.instructions) <= 0 {
		return ``, errors.New("no instructions")
	}
	result := ""
	for _, instruction := range vm.instructions {
		tokenType, ok := OpCodeToTokenType[instruction.opCode]
		if !ok {
			continue
		}
		str, ok := TokenTypeToString[tokenType]
		if !ok {
			continue
		}
		for _, arg := range instruction.args {
			if arg.valueType == ValueType_STRING {
				str += " " + `"` + fmt.Sprint(arg.data) + `"`
			} else {
				str += " " + fmt.Sprint(arg.data)
			}
		}
		result += str + "\n"
	}
	return result, nil
}

func (vm *Vm) fetchInstruction() *Instruction {
	return vm.instructions[vm.instructionPointer]
}

func (vm *Vm) stepInstruction() {
	vm.instructionPointer++
}

func (vm *Vm) Run() error {
	for {
		instruction := vm.fetchInstruction()
		switch instruction.opCode {
		case ADD:
			currentStackFrame := vm.stack.GetTopStackFrame()
			v1, err := currentStackFrame.Pop().GetNumber()
			if err != nil {
				return err
			}
			v2, err := currentStackFrame.Pop().GetNumber()
			if err != nil {
				return err
			}
			currentStackFrame.Push(&Value{data: v1 + v2, valueType: ValueType_NUMBER})
		case CONST:
			currentStackFrame := vm.stack.GetTopStackFrame()
			if len(instruction.args) < 1 {
				panic(fmt.Errorf("instruction error - %v", instruction))
			}
			currentStackFrame.Push(instruction.args[0])
		case PRINT:
			currentStackFrame := vm.stack.GetTopStackFrame()
			fmt.Println(currentStackFrame.Pop().data)
		case HALT:
			return nil
		}
		vm.stepInstruction()
	}
}
