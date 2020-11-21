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

func TestParseAST(t *testing.T) {
	testcases := []struct {
		in     []*Token
		msg    string
		expect string
	}{
		{
			in: []*Token{
				{tt: Lparen, value: ""},
				{tt: Rparen, value: ""},
			},
			msg:    "empty",
			expect: "[]",
		},
		{
			in: []*Token{
				{tt: Lparen, value: ""},
				{tt: TypeOpr, value: "add"},
				{tt: TypeInteger, value: "1"},
				{tt: TypeInteger, value: "2"},
				{tt: Rparen, value: ""},
			},
			msg:    "simple add",
			expect: "[[+[1,2]]]",
		},
		{
			in: []*Token{
				{tt: Lparen, value: ""},
				{tt: TypeOpr, value: "sub"},
				{tt: TypeInteger, value: "1"},
				{tt: TypeInteger, value: "2"},
				{tt: Rparen, value: ""},
			},
			msg:    "simple sub",
			expect: "[[-[1,2]]]",
		},
		{
			in: []*Token{
				{tt: Lparen, value: ""},
				{tt: TypeOpr, value: "mul"},
				{tt: TypeInteger, value: "1"},
				{tt: TypeInteger, value: "2"},
				{tt: Rparen, value: ""},
			},
			msg:    "simple mul",
			expect: "[[*[1,2]]]",
		},
		{
			in: []*Token{
				{tt: Lparen, value: ""},
				{tt: TypeOpr, value: "div"},
				{tt: TypeInteger, value: "1"},
				{tt: TypeInteger, value: "2"},
				{tt: Rparen, value: ""},
			},
			msg:    "simple div",
			expect: "[[/[1,2]]]",
		},
		{
			in: []*Token{
				{tt: Lparen, value: ""},
				{tt: TypeOpr, value: "add"},
				{tt: Lparen, value: ""},
				{tt: TypeOpr, value: "add"},
				{tt: TypeInteger, value: "1"},
				{tt: TypeInteger, value: "2"},
				{tt: Rparen, value: ""},
				{tt: Lparen, value: ""},
				{tt: TypeOpr, value: "sub"},
				{tt: TypeInteger, value: "3"},
				{tt: TypeInteger, value: "4"},
				{tt: Rparen, value: ""},
				{tt: Rparen, value: ""},
			},
			msg:    "nested",
			expect: "[[+[[+[1,2]],[-[3,4]]]]]",
		},
		{
			in: []*Token{
				{tt: Lparen, value: ""},
				{tt: TypeSymbol, value: "if"},
				{tt: Lparen, value: ""},
				{tt: TypeOpr, value: "gt"},
				{tt: TypeInteger, value: "1"},
				{tt: TypeInteger, value: "2"},
				{tt: Rparen, value: ""},
				{tt: TypeInteger, value: "3"},
				{tt: TypeInteger, value: "4"},
				{tt: Rparen, value: ""},
			},
			msg:    "if stmt",
			expect: "[[if[[>[1,2]],3,4]]]",
		},
		{
			in: []*Token{
				{tt: Lparen, value: ""},
				{tt: TypeOpr, value: "defun"},
				{tt: TypeSymbol, value: "incr"},

				{tt: Lparen, value: ""},
				{tt: TypeOpr, value: "n"},
				{tt: Rparen, value: ""},

				{tt: Lparen, value: ""},
				{tt: TypeOpr, value: "add"},
				{tt: TypeSymbol, value: "n"},
				{tt: TypeInteger, value: "1"},
				{tt: Rparen, value: ""},

				{tt: Rparen, value: ""},

				{tt: Lparen, value: ""},
				{tt: TypeOpr, value: "incr"},
				{tt: TypeInteger, value: "1"},
				{tt: Rparen, value: ""},
			},
			msg:    "defun stmt",
			expect: "[[defun[incr,n,[+[n,1]]]],[incr[1]]]",
		},
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
	}

	for _, tt := range testcases {
		tl := NewTokenList()
		for _, token := range tt.in {
			tl.Push(token)
		}
		actual := CreateAST(tl)
		fmt.Printf("%v\n", actual)
		if actual.Text() != tt.expect {
			t.Errorf("[%s] expect: %s, actual: %s\n", tt.msg, tt.expect, actual.Text())
			actual.Show()
		}
	}
}
