package go_vm

import (
	"errors"
	"fmt"
	"github.com/pefish/go-decimal"
	"reflect"
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
		//fmt.Printf("%v\n", instruction.opCode)
		currentStackFrame := vm.stack.GetTopStackFrame()
		//fmt.Printf("currentStackFrame: %v\n", currentStackFrame)
		switch instruction.opCode {
		case ADD:
			v1 := currentStackFrame.Pop().data
			v2 := currentStackFrame.Pop().data
			float64_, err := strconv.ParseFloat(go_decimal.Decimal.Start(v1).AddForString(v2), 64)
			if err != nil {
				return err
			}
			currentStackFrame.Push(&Value{data: float64_, valueType: ValueType_NUMBER})
		case SUB:
			v1 := currentStackFrame.Pop().data
			v2 := currentStackFrame.Pop().data
			float64_, err := strconv.ParseFloat(go_decimal.Decimal.Start(v1).SubForString(v2), 64)
			if err != nil {
				return err
			}
			currentStackFrame.Push(&Value{data: float64_, valueType: ValueType_NUMBER})
		case MUL:
			v1 := currentStackFrame.Pop().data
			v2 := currentStackFrame.Pop().data
			float64_, err := strconv.ParseFloat(go_decimal.Decimal.Start(v1).MultiForString(v2), 64)
			if err != nil {
				return err
			}
			currentStackFrame.Push(&Value{data: float64_, valueType: ValueType_NUMBER})
		case DIV:
			v1 := currentStackFrame.Pop().data
			v2 := currentStackFrame.Pop().data
			float64_, err := strconv.ParseFloat(go_decimal.Decimal.Start(v1).DivForString(v2), 64)
			if err != nil {
				return err
			}
			currentStackFrame.Push(&Value{data: float64_, valueType: ValueType_NUMBER})
		case CONST:
			if len(instruction.args) < 1 {
				panic(fmt.Errorf("instruction error - %v", instruction))
			}
			currentStackFrame.Push(instruction.args[0])
		case JMP:
			if len(instruction.args) < 1 {
				panic(fmt.Errorf("instruction error - %v", instruction))
			}
			targetPos, err := instruction.args[0].GetNumber()
			if err != nil {
				panic(err)
			}
			vm.instructionPointer = int64(targetPos) - 1
		case JNE:
			v1 := currentStackFrame.Pop().data
			v2 := currentStackFrame.Pop().data
			if !reflect.DeepEqual(v1, v2) {
				if len(instruction.args) < 1 {
					panic(fmt.Errorf("instruction error - %v", instruction))
				}
				targetPos, err := instruction.args[0].GetNumber()
				if err != nil {
					panic(err)
				}
				vm.instructionPointer = int64(targetPos) - 1
			}
		case JEQ:
			v1 := currentStackFrame.Pop().data
			v2 := currentStackFrame.Pop().data
			if reflect.DeepEqual(v1, v2) {
				if len(instruction.args) < 1 {
					panic(fmt.Errorf("instruction error - %v", instruction))
				}
				targetPos, err := instruction.args[0].GetNumber()
				if err != nil {
					panic(err)
				}
				vm.instructionPointer = int64(targetPos) - 1
			}
		case PRINT:
			fmt.Println(currentStackFrame.Pop().data)
		case CALL:
			if len(instruction.args) < 2 {
				panic(fmt.Errorf("instruction error - %v", instruction))
			}
			targetPos, err := instruction.args[0].GetNumber()  // 取出跳转到的位置
			if err != nil {
				return err
			}
			argsCount, err := instruction.args[1].GetNumber()  // 取出参数个数
			if err != nil {
				return err
			}
			newStackFrame := NewStackFrame()  // 新建一个栈帧
			for argsCount > 0 {
				newStackFrame.Push(currentStackFrame.Pop())  // 参数移动过来
				argsCount--
			}
			vm.stack.Push(newStackFrame)  // 放入栈帧

			currentStackFrame.retAddress = vm.instructionPointer  // 记录返回地址
			vm.instructionPointer = int64(targetPos) - 1  // jmp
		case RET:
			if len(instruction.args) < 1 {
				panic(fmt.Errorf("instruction error - %v", instruction))
			}
			returnCount, err := instruction.args[0].GetNumber()
			if err != nil {
				return err
			}
			lastStackFrame := vm.stack.GetLastStackFrame()
			for returnCount > 0 {
				lastStackFrame.Push(currentStackFrame.Pop())
				returnCount--
			}
			vm.stack.Pop()

			vm.instructionPointer = lastStackFrame.retAddress
		case HALT:
			return nil
		}
		vm.stepInstruction()
	}
}
