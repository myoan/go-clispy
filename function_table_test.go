package main

import "testing"

func TestFunctionTable_Expand(t *testing.T) {
	ft := NewFunctionTable()
	mtd := &Node{
		nodeType: Add,
		children: []*Node{
			{
				nodeType: Var,
				children: make([]*Node, 0),
				value:    0,
				vari:     "n",
			},
			{
				nodeType: Num,
				children: make([]*Node, 0),
				value:    1,
				vari:     "",
			},
		},
		value: 0,
		vari:  "+",
	}
	ft.Register("incr", "n", mtd)

	tokens := []*Token{
		{tt: Lparen, value: ""},
		{tt: TypeOpr, value: "incr"},
		{tt: TypeInteger, value: "1"},
		{tt: Rparen, value: ""},
	}
	tl := NewTokenList()
	for _, token := range tokens {
		tl.Push(token)
	}
	node := CreateAST(tl)
	ft.Expand(node.children[0])
	expect := "[[+[1,1]]]"
	if node.Text() != expect {
		t.Errorf("expect: %s, actual: %s\n", expect, node.Text())
	}
}
