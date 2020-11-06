package main

type TokenType string

const (
	None            = "NONE"
	Lparen          = "LPRN"
	Rparen          = "RPRN"
	Dot             = "DOT"
	Quote           = "QUOT"
	QuasiQuote      = "QQ"
	Unquote         = "UNQT"
	UnquoteSplicing = "UNQS"
	Eof             = "EOF"
	TypeString      = "STR"
	TypeInteger     = "INT"
	TypeRational    = "RAT"
	TypeReal        = "REAL"
	TypeComplex     = "CMPX"
	TypeSymbol      = "SYM"
	TypeBoolean     = "BOOL"
	Terminal        = "TERM"
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

func (s *Scanner) Char() string {
	return string(s.ch)
}

func Tokenize(program string) []*Token {
	s := NewScanner(program)
	tokens := make([]*Token, 0)
	for s.EachChar() {
		switch s.Char() {
		case "(":
			tokens = append(tokens, NewToken(Lparen, ""))
		case ")":
			tokens = append(tokens, NewToken(Rparen, ""))
		case "+":
			tokens = append(tokens, NewToken(TypeSymbol, "add"))
		case "-":
			tokens = append(tokens, NewToken(TypeSymbol, "sub"))
		case "*":
			tokens = append(tokens, NewToken(TypeSymbol, "mul"))
		case "/":
			tokens = append(tokens, NewToken(TypeSymbol, "div"))
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			tokens = append(tokens, NewToken(TypeInteger, s.Char()))
		case " ":
			break
		default:
			tokens = append(tokens, NewToken(None, s.Char()))
		}
	}
	return tokens
}
