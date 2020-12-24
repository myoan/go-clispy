package main

import "fmt"

type InstructionSet struct {
	insts []*Instruction
	ft    *FunctionTable
}

func (is *InstructionSet) IsFunction(name string) bool {
	fmt.Println("Function")
	for _, fn := range is.ft.tbl {
		fmt.Printf("function: %s\n", fn)
		if fn == name {
			fmt.Printf("function found: %s\n", name)
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
	fmt.Println("Merge")
	exists := false
	for i, fname := range dst.tbl {
		for _, srcFuncName := range src.tbl {
			if srcFuncName == fname {
				exists = true
			}
		}
		if !exists {
			fmt.Printf("  merge %s\n", fname)
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
	// fnode.Show()
	code := Compile(fnode)
	// is.Merge(Compile(node))
	// code.Show()
	fmt.Printf("Register %s\n", name)
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

func Compile(ast *Node) *InstructionSet {
	is := &InstructionSet{}
	is.ft = NewFunctionTable()
	if ast.nodeType == Defun {
		is.ft.RegisterFunction(ast)
		return is
	}
	for i, node := range ast.children {
		is.Merge(Compile(node))
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
		is.Add(&Instruction{
			iType: InsAdd,
		})
	case Sub:
		is.Add(&Instruction{
			iType: InsSub,
		})
	case Mul:
		is.Add(&Instruction{
			iType: InsMul,
		})
	case Div:
		is.Add(&Instruction{
			iType: InsDiv,
		})
	case Lt:
		is.Add(&Instruction{
			iType: InsLt,
		})
	case Lte:
		is.Add(&Instruction{
			iType: InsLte,
		})
	case Gt:
		is.Add(&Instruction{
			iType: InsGt,
		})
	case Gte:
		is.Add(&Instruction{
			iType: InsGte,
		})
	case Eq:
		is.Add(&Instruction{
			iType: InsEq,
		})
	case Var:
		if is.IsFunction(ast.vari) {
			is.Add(&Instruction{
				iType: InsCall,
			})
		} else {
			fmt.Println("ELSE!!")
			/*
				is.Add(&Instruction{
					iType: InsVar,
				})
			*/
		}
	case Num:
		is.Add(&Instruction{
			iType:  InsPush,
			value1: ast.value,
		})
	}
	fmt.Println(is.ft.tbl)
	return is
}
