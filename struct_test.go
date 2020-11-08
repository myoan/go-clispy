package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
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
			expect: "",
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
			expect: "12+",
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
			expect: "12-",
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
			expect: "12*",
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
			expect: "12/",
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
				{tt: TypeInteger, value: "3"},
				{tt: Rparen, value: ""},
			},
			msg:    "nested",
			expect: "12+3+",
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
			expect: "12+34-+",
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
			expect: "12>34if",
		},
	}

	for _, tt := range testcases {
		actual := Parse(tt.in)
		fmt.Printf("%v\n", actual)
		if actual.Text() != tt.expect {
			t.Errorf("[%s] expect: %s, actual: %s\n", tt.msg, tt.expect, actual.Text())
		}
	}
}
