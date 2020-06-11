package go_vm

import (
	"fmt"
	"testing"
)

func ExampleVm_Run() {
	// CONST "haha"
	// PRINT
	// HALT
	vm := NewVm([]*Instruction{
		{
			opCode: CONST,
			args: []*Value{
				&Value{data:"haha", valueType:ValueType_STRING},
			},
		},
		{
			opCode: PRINT,
		},
		{
			opCode: HALT,
		},
	})
	err := vm.Run()
	if err != nil {
		panic(err)
	}
	// OUTPUT:
	// haha
}

func TestNewVmFromText(t *testing.T) {
	vm := NewVmFromText(`
CONST 'Hello World'
PRINT 
halt
`)
	vm.Run()
	// Output:
	// Hello World
}

func TestNewVmFromText1(t *testing.T) {
	vm := NewVmFromText(`
CONST 1
CONST 2
ADD
PRINT 
halt
`)
	vm.Run()
	// Output:
	// 3
}

func TestNewVmFromText2(t *testing.T) {
	vm := NewVmFromText(`
CONST 1.11
CONST 2.22
ADD
PRINT 
halt
`)
	vm.Run()
	// Output:
	// 3.33
}

func ExampleDecompileText() {
	vm := NewVmFromText(`
CONST "Hello World"
PRINT 
halt
`)
	result, _ := vm.DecompileText()
	fmt.Println(result)
	// Output:
	// CONST "Hello World"
	// PRINT
	// HALT
}