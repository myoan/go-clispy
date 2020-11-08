package main

import (
	"strings"
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
		expect string
	}{
		{
			in:     "()",
			msg:    "empty",
			expect: "( )",
		},
		{
			in:     "(+ 1 2)",
			msg:    "simple add",
			expect: "( add 1 2 )",
		},
		{
			in:     "(+ 1 (* 2 3))",
			msg:    "complex add",
			expect: "( add 1 ( mul 2 3 ) )",
		},
		{
			in:     "(> 1 2)",
			msg:    "compare",
			expect: "( gt 1 2 )",
		},
		{
			in:     "(if (> 0 3) 1 2)",
			msg:    "if stmt",
			expect: "( if ( gt 0 3 ) 1 2 )",
		},
		{
			in:     "(defun incr (n) (+ n 1)) (incr 2)",
			msg:    "defun stmt",
			expect: "( defun incr ( n ) ( add n 1 ) ) ( incr 2 )",
		},
	}

	for _, tt := range testcases {
		tokens, _ := Tokenize(tt.in)
		vals := make([]string, 0)
		for tokens.Next() {
			vals = append(vals, tokens.token.value)
		}
		actual := strings.Join(vals, " ")
		if actual != tt.expect {
			t.Errorf("[%s] expect: '%s', actual: '%s'\n", tt.msg, tt.expect, actual)
		}
	}
}
