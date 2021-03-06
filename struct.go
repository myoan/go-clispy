package main

import (
	"fmt"
	"strconv"
	"strings"
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

func (n *Node) isOpr() bool {
	if n.nodeType == Add ||
		n.nodeType == Sub ||
		n.nodeType == Mul ||
		n.nodeType == Div ||
		n.nodeType == Lt ||
		n.nodeType == Lte ||
		n.nodeType == Gt ||
		n.nodeType == Gte ||
		n.nodeType == Eq ||
		n.nodeType == If ||
		n.nodeType == Defun ||
		n.nodeType == Var ||
		n.nodeType == Func {
		return true
	}
	return false
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
	fmt.Printf("(%s %d %s)\n", n.nodeType, n.value, n.vari)
	for _, child := range n.children {
		child.show(indent + 1)
	}
}

func (n *Node) Text() string {
	result := ""
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
	if len(n.children) > 0 {
		result += "["
		children := make([]string, 0)
		for _, c := range n.children {
			children = append(children, c.Text())
		}
		result += strings.Join(children, ",")
		result += "]"
	}
	return result
}

func Parse(tl *TokenList) (*Node, error) {
	node := CreateAST(tl)
	return node, nil
}

func CreateAST(tl *TokenList) *Node {
	root := &Node{
		nodeType: Non,
		children: make([]*Node, 0),
		parent:   nil,
		value:    0,
	}
	current := root
	for tl.Next() {
		token := tl.token
		// fmt.Printf("{tt: %s, value: \"%s\"},\n", token.tt, token.value)
		switch token.tt {
		case TokenTypeLParen:
			node := &Node{
				nodeType: Non,
				children: make([]*Node, 0),
				value:    0,
				vari:     "paren",
			}
			current.addChild(node)
			current = node
		case TokenTypeRParen:
			if current.isOpr() {
				current = current.parent
			}
			current = current.parent
		case TokenTypeAdd:
			node := &Node{
				nodeType: Add,
				children: make([]*Node, 0),
				value:    0,
				vari:     "+",
			}
			current.addChild(node)
			current = node
		case TokenTypeSub:
			node := &Node{
				nodeType: Sub,
				children: make([]*Node, 0),
				value:    0,
				vari:     "-",
			}
			current.addChild(node)
			current = node
		case TokenTypeMul:
			node := &Node{
				nodeType: Mul,
				children: make([]*Node, 0),
				value:    0,
				vari:     "*",
			}
			current.addChild(node)
			current = node
		case TokenTypeDiv:
			node := &Node{
				nodeType: Div,
				children: make([]*Node, 0),
				value:    0,
				vari:     "/",
			}
			current.addChild(node)
			current = node
		case TokenTypeLt:
			node := &Node{
				nodeType: Lt,
				children: make([]*Node, 0),
				value:    0,
				vari:     "<",
			}
			current.addChild(node)
			current = node
		case TokenTypeLte:
			node := &Node{
				nodeType: Lte,
				children: make([]*Node, 0),
				value:    0,
				vari:     "<=",
			}
			current.addChild(node)
			current = node
		case TokenTypeGt:
			node := &Node{
				nodeType: Gt,
				children: make([]*Node, 0),
				value:    0,
				vari:     ">",
			}
			current.addChild(node)
			current = node
		case TokenTypeGte:
			node := &Node{
				nodeType: Gte,
				children: make([]*Node, 0),
				value:    0,
				vari:     ">=",
			}
			current.addChild(node)
			current = node
		case TokenTypeEq:
			node := &Node{
				nodeType: Eq,
				children: make([]*Node, 0),
				value:    0,
				vari:     "==",
			}
			current.addChild(node)
			current = node
		case TokenTypeKeyword:
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
		case TokenTypeInt:
			data, _ := strconv.Atoi(token.value)
			node := &Node{
				nodeType: Num,
				children: make([]*Node, 0),
				value:    data,
				vari:     "",
			}
			current.addChild(node)
		case TokenTypeIdent:
			node := &Node{
				nodeType: Var,
				children: make([]*Node, 0),
				value:    0,
				vari:     token.value,
			}
			current.addChild(node)
			if current.nodeType == Non { // LParen
				current = node
			}
		}
	}
	return root
}
