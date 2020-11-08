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
	Lt
	Lte
	Gt
	Gte
	Eq
	Num
	If
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
	case Lt:
		result += "<"
	case Lte:
		result += "<="
	case Gt:
		result += ">"
	case Gte:
		result += ">="
	case Eq:
		result += "=="
	case Num:
		result += strconv.Itoa(n.value)
	case If:
		result += "if"
	default:
		result += "hoge"
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
		case TypeOpr:
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
			case "lt":
				node = &Node{
					nodeType: Lt,
					children: make([]*Node, 0),
					value:    0,
				}
			case "lte":
				node = &Node{
					nodeType: Lte,
					children: make([]*Node, 0),
					value:    0,
				}
			case "gt":
				node = &Node{
					nodeType: Gt,
					children: make([]*Node, 0),
					value:    0,
				}
			case "gte":
				node = &Node{
					nodeType: Gte,
					children: make([]*Node, 0),
					value:    0,
				}
			case "eq":
				node = &Node{
					nodeType: Eq,
					children: make([]*Node, 0),
					value:    0,
				}
			}
			current.addChild(node)
			current = node
		case TypeSymbol:
			node := &Node{
				nodeType: If,
				children: make([]*Node, 0),
				value:    0,
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
