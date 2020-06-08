package go_vm

import (
	"fmt"
	"testing"
)

func ExampleVm_Run() {
	// CONSTQ "haha"
	// PRINT
	// HALT
	vm := NewVm([]*Instruction{
		{
			opCode: CONSTQ,
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
	vm.Run()

	// OUTPUT:
	// haha
}

func TestNewVmFromText(t *testing.T) {
	vm := NewVmFromText(`
CONSTQ 'Hello World'
PRINT 
halt
`)
	vm.Run()
	// Output:
	// Hello World
}

func ExampleDecompileText() {
	vm := NewVmFromText(`
CONSTQ "Hello World"
PRINT 
halt
`)
	result, _ := vm.DecompileText()
	fmt.Println(result)
	// Output:
	// CONSTQ "Hello World"
	// PRINT
	// HALT
}