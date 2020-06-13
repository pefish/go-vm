package go_vm

import (
	"fmt"
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

func ExampleNewVmFromText() {
	vm := NewVmFromText(`
CONST 'Hello World'
PRINT 
halt
`)
	err := vm.Run()
	if err != nil {
		panic(err)
	}
	// Output:
	// Hello World
}

func ExampleNewVmFromText1() {
	vm := NewVmFromText(`
CONST 1
CONST 2
ADD
PRINT 
halt
`)
	err := vm.Run()
	if err != nil {
		panic(err)
	}
	// Output:
	// 3
}

func ExampleNewVmFromText2() {
	vm := NewVmFromText(`
CONST 1.11
CONST 2.22
ADD
PRINT 
halt
`)
	err := vm.Run()
	if err != nil {
		panic(err)
	}
	// Output:
	// 3.33
}

func ExampleNewVmFromText3() {
	vm := NewVmFromText(`
CONST 1.11
CONST 2.22
SUB
PRINT 
halt
`)
	err := vm.Run()
	if err != nil {
		panic(err)
	}
	// Output:
	// 1.11
}

func ExampleNewVmFromText4() {
	vm := NewVmFromText(`
CONST 1.11
CONST 2.22
MUL
PRINT 
halt
`)
	err := vm.Run()
	if err != nil {
		panic(err)
	}
	// Output:
	// 2.4642
}

func ExampleNewVmFromText5() {
	vm := NewVmFromText(`
CONST 1.11
CONST 2.22
DIV
PRINT 
halt
`)
	err := vm.Run()
	if err != nil {
		panic(err)
	}
	// Output:
	// 2
}

func ExampleNewVmFromText6() {
	vm := NewVmFromText(`
CONST 3.33
CONST 3.33
JNE 6
CONST "相等"
PRINT
halt
CONST "不相等"
PRINT
halt
`)
	err := vm.Run()
	if err != nil {
		panic(err)
	}
	// Output:
	// 相等
}

func ExampleNewVmFromText7() {
	vm := NewVmFromText(`
CONST 3.33
CONST 3.334
JNE 6
CONST "相等"
PRINT
halt
CONST "不相等"
PRINT
halt
`)
	err := vm.Run()
	if err != nil {
		panic(err)
	}
	// Output:
	// 不相等
}

func ExampleNewVmFromText8() {
	vm := NewVmFromText(`
CONST "3.33"
CONST "3.31"
JNE 6
CONST "相等"
PRINT
halt
CONST "不相等"
PRINT
halt
`)
	err := vm.Run()
	if err != nil {
		panic(err)
	}
	// Output:
	// 不相等
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