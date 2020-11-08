package main

import "fmt"

type Opr int

const (
	OprAdd = iota
	OprSub
	OprMul
	OprDiv
	OprPush
	OprPop
	TypeNum
)

type StackMachine struct {
	stack []int
}

func NewStackMachine() *StackMachine {
	return &StackMachine{stack: make([]int, 0)}
}

func (sm *StackMachine) Push(v int) {
	sm.stack = append(sm.stack, v)
}

func (sm *StackMachine) Pop() int {
	ret := sm.stack[len(sm.stack)-1]
	sm.stack = sm.stack[:len(sm.stack)-1]
	return ret
}

func (sm *StackMachine) Result() int {
	return sm.stack[0]
}

func Eval(sm *StackMachine, node *Node) {
	for _, n := range node.children {
		Eval(sm, n)
	}
	switch node.nodeType {
	case Non:
		fmt.Println("NON")
	case Add:
		a := sm.Pop()
		b := sm.Pop()
		fmt.Printf("ADD %d, %d\n", a, b)
		sm.Push(b + a)
	case Sub:
		a := sm.Pop()
		b := sm.Pop()
		fmt.Printf("SUB %d, %d\n", a, b)
		sm.Push(b - a)
	case Mul:
		a := sm.Pop()
		b := sm.Pop()
		fmt.Printf("MUL %d, %d\n", a, b)
		sm.Push(b * a)
	case Div:
		a := sm.Pop()
		b := sm.Pop()
		fmt.Printf("DIV %d, %d\n", a, b)
		sm.Push(a / b)
	case Lt:
		a := sm.Pop()
		b := sm.Pop()
		fmt.Printf("LT %d, %d\n", a, b)
		if b < a {
			sm.Push(1)
		} else {
			sm.Push(0)
		}
	case Lte:
		a := sm.Pop()
		b := sm.Pop()
		fmt.Printf("LTE %d, %d\n", a, b)
		if b <= a {
			sm.Push(1)
		} else {
			sm.Push(0)
		}
	case Gt:
		a := sm.Pop()
		b := sm.Pop()
		fmt.Printf("GT %d, %d\n", a, b)
		if b > a {
			sm.Push(1)
		} else {
			sm.Push(0)
		}
	case Gte:
		a := sm.Pop()
		b := sm.Pop()
		fmt.Printf("GTE %d, %d\n", a, b)
		sm.Push(a / b)
		if b >= a {
			sm.Push(1)
		} else {
			sm.Push(0)
		}
	case Eq:
		a := sm.Pop()
		b := sm.Pop()
		fmt.Printf("EQ %d, %d\n", a, b)
		if b == a {
			sm.Push(1)
		} else {
			sm.Push(0)
		}
	case If:
		a := sm.Pop()
		b := sm.Pop()
		c := sm.Pop()
		fmt.Printf("IF %d, %d, %d\n", c, b, a)
		if c == 1 {
			sm.Push(b)
		} else {
			sm.Push(a)
		}
	case Num:
		fmt.Printf("PUSH %d\n", node.value)
		sm.Push(node.value)
	}
	fmt.Printf("%v\n", sm.stack)
}
