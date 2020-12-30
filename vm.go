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
	stack []int
	cfp   int
	sp    int
	frame []*ControlFrame
}

type ControlFrame struct {
	pc    int
	arg   int
	isset []*Instruction
}

func NewVirtualMachine() *VirtualMachine {
	return &VirtualMachine{
		stack: make([]int, 0),
		sp:    0,
		cfp:   0,
		frame: make([]*ControlFrame, 0),
	}
}

func (vm *VirtualMachine) CurrentFrame() *ControlFrame {
	if len(vm.frame) == 0 {
		return nil
	}
	return vm.frame[vm.cfp]
}

func (vm *VirtualMachine) PushFunction(iss []*Instruction) {
	frame := &ControlFrame{
		pc:    0,
		arg:   0,
		isset: iss,
	}
	vm.cfp += 1
	if vm.cfp >= len(vm.frame) {
		vm.frame = append(vm.frame, frame)
	} else {
		vm.frame[vm.cfp] = frame
	}
}

func (vm *VirtualMachine) PopFunction() {
	vm.cfp -= 1
}

func (vm *VirtualMachine) Show() {
	fmt.Println("--- start stack ---")
	for i, s := range vm.stack {
		fmt.Printf("[%d] %d", i, s)
		if i == vm.sp-1 {
			fmt.Println(" <- sp")
		} else {
			fmt.Println("")
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
	vm.sp--
	vm.stack = vm.stack[:len(vm.stack)-1]
	// if vm.sp > 1 {
	// 	vm.sp--
	// } else {
	// 	vm.sp = 1
	// }
	return ret, nil
}

func (vm *VirtualMachine) IncrPC() {
	vm.CurrentFrame().pc += 1
}

func (vm *VirtualMachine) IsFinish() bool {
	if vm.CurrentFrame().pc >= len(vm.CurrentFrame().isset) {
		if vm.cfp > 0 {
			vm.PopFunction()
			vm.IncrPC()
			return false
		}
		return true
	}
	return false
}

func (vm *VirtualMachine) NextInstruction() *Instruction {
	return vm.CurrentFrame().isset[vm.CurrentFrame().pc]
}

func (vm *VirtualMachine) AddArgs(val int) {
	vm.CurrentFrame().arg = val
}

func (vm *VirtualMachine) Exec(is *InstructionSet) {
	vm.frame = append(vm.frame, &ControlFrame{pc: 0, isset: is.insts})
	for !vm.IsFinish() {
		inst := vm.NextInstruction()
		// vm.Show()
		switch inst.iType {
		case InsPush:
			// fmt.Printf("push %d\n", inst.value1)
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
			// fmt.Printf("add %d, %d\n", a, b)
			vm.Push(a + b)
		case InsSub:
			b, _ := vm.Pop()
			a, _ := vm.Pop()
			// fmt.Printf("sub %d, %d\n", a, b)
			vm.Push(a - b)
		case InsMul:
			b, _ := vm.Pop()
			a, _ := vm.Pop()
			// fmt.Printf("mul %d, %d\n", a, b)
			vm.Push(a * b)
		case InsDiv:
			b, _ := vm.Pop()
			a, _ := vm.Pop()
			// fmt.Printf("div %d, %d\n", a, b)
			vm.Push(a / b)
		case InsLt:
			b, _ := vm.Pop()
			a, _ := vm.Pop()
			// fmt.Printf("lt  %d, %d\n", a, b)
			if a < b {
				vm.Push(1)
			} else {
				vm.Push(0)
			}
		case InsLte:
			b, _ := vm.Pop()
			a, _ := vm.Pop()
			// fmt.Printf("lte %d, %d\n", a, b)
			if a <= b {
				vm.Push(1)
			} else {
				vm.Push(0)
			}
		case InsGt:
			b, _ := vm.Pop()
			a, _ := vm.Pop()
			// fmt.Printf("gt  %d, %d\n", a, b)
			if a > b {
				vm.Push(1)
			} else {
				vm.Push(0)
			}
		case InsGte:
			b, _ := vm.Pop()
			a, _ := vm.Pop()
			// fmt.Printf("gte %d, %d\n", a, b)
			if a >= b {
				vm.Push(1)
			} else {
				vm.Push(0)
			}
		case InsEq:
			b, _ := vm.Pop()
			a, _ := vm.Pop()
			// fmt.Printf("eq  %d, %d\n", a, b)
			if a == b {
				vm.Push(1)
			} else {
				vm.Push(0)
			}
		case InsIf:
			a, _ := vm.Pop()
			// fmt.Printf("if  %d, %d\n", a, inst.value1)
			if a == 0 {
				vm.CurrentFrame().pc = inst.value1
			}
		case InsCall:
			// fmt.Printf("call %d(%d)\n", inst.value1, inst.value2)
			f := is.ft.funcs[inst.value1]
			// fmt.Printf("fn len: %d\n", len(is.ft.funcs))
			// for i, ins := range f.insts {
			// 	fmt.Printf("[%d] {type: %d, v1: %d, v2: %d}\n", i, ins.iType, ins.value1, ins.value2)
			// }
			vm.PushFunction(f.insts)
			for i := 0; i < inst.value2; i++ {
				arg, _ := vm.Pop()
				vm.AddArgs(arg)
			}
			continue
		case InsVar:
			// fmt.Println("var")
			vm.Push(vm.CurrentFrame().arg)
		case InsJump:
			// fmt.Printf("jmp  %d\n", inst.value1)
			vm.CurrentFrame().pc = inst.value1
		}
		vm.IncrPC()
	}
}

func (vm *VirtualMachine) Result() int {
	return vm.stack[0]
}
