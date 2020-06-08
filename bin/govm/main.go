package main

import (
	go_vm "github.com/pefish/go-vm"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("args error")
	}
	vm := go_vm.NewVmFromText(os.Args[1])
	vm.Run()
}
