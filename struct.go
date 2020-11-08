package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	nodeType string
	children []*Node
	parent   *Node
	value    int
	vari     string
}

func (n *Node) addChild(node *Node) {
	node.parent = n
	n.children = append(n.children, node)
}

func (n *Node) delChild(node *Node) {
	node.parent = nil
	newChildren := make([]*Node, len(n.children)-1)
	i := 0
	for _, c := range n.children {
		if c == node {
			continue
		}
		newChildren[i] = c
		i++
	}
	n.children = newChildren
}

const (
	Non   = "NON"
	Add   = "ADD"
	Sub   = "SUB"
	Mul   = "MUL"
	Div   = "DIV"
	Lt    = "LT"
	Lte   = "LTE"
	Gt    = "GT"
	Gte   = "GTE"
	Eq    = "EQ"
	Num   = "NUM"
	If    = "IF"
	Defun = "DEFUN"
	Var   = "VAR"
	Args  = "ARGS"
	Func  = "FUNC"
)

func (n *Node) Show() {
	fmt.Println("Node Show ======")
	n.show(0)
	fmt.Println("Node Show ======")
}

func (n *Node) show(indent int) {
	for range make([]int, indent) {
		fmt.Print("  ")
	}
	fmt.Printf("%v\n", n)
	for _, child := range n.children {
		child.show(indent + 1)
	}
}

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
	case Defun:
		result += "defun"
	default:
		result += n.vari
	}
	return result
}

func Parse(tl *TokenList) *Node {
	node := parse(tl)
	ft := NewFunctionTable()
	node.Show()
	ft.ExpandAll(node)
	return node
}

func parse(tl *TokenList) *Node {
	root := &Node{
		nodeType: Non,
		children: make([]*Node, 0),
		parent:   nil,
		value:    0,
	}
	current := root
	for tl.Next() {
		token := tl.token
		switch token.tt {
		case Lparen:
			node := &Node{
				nodeType: Non,
				children: make([]*Node, 0),
				value:    0,
				vari:     "paren",
			}
			current.addChild(node)
			current = node
		case Rparen:
			current = current.parent
			if current.nodeType == Defun || current.nodeType == TypeOpr {
				current = current.parent
			}
		case TypeOpr:
			var node *Node
			switch token.value {
			case "add":
				node = &Node{
					nodeType: Add,
					children: make([]*Node, 0),
					value:    0,
					vari:     "+",
				}
				current.addChild(node)
				current = node
			case "sub":
				node = &Node{
					nodeType: Sub,
					children: make([]*Node, 0),
					value:    0,
					vari:     "-",
				}
				current.addChild(node)
				current = node
			case "mul":
				node = &Node{
					nodeType: Mul,
					children: make([]*Node, 0),
					value:    0,
					vari:     "*",
				}
				current.addChild(node)
				current = node
			case "div":
				node = &Node{
					nodeType: Div,
					children: make([]*Node, 0),
					value:    0,
					vari:     "/",
				}
				current.addChild(node)
				current = node
			case "lt":
				node = &Node{
					nodeType: Lt,
					children: make([]*Node, 0),
					value:    0,
					vari:     "<",
				}
				current.addChild(node)
				current = node
			case "lte":
				node = &Node{
					nodeType: Lte,
					children: make([]*Node, 0),
					value:    0,
					vari:     "<=",
				}
				current.addChild(node)
				current = node
			case "gt":
				node = &Node{
					nodeType: Gt,
					children: make([]*Node, 0),
					value:    0,
					vari:     ">",
				}
				current.addChild(node)
				current = node
			case "gte":
				node = &Node{
					nodeType: Gte,
					children: make([]*Node, 0),
					value:    0,
					vari:     ">=",
				}
				current.addChild(node)
				current = node
			case "eq":
				node = &Node{
					nodeType: Eq,
					children: make([]*Node, 0),
					value:    0,
					vari:     "==",
				}
				current.addChild(node)
				current = node
			case "if":
				node = &Node{
					nodeType: If,
					children: make([]*Node, 0),
					value:    0,
					vari:     "IF",
				}
				current.addChild(node)
				current = node
			case "defun":
				node = &Node{
					nodeType: Defun,
					children: make([]*Node, 0),
					value:    0,
					vari:     "DEFUN",
				}
				current.addChild(node)
				current = node

				tl.Next()
				nameNode := &Node{
					nodeType: Var,
					children: make([]*Node, 0),
					value:    0,
					vari:     tl.token.value,
				}
				current.addChild(nameNode)

				tl.Next() // maybe (
				tl.Next()
				argsNode := &Node{
					nodeType: Args,
					children: make([]*Node, 0),
					value:    0,
					vari:     tl.token.value,
				}
				current.addChild(argsNode)
				tl.Next() // maybe )
			default:
				node = &Node{
					nodeType: Func,
					children: make([]*Node, 0),
					value:    0,
					vari:     tl.token.value,
				}
				current.addChild(node)
				current = node
			}
		case TypeSymbol:
			var node *Node
			switch token.value {
			case "if":
				node = &Node{
					nodeType: If,
					children: make([]*Node, 0),
					value:    0,
					vari:     "IF",
				}
				current.addChild(node)
				current = node
			case "defun":
				node = &Node{
					nodeType: Defun,
					children: make([]*Node, 0),
					value:    0,
					vari:     "DEFUN",
				}
				current.addChild(node)
				current = node

				tl.Next()
				nameNode := &Node{
					nodeType: Var,
					children: make([]*Node, 0),
					value:    0,
					vari:     tl.token.value,
				}
				current.addChild(nameNode)

				tl.Next() // maybe (
				tl.Next()
				argsNode := &Node{
					nodeType: Args,
					children: make([]*Node, 0),
					value:    0,
					vari:     tl.token.value,
				}
				current.addChild(argsNode)
				tl.Next() // maybe )
			default:
				node = &Node{
					nodeType: Var,
					children: make([]*Node, 0),
					value:    0,
					vari:     token.value,
				}
				current.addChild(node)
			}
		case TypeInteger:
			data, _ := strconv.Atoi(token.value)
			node := &Node{
				nodeType: Num,
				children: make([]*Node, 0),
				value:    data,
				vari:     "",
			}
			current.addChild(node)
		}
	}
	return root
}
