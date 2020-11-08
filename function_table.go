package main

import (
	"fmt"
)

type Function struct {
	args string
	mtd  *Node
}

type FunctionTable struct {
	fns map[string]*Function
}

func NewFunctionTable() *FunctionTable {
	return &FunctionTable{
		fns: make(map[string]*Function, 0),
	}
}

func (ft *FunctionTable) ExpandAll(node *Node) {
	ft.RegisterAll(node)
	ft.expandAll(node)
}

func (ft *FunctionTable) HasFunc(name string) bool {
	fn := ft.fns[name]
	return fn != nil
}

func (ft *FunctionTable) GetFunc(name string) (*Function, error) {
	fn := ft.fns[name]
	if fn == nil {
		return nil, fmt.Errorf("Undefined Function: %s", name)
	}
	return fn, nil
}

func (ft *FunctionTable) RegisterAll(node *Node) {
	if node.nodeType == Defun {
		name := node.children[0].vari
		args := node.children[1].vari
		mtd := node.children[2].children[0]
		ft.Register(name, args, mtd)
		node.parent.delChild(node)
		return
	}
	for _, child := range node.children {
		ft.RegisterAll(child)
	}
}

func (ft *FunctionTable) Register(name string, args string, node *Node) {
	fmt.Printf("Register: %s\n", name)
	ft.fns[name] = &Function{
		args: args,
		mtd:  node,
	}
}

func (ft *FunctionTable) expandAll(node *Node) {
	if node.nodeType == Func {
		if ft.HasFunc(node.vari) {
			ft.Expand(node.parent)
		}
		return
	}
	for _, child := range node.children {
		ft.expandAll(child)
	}
}

func (ft *FunctionTable) Expand(node *Node) {
	fNode := node.children[0]
	param := fNode.children[0].value
	fn, err := ft.GetFunc(fNode.vari)
	if err != nil {
		panic(err)
	}
	*fNode = *ft.expandParams(fn.mtd, fn.args, param)
}

func (ft *FunctionTable) expandParams(node *Node, param string, value int) *Node {
	if node.nodeType == Var && node.vari == param {
		return &Node{
			nodeType: Num,
			children: make([]*Node, 0),
			value:    value,
			vari:     "",
		}
	}
	children := make([]*Node, 0)
	for _, ch := range node.children {
		children = append(children, ft.expandParams(ch, param, value))
	}
	return &Node{
		nodeType: node.nodeType,
		children: children,
		value:    node.value,
		vari:     node.vari,
	}
}
