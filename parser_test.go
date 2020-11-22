package main

import (
	"testing"
)

func TestScanner_EachChar(t *testing.T) {
	var ch byte
	s := NewScanner("hoge")
	ret := s.EachChar()
	if !ret {
		t.Errorf("return false")
		return
	}
	if string(s.ch) != "h" {
		t.Errorf("string not match: expect 'h', but actual: '%s'", string(ch))
		return
	}

	ret = s.EachChar()
	if !ret {
		t.Errorf("return false")
		return
	}
	if string(s.ch) != "o" {
		t.Errorf("string not match: expect 'o', but actual: '%s'", string(ch))
		return
	}

	ret = s.EachChar()
	if !ret {
		t.Errorf("return false")
		return
	}
	if string(s.ch) != "g" {
		t.Errorf("string not match: expect 'g', but actual: '%s'", string(ch))
		return
	}

	ret = s.EachChar()
	if !ret {
		t.Errorf("return false")
		return
	}
	if string(s.ch) != "e" {
		t.Errorf("string not match: expect 'e', but actual: '%s'", string(ch))
		return
	}

	ret = s.EachChar()
	if ret {
		t.Errorf("return true")
		return
	}
}

func TestScanner_GetWord(t *testing.T) {
	testcases := []struct {
		in     string
		expect string
	}{
		{
			in:     "hoge",
			expect: "hoge",
		},
		{
			in:     "hoge1",
			expect: "hoge1",
		},
		{
			in:     "hoge+1",
			expect: "hoge",
		},
		{
			in:     "hoge\n1",
			expect: "hoge",
		},
	}

	for _, tt := range testcases {
		s := NewScanner(tt.in)
		actual := s.GetWord()
		if actual != tt.expect {
			t.Errorf("expect: '%s', actual: '%s'\n", tt.expect, actual)
		}
	}
}

func TestTokenize(t *testing.T) {
	testcases := []struct {
		in     string
		msg    string
		expect *TokenList
	}{
		{
			in:  "()",
			msg: "empty",
			expect: &TokenList{
				idx:   0,
				token: nil,
				tokens: []*Token{
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeRParen, value: ""},
				},
			},
		},
		{
			in:  "(+ 1 2)",
			msg: "simple add",
			expect: &TokenList{
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
		},
		{
			in:  "(+ 10 22)",
			msg: "multi digit",
			expect: &TokenList{
				idx:   0,
				token: nil,
				tokens: []*Token{
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeAdd, value: ""},
					{tt: TokenTypeInt, value: "10"},
					{tt: TokenTypeInt, value: "22"},
					{tt: TokenTypeRParen, value: ""},
				},
			},
		},
		{
			in:  "(+ 1 (* 2 3))",
			msg: "complex add",
			expect: &TokenList{
				idx:   0,
				token: nil,
				tokens: []*Token{
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeAdd, value: ""},
					{tt: TokenTypeInt, value: "1"},
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeMul, value: ""},
					{tt: TokenTypeInt, value: "2"},
					{tt: TokenTypeInt, value: "3"},
					{tt: TokenTypeRParen, value: ""},
					{tt: TokenTypeRParen, value: ""},
				},
			},
		},
		{
			in:  "(> 1 2)",
			msg: "compare",
			expect: &TokenList{
				idx:   0,
				token: nil,
				tokens: []*Token{
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeGt, value: ""},
					{tt: TokenTypeInt, value: "1"},
					{tt: TokenTypeInt, value: "2"},
					{tt: TokenTypeRParen, value: ""},
				},
			},
		},
		{
			in:  "(if (> 0 3) 1 2)",
			msg: "if stmt",
			expect: &TokenList{
				idx:   0,
				token: nil,
				tokens: []*Token{
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeKeyword, value: "if"},
					{tt: TokenTypeLParen, value: ""},
					{tt: TokenTypeGt, value: ""},
					{tt: TokenTypeInt, value: "0"},
					{tt: TokenTypeInt, value: "3"},
					{tt: TokenTypeRParen, value: ""},
					{tt: TokenTypeInt, value: "1"},
					{tt: TokenTypeInt, value: "2"},
					{tt: TokenTypeRParen, value: ""},
				},
			},
		},
		{
			in:  "(defun incr (n) (+ n 1)) (incr 2)",
			msg: "defun stmt",
			expect: &TokenList{
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
					{tt: TokenTypeInt, value: "2"},
					{tt: TokenTypeRParen, value: ""},
				},
			},
		},
	}

	for _, tt := range testcases {
		tl, _ := Tokenize(tt.in)
		for i, actual := range tl.tokens {
			if actual.tt != tt.expect.tokens[i].tt {
				t.Errorf("[%s](%d) expect: '%s', actual: '%s'\n", tt.msg, i, tt.expect.tokens[i].tt, actual.tt)
			}
		}
	}
}
