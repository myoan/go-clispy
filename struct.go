package main

import (
	"strconv"
)

type Node struct {
	nodeType int
	children []*Node
	parent   *Node
	value    int
}

func (n *Node) addChild(node *Node) {
	node.parent = n
	n.children = append(n.children, node)
}

const (
	Non = iota
	Add
	Sub
	Mul
	Div
	Num
)

func (n *Node) Text() string {
	result := ""
	for _, c := range n.children {
		result += c.Text()
	}
	switch n.nodeType {
	case Non:
	case Add:
		result += "+"
	case Sub:
		result += "-"
	case Mul:
		result += "*"
	case Div:
		result += "/"
	case Num:
		result += strconv.Itoa(n.value)
	}
	return result
}

func Parse(tokens []*Token) *Node {
	root := &Node{
		nodeType: Non,
		children: make([]*Node, 0),
		parent:   nil,
		value:    0,
	}
	current := root
	for _, token := range tokens {
		switch token.tt {
		case Lparen:
		case Rparen:
			current = current.parent
		case TypeSymbol:
			var node *Node
			switch token.value {
			case "add":
				node = &Node{
					nodeType: Add,
					children: make([]*Node, 0),
					value:    0,
				}
			case "sub":
				node = &Node{
					nodeType: Sub,
					children: make([]*Node, 0),
					value:    0,
				}
			case "mul":
				node = &Node{
					nodeType: Mul,
					children: make([]*Node, 0),
					value:    0,
				}
			case "div":
				node = &Node{
					nodeType: Div,
					children: make([]*Node, 0),
					value:    0,
				}
			}
			current.addChild(node)
			current = node
		case TypeInteger:
			data, _ := strconv.Atoi(token.value)
			node := &Node{
				nodeType: Num,
				children: make([]*Node, 0),
				value:    data,
			}
			current.addChild(node)
		}
	}
	return root
}
