package main

import "fmt"

type TokenType string

const (
	None        = "NONE"
	Lparen      = "LPRN"
	Rparen      = "RPRN"
	TypeInteger = "INT"
	TypeOpr     = "OPR"
	TypeSymbol  = "SYM"
	Reserved    = "RES"
)

type Token struct {
	tt    TokenType
	value string
}

func NewToken(tt TokenType, value string) *Token {
	return &Token{tt: tt, value: value}
}

type Scanner struct {
	program string
	idx     int
	ch      byte
}

func NewScanner(program string) *Scanner {
	return &Scanner{
		program: program,
		idx:     0,
		ch:      0,
	}
}

func (s *Scanner) EachChar() bool {
	if s.idx == len(s.program) {
		return false
	}
	s.ch = s.program[s.idx]
	s.idx++

	return true
}

func (s *Scanner) decr(index int) {
	s.idx -= index
}

func (s *Scanner) NextChar() string {
	if s.idx+1 >= len(s.program) {
		return ""
	}
	return string(s.program[s.idx])
}

func (s *Scanner) GetWord() string {
	token := ""
	for s.EachChar() {
		switch s.Char() {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
			"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
			"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z":
			token += s.Char()
		default:
			fmt.Println(s.Char())
			s.decr(1)
			goto exit_loop
		}
	}
exit_loop:
	return token
}

func (s *Scanner) Char() string {
	return string(s.ch)
}

func Tokenize(program string) ([]*Token, error) {
	s := NewScanner(program)
	b := []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	fmt.Printf("%v\n", b)
	tokens := make([]*Token, 0)
	for s.EachChar() {
		switch s.Char() {
		case "(":
			tokens = append(tokens, NewToken(Lparen, "("))
		case ")":
			tokens = append(tokens, NewToken(Rparen, ")"))
		case "+":
			tokens = append(tokens, NewToken(TypeOpr, "add"))
		case "-":
			tokens = append(tokens, NewToken(TypeOpr, "sub"))
		case "*":
			tokens = append(tokens, NewToken(TypeOpr, "mul"))
		case "/":
			tokens = append(tokens, NewToken(TypeOpr, "div"))
		case ">":
			tokens = append(tokens, NewToken(TypeOpr, "gt"))
		case ">=":
			tokens = append(tokens, NewToken(TypeOpr, "gte"))
		case "<":
			tokens = append(tokens, NewToken(TypeOpr, "lt"))
		case "<=":
			tokens = append(tokens, NewToken(TypeOpr, "lte"))
		case "==":
			tokens = append(tokens, NewToken(TypeOpr, "eq"))
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			tokens = append(tokens, NewToken(TypeInteger, s.Char()))
		case "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
			"o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z":
			s.decr(1)
			token := s.GetWord()
			fmt.Printf("token: '%s'\n", token)
			tokens = append(tokens, NewToken(TypeSymbol, token))
		case " ":
			break
		default:
			tokens = append(tokens, NewToken(None, s.Char()))
		}
	}
	return tokens, nil
}
