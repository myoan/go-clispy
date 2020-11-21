package main

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyStack = errors.New("Stack is empty")
	ErrInvalidSP  = errors.New("Invalid stack pointer")
)

type VirtualMachine struct {
	sp    int
	pc    int
	is    *InstructionSet
	stack []int
}

func NewVirtualMachine() *VirtualMachine {
	return &VirtualMachine{
		sp:    0,
		pc:    0,
		is:    nil,
		stack: make([]int, 0),
	}
}

func (vm *VirtualMachine) Show() {
	fmt.Println("--- start stack ---")
	for i, s := range vm.stack {
		fmt.Printf("[%d] %d\n", i, s)
		if i == vm.sp-1 {
			fmt.Println("- sp")
		}
	}
	fmt.Println("--- end stack ---")
}

func (vm *VirtualMachine) Push(v int) {
	if vm.sp < len(vm.stack) {
		vm.stack[vm.sp] = v
	} else {
		vm.stack = append(vm.stack, v)
	}
	vm.sp += 1
}

func (vm *VirtualMachine) Pop() (int, error) {
	if len(vm.stack) == 0 {
		return -1, ErrEmptyStack
	}
	if vm.sp <= 0 {
		return -1, ErrInvalidSP
	}
	ret := vm.stack[vm.sp-1]
	vm.sp -= 1
	return ret, nil
}

func (vm *VirtualMachine) IncrPC() {
	vm.pc += 1
}

func (vm *VirtualMachine) IsFinish() bool {
	if vm.pc >= len(vm.is.insts) {
		return true
	}
	return false
}

func (vm *VirtualMachine) NextInstruction() *Instruction {
	return vm.is.insts[vm.pc]
}

func (vm *VirtualMachine) Exec(is *InstructionSet) {
	vm.is = is
	for !vm.IsFinish() {
		inst := vm.NextInstruction()
		vm.Show()
		switch inst.iType {
		case InsPush:
			fmt.Printf("push %d\n", inst.value1)
			vm.Push(inst.value1)
		case InsAdd:
			b, err := vm.Pop()
			if err != nil {
				panic(err)
			}
			a, err := vm.Pop()
			if err != nil {
				panic(err)
			}
			fmt.Printf("add %d, %d\n", a, b)
			vm.Push(a + b)
		case InsSub:
			b, _ := vm.Pop()
			a, _ := vm.Pop()
			fmt.Printf("sub %d, %d\n", a, b)
			vm.Push(a - b)
		case InsMul:
			b, _ := vm.Pop()
			a, _ := vm.Pop()
			fmt.Printf("mul %d, %d\n", a, b)
			vm.Push(a * b)
		case InsDiv:
			b, _ := vm.Pop()
			a, _ := vm.Pop()
			fmt.Printf("div %d, %d\n", a, b)
			vm.Push(a / b)
		case InsLt:
			b, _ := vm.Pop()
			a, _ := vm.Pop()
			fmt.Printf("lt  %d, %d\n", a, b)
			if a < b {
				vm.Push(1)
			} else {
				vm.Push(0)
			}
		case InsLte:
			b, _ := vm.Pop()
			a, _ := vm.Pop()
			fmt.Printf("lte %d, %d\n", a, b)
			if a <= b {
				vm.Push(1)
			} else {
				vm.Push(0)
			}
		case InsGt:
			b, _ := vm.Pop()
			a, _ := vm.Pop()
			fmt.Printf("gt  %d, %d\n", a, b)
			if a > b {
				vm.Push(1)
			} else {
				vm.Push(0)
			}
		case InsGte:
			b, _ := vm.Pop()
			a, _ := vm.Pop()
			fmt.Printf("gte %d, %d\n", a, b)
			if a >= b {
				vm.Push(1)
			} else {
				vm.Push(0)
			}
		case InsEq:
			b, _ := vm.Pop()
			a, _ := vm.Pop()
			fmt.Printf("eq  %d, %d\n", a, b)
			if a == b {
				vm.Push(1)
			} else {
				vm.Push(0)
			}
		case InsIf:
			a, _ := vm.Pop()
			fmt.Printf("if  %d, %d\n", a, inst.value1)
			if a == 0 {
				vm.pc = inst.value1
			}
		case InsJump:
			fmt.Printf("jmp  %d\n", inst.value1)
			vm.pc = inst.value1
		}
		vm.IncrPC()
	}
}

func (vm *VirtualMachine) Result() int {
	return vm.stack[0]
}
