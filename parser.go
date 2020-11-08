package main

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

type TokenList struct {
	idx    int
	tokens []*Token
	token  *Token
}

func NewTokenList() *TokenList {
	return &TokenList{
		idx:    0,
		tokens: make([]*Token, 0),
		token:  nil,
	}
}

func (tl *TokenList) Push(t *Token) {
	tl.tokens = append(tl.tokens, t)
}

func (tl *TokenList) Next() bool {
	if len(tl.tokens) <= tl.idx {
		return false
	}
	tl.token = tl.tokens[tl.idx]
	tl.idx++
	return true
}

func (tl *TokenList) LastToken() *Token {
	return tl.tokens[len(tl.tokens)-1]
}

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

func Tokenize(program string) (*TokenList, error) {
	s := NewScanner(program)
	tl := NewTokenList()
	for s.EachChar() {
		switch s.Char() {
		case "(":
			tl.Push(NewToken(Lparen, "("))
		case ")":
			tl.Push(NewToken(Rparen, ")"))
		case "+":
			tl.Push(NewToken(TypeOpr, "add"))
		case "-":
			tl.Push(NewToken(TypeOpr, "sub"))
		case "*":
			tl.Push(NewToken(TypeOpr, "mul"))
		case "/":
			tl.Push(NewToken(TypeOpr, "div"))
		case ">":
			tl.Push(NewToken(TypeOpr, "gt"))
		case ">=":
			tl.Push(NewToken(TypeOpr, "gte"))
		case "<":
			tl.Push(NewToken(TypeOpr, "lt"))
		case "<=":
			tl.Push(NewToken(TypeOpr, "lte"))
		case "==":
			tl.Push(NewToken(TypeOpr, "eq"))
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			tl.Push(NewToken(TypeInteger, s.Char()))
		case "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
			"o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z":
			s.decr(1)
			token := s.GetWord()
			if tl.LastToken().tt == Lparen {
				tl.Push(NewToken(TypeOpr, token))
			} else {
				tl.Push(NewToken(TypeSymbol, token))
			}
		case " ":
			break
		default:
			tl.Push(NewToken(None, s.Char()))
		}
	}
	return tl, nil
}
