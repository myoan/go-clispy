package main

import (
	"errors"
	"fmt"
)

type NFunction struct {
	name  string
	insts []*Instruction
}

type NFunctionTable struct {
	table []*NFunction
}

func NewNFunctionTable() *NFunctionTable {
	return &NFunctionTable{
		table: make([]*NFunction, 0),
	}
}

func (ft *NFunctionTable) Add(f *NFunction) error {
	_, err := ft.GetFunc(f.name)
	if err == nil {
		return errors.New("Duplicated function")
	}
	ft.table = append(ft.table, f)
	return nil
}

func (ft *NFunctionTable) Register(node *Node) {
	/*
		name := node.children[0].vari
		fnode := node.children[2]
		// fnode.Show()
		code := Compile(fnode)
		// is.Merge(Compile(node))
		// code.Show()
		fmt.Printf("Register %s\n", name)
		t.tbl = append(t.tbl, name)
		t.funcs = append(t.funcs, &Function{insts: code.insts})
	*/
}

func (ft *NFunctionTable) GetFunc(name string) (*NFunction, error) {
	for _, f := range ft.table {
		fmt.Printf("GetFunc serch: %s\n", f.name)
		if f.name == name {
			fmt.Println("  matched")
			return f, nil
		}
	}
	return nil, errors.New("Function not found")
}

func (ft *NFunctionTable) Size() int {
	return len(ft.table)
}

func (ft *NFunctionTable) Show() {
	fmt.Println("FunctionTable Show ======")
	for _, f := range ft.table {
		fmt.Printf("%s:\n", f.name)
		for i, ins := range f.insts {
			fmt.Printf("  [%d] {type: %d, v1: %d, v2: %d}\n", i, ins.iType, ins.value1, ins.value2)
		}
	}
	fmt.Println("FunctionTable Show ======")
}

func NewCompile(ast *Node) *InstructionSet {
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
