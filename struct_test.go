package main

import (
	"fmt"
	"testing"
)

/*
func TestExpandFunction(t *testing.T) {
	testcases := []struct {
		in     []*Token
		msg    string
		expect string
	}{
		{
			in: []*Token{
				{tt: Lparen, value: ""},
				{tt: TypeOpr, value: "incr"},
				{tt: TypeInteger, value: "1"},
				{tt: Rparen, value: ""},
			},
			msg:    "defun stmt",
			expect: "incrnn1+defun",
		},
	}

	mtd := []*Token {
		{tt: Lparen, value: ""},
		{tt: TypeOpr, value: "add"},
		{tt: TypeSymbol, value: "n"},
		{tt: TypeInteger, value: "1"},
		{tt: Rparen, value: ""},
	}

	for _, tt := range testcases {
		tl := NewTokenList()
		for _, token := range tt.in {
			tl.Push(token)
		}
		node := parse(tl)
		ftable := make(map[string]*Function)
		actual := ExpandFunction(node, ftable)
		fmt.Printf("%v\n", actual)
		if actual.Text() != tt.expect {
			t.Errorf("[%s] expect: %s, actual: %s\n", tt.msg, tt.expect, actual.Text())
			actual.Show()
		}
	}
}
*/

func IsSameNode(actual, expect *Node) bool {
	for i, child := range actual.children {
		if !IsSameNode(child, expect.children[i]) {
			return false
		}
	}
	if actual.nodeType != expect.nodeType {
		fmt.Println("nodetype")
		return false
	}
	if len(actual.children) != len(expect.children) {
		fmt.Printf("child len e: %d, a: %d\n", len(actual.children), len(expect.children))
		return false
	}
	if actual.value != expect.value {
		fmt.Println("value")
		return false
	}
	if actual.vari != expect.vari {
		fmt.Println("vari")
		return false
	}
	return true
}

func TestCreateAST(t *testing.T) {
	testcases := []struct {
		in     *TokenList
		msg    string
		expect *Node
	}{
		{
			in: &TokenList{
				idx:   0,
				token: nil,
				tokens: []*Token{
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeRParen, value: ""},
				},
			},
			msg: "empty",
			expect: &Node{
				nodeType: Non,
				children: []*Node{
					{
						nodeType: Non,
						vari:     "paren",
						children: []*Node{},
					},
				},
			},
		},
		{
			in: &TokenList{
				idx:   0,
				token: nil,
				tokens: []*Token{
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeAdd, value: ""},
					{tt: TokenTypeInt, value: "1"},
					{tt: TokenTypeInt, value: "2"},
					{tt: TokenTypeRParen, value: ""},
				},
			},
			msg: "simple add",
			expect: &Node{
				nodeType: Non,
				children: []*Node{
					{
						nodeType: Non,
						vari:     "paren",
						children: []*Node{
							{
								nodeType: Add,
								vari:     "+",
								children: []*Node{
									{
										nodeType: Num,
										vari:     "",
										value:    1,
										children: []*Node{},
									},
									{
										nodeType: Num,
										vari:     "",
										value:    2,
										children: []*Node{},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			in: &TokenList{
				idx:   0,
				token: nil,
				tokens: []*Token{
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeSub, value: ""},
					{tt: TokenTypeInt, value: "1"},
					{tt: TokenTypeInt, value: "2"},
					{tt: TokenTypeRParen, value: ""},
				},
			},
			msg: "simple sub",
			expect: &Node{
				nodeType: Non,
				children: []*Node{
					{
						nodeType: Non,
						vari:     "paren",
						children: []*Node{
							{
								nodeType: Sub,
								vari:     "-",
								children: []*Node{
									{
										nodeType: Num,
										vari:     "",
										value:    1,
										children: []*Node{},
									},
									{
										nodeType: Num,
										vari:     "",
										value:    2,
										children: []*Node{},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			in: &TokenList{
				idx:   0,
				token: nil,
				tokens: []*Token{
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeAdd, value: ""},
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeAdd, value: ""},
					{tt: TokenTypeInt, value: "1"},
					{tt: TokenTypeInt, value: "2"},
					{tt: TokenTypeRParen, value: ""},
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeSub, value: ""},
					{tt: TokenTypeInt, value: "3"},
					{tt: TokenTypeInt, value: "4"},
					{tt: TokenTypeRParen, value: ""},
					{tt: TokenTypeRParen, value: ""},
				},
			},
			msg: "nested",
			expect: &Node{
				nodeType: Non,
				children: []*Node{
					{
						nodeType: Non,
						vari:     "paren",
						children: []*Node{
							{
								nodeType: Add,
								vari:     "+",
								children: []*Node{
									{
										nodeType: Non,
										vari:     "paren",
										children: []*Node{
											{
												nodeType: Add,
												vari:     "+",
												children: []*Node{
													{
														nodeType: Num,
														vari:     "",
														value:    1,
														children: []*Node{},
													},
													{
														nodeType: Num,
														vari:     "",
														value:    2,
														children: []*Node{},
													},
												},
											},
										},
									},
									{
										nodeType: Non,
										vari:     "paren",
										children: []*Node{
											{
												nodeType: Sub,
												vari:     "-",
												children: []*Node{
													{
														nodeType: Num,
														vari:     "",
														value:    3,
														children: []*Node{},
													},
													{
														nodeType: Num,
														vari:     "",
														value:    4,
														children: []*Node{},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			in: &TokenList{
				idx:   0,
				token: nil,
				tokens: []*Token{
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeKeyword, value: "if"},
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeGt, value: ""},
					{tt: TokenTypeInt, value: "1"},
					{tt: TokenTypeInt, value: "2"},
					{tt: TokenTypeRParen, value: ""},
					{tt: TokenTypeInt, value: "3"},
					{tt: TokenTypeInt, value: "4"},
					{tt: TokenTypeRParen, value: ""},
				},
			},
			msg: "if stmt",
			expect: &Node{
				nodeType: Non,
				children: []*Node{
					{
						nodeType: Non,
						vari:     "paren",
						children: []*Node{
							{
								nodeType: If,
								vari:     "IF",
								children: []*Node{
									{
										nodeType: Non,
										vari:     "paren",
										children: []*Node{
											{
												nodeType: Gt,
												vari:     ">",
												children: []*Node{
													{
														nodeType: Num,
														vari:     "",
														value:    1,
														children: []*Node{},
													},
													{
														nodeType: Num,
														vari:     "",
														value:    2,
														children: []*Node{},
													},
												},
											},
										},
									},
									{
										nodeType: Num,
										vari:     "",
										value:    3,
										children: []*Node{},
									},
									{
										nodeType: Num,
										vari:     "",
										value:    4,
										children: []*Node{},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			in: &TokenList{
				idx:   0,
				token: nil,
				tokens: []*Token{
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeKeyword, value: "defun"},
					{tt: TokenTypeIdent, value: "incr"},
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeIdent, value: "n"},
					{tt: TokenTypeRParen, value: ""},
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeAdd, value: ""},
					{tt: TokenTypeIdent, value: "n"},
					{tt: TokenTypeInt, value: "1"},
					{tt: TokenTypeRParen, value: ""},
					{tt: TokenTypeRParen, value: ""},
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeIdent, value: "incr"},
					{tt: TokenTypeInt, value: "1"},
					{tt: TokenTypeRParen, value: ""},
				},
			},
			msg: "defun",
			expect: &Node{
				nodeType: Non,
				children: []*Node{
					{
						nodeType: Non,
						vari:     "paren",
						children: []*Node{
							{
								nodeType: Defun,
								vari:     "DEFUN",
								children: []*Node{
									{
										nodeType: Var,
										vari:     "incr",
										children: []*Node{},
									},
									{
										nodeType: Args,
										vari:     "n",
										children: []*Node{},
									},
									{
										nodeType: Non,
										vari:     "paren",
										children: []*Node{
											{
												nodeType: Add,
												vari:     "+",
												children: []*Node{
													{
														nodeType: Var,
														vari:     "n",
														children: []*Node{},
													},
													{
														nodeType: Num,
														vari:     "",
														value:    1,
														children: []*Node{},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					{
						nodeType: Non,
						vari:     "paren",
						children: []*Node{
							{
								nodeType: Var,
								vari:     "incr",
								children: []*Node{
									{
										nodeType: Num,
										vari:     "",
										value:    1,
										children: []*Node{},
									},
								},
							},
						},
					},
				},
			},
		},
		/*
			{
				in: []*Token{
					{tt: Lparen, value: ""},
					{tt: TypeOpr, value: "defun"},
					{tt: TypeSymbol, value: "fib"},
					{tt: Lparen, value: ""},
					{tt: TypeSymbol, value: "n"},
					{tt: Rparen, value: ""},
					{tt: Lparen, value: ""},
					{tt: TypeSymbol, value: "if"},
					{tt: Lparen, value: ""},
					{tt: TypeOpr, value: "lte"},
					{tt: TypeSymbol, value: "n"},
					{tt: TypeInteger, value: "1"},
					{tt: Rparen, value: ""},
					{tt: TypeInteger, value: "1"},
					{tt: Lparen, value: ""},
					{tt: TypeOpr, value: "add"},
					{tt: Lparen, value: ""},
					{tt: TypeOpr, value: "fib"},
					{tt: Lparen, value: ""},
					{tt: TypeOpr, value: "sub"},
					{tt: TypeSymbol, value: "n"},
					{tt: TypeInteger, value: "1"},
					{tt: Rparen, value: ""},
					{tt: Rparen, value: ""},
					{tt: Lparen, value: ""},
					{tt: TypeOpr, value: "fib"},
					{tt: Lparen, value: ""},
					{tt: TypeOpr, value: "sub"},
					{tt: TypeSymbol, value: "n"},
					{tt: TypeInteger, value: "2"},
					{tt: Rparen, value: ""},
					{tt: Rparen, value: ""},
					{tt: Rparen, value: ""},
					{tt: Rparen, value: ""},
					{tt: Rparen, value: ""},
					{tt: Lparen, value: ""},
					{tt: TypeOpr, value: "fib"},
					{tt: TypeInteger, value: "4"},
					{tt: Rparen, value: ""},
				},
				msg:    "fibonacci stmt",
				expect: "[[defun[fib,n,[if[[<=[n,1]],1,[+[[fib[[-[n,1]]]],[fib[[-[n,2]]]]]]]]]],[fib[4]]]",
			},
		*/
	}

	for _, tt := range testcases {
		actual := CreateAST(tt.in)
		fmt.Printf("%v\n", actual)
		if !IsSameNode(actual, tt.expect) {
			t.Errorf("[%s] expect: %s, actual: %s\n", tt.msg, tt.expect.Text(), actual.Text())
			tt.expect.Show()
			actual.Show()
		}
	}
}
