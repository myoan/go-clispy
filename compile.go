package main

import "fmt"

type NFunctionTable struct {
	fns map[string]*NFunction
}

type NFunction struct {
	args  string
	insts []*Instruction
}

func (t *NFunctionTable) Merge(ft *NFunctionTable) {
	// for _, s, f := range ft.fns {
	// 	t.fns[f.]
	// }
}

type InstructionSet struct {
	insts []*Instruction
	ft    *NFunctionTable
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
	InsFunc
)

type Instruction struct {
	iType  InstType
	value1 int
	value2 int
}

func Compile(ast *Node) *InstructionSet {
	is := &InstructionSet{}
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
	case Defun:
		is.Add(&Instruction{
			iType: InsDefun,
		})
	case Func:
		is.Add(&Instruction{
			iType: InsFunc,
		})
	case Num:
		is.Add(&Instruction{
			iType:  InsPush,
			value1: ast.value,
		})
	}
	return is
}
