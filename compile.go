package main

import (
	"fmt"
)

type InstructionSet struct {
	insts []*Instruction
	ft    *FunctionTable
}

func (is *InstructionSet) funcs() {
	fmt.Printf("func size: %d\n", len(is.ft.tbl))
	for _, fn := range is.ft.tbl {
		fmt.Printf("func: %s\n", fn)
	}
}

func (is *InstructionSet) IsFunction(name string) bool {
	for _, fn := range is.ft.tbl {
		if fn == name {
			return true
		}
	}
	return false
}

type FunctionTable struct {
	tbl   []string
	funcs []*Function
}

func NewFunctionTable() *FunctionTable {
	return &FunctionTable{
		tbl:   make([]string, 0),
		funcs: make([]*Function, 0),
	}
}

type Function struct {
	insts []*Instruction
}

func (src *FunctionTable) Merge(dst *FunctionTable) error {
	exists := false
	for i, fname := range dst.tbl {
		for _, srcFuncName := range src.tbl {
			if srcFuncName == fname {
				exists = true
			}
		}
		if !exists {
			f := dst.funcs[i]
			src.tbl = append(src.tbl, fname)
			src.funcs = append(src.funcs, f)
		}
		exists = false
	}
	return nil
}

func (t *FunctionTable) RegisterFunction(node *Node) {
	name := node.children[0].vari
	fnode := node.children[2]
	code := Compile(fnode, nil)
	t.tbl = append(t.tbl, name)
	t.funcs = append(t.funcs, &Function{insts: code.insts})
}

func (is *InstructionSet) GetFunc(idx int) *Function {
	return is.ft.funcs[idx]
}

func (is *InstructionSet) Size() int {
	return len(is.insts)
}

func (is *InstructionSet) Add(i *Instruction) {
	is.insts = append(is.insts, i)
}

func (is *InstructionSet) Merge(is2 *InstructionSet) {
	for _, i := range is2.insts {
		is.insts = append(is.insts, i)
	}
	is.ft.Merge(is2.ft)
}

func (is *InstructionSet) Show() {
	fmt.Println("InstSet Show ======")
	for i, ins := range is.insts {
		fmt.Printf("[%d] {type: %d, v1: %d, v2: %d}\n", i, ins.iType, ins.value1, ins.value2)
	}
	fmt.Println("InstSet Show ======")
}

type InstType int

const (
	InsPush = iota
	InsPop
	InsSend
	InsAdd
	InsSub
	InsMul
	InsDiv
	InsLt
	InsLte
	InsGt
	InsGte
	InsEq
	InsIf
	InsJump
	InsDefun
	InsCall
	InsVar
	InsLabel
	InsRet
)

type Instruction struct {
	iType  InstType
	value1 int
	value2 int
}

func Compile(ast *Node, ft *FunctionTable) *InstructionSet {
	is := &InstructionSet{}
	if ft == nil {
		is.ft = NewFunctionTable()
	} else {
		is.ft = ft
	}
	if ast.nodeType == Defun {
		is.ft.RegisterFunction(ast)
		return is
	}
	for i, node := range ast.children {
		is.Merge(Compile(node, is.ft))
		if ast.nodeType == If {
			if i == 0 {
				is.Add(&Instruction{
					iType:  InsIf,
					value1: is.Size() + 2,
				})
			} else if i == 1 {
				is.Add(&Instruction{
					iType:  InsJump,
					value1: is.Size() + 3,
				})
			}
		}
	}
	switch ast.nodeType {
	case Add:
		fmt.Printf("Add\n")
		is.Add(&Instruction{
			iType: InsAdd,
		})
	case Sub:
		fmt.Printf("Sub\n")
		is.Add(&Instruction{
			iType: InsSub,
		})
	case Mul:
		fmt.Printf("Mul\n")
		is.Add(&Instruction{
			iType: InsMul,
		})
	case Div:
		fmt.Printf("Div\n")
		is.Add(&Instruction{
			iType: InsDiv,
		})
	case Lt:
		fmt.Printf("Lt\n")
		is.Add(&Instruction{
			iType: InsLt,
		})
	case Lte:
		fmt.Printf("Lte\n")
		is.Add(&Instruction{
			iType: InsLte,
		})
	case Gt:
		fmt.Printf("Gt\n")
		is.Add(&Instruction{
			iType: InsGt,
		})
	case Gte:
		fmt.Printf("Gte\n")
		is.Add(&Instruction{
			iType: InsGte,
		})
	case Eq:
		fmt.Printf("Eq\n")
		is.Add(&Instruction{
			iType: InsEq,
		})
	case Var:
		fmt.Printf("Var: %s\n", ast.vari)
		if is.IsFunction(ast.vari) {
			is.Add(&Instruction{
				iType: InsCall,
			})
		} else {
			/*
				is.Add(&Instruction{
					iType: InsVar,
				})
			*/
		}
	case Num:
		fmt.Printf("Num %d\n", ast.value)
		is.Add(&Instruction{
			iType:  InsPush,
			value1: ast.value,
		})
	}
	return is
}
