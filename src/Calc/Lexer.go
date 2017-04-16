package Calc

import (
	"strings"
	"text/scanner"
	"regexp"
	"errors"
)

const (
	Number = iota+1
	ADD
	SUB
	MUL
	DIV
	LB
	RB
)

type Token struct {
	tok string
	typ int
}

type Lexer struct {
	sc scanner.Scanner
	tokens []Token
}

func Init() *Lexer {
	lex := new(Lexer)
	return lex
}

func (lex *Lexer) Lex(code string) (error, []Token) {
	lex.sc.Init(strings.NewReader(code))
	var tokens []Token
	re := regexp.MustCompile(`^(0|[1-9][0-9]*)$`)
	for lex.sc.Scan() != scanner.EOF {
		tok := lex.sc.TokenText()
		switch {
		case re.MatchString(tok):
			tokens = append(tokens, Token{tok, Number})
		case tok == "+":
			tokens = append(tokens, Token{tok, ADD})
		case tok == "-":
			tokens = append(tokens, Token{tok, SUB})
		case tok == "*":
			tokens = append(tokens, Token{tok, MUL})
		case tok == "/":
			tokens = append(tokens, Token{tok, DIV})
		case tok == "(":
			tokens = append(tokens, Token{tok, LB})
		case tok == ")":
			tokens = append(tokens, Token{tok, RB})
		default:
			er := errors.New("there is something unrecognized")
			return er, nil
		}
	}
	return nil, tokens
}