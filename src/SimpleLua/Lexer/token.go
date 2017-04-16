package Lexer

import "fmt"

const (
	TokenAnd          string = "and"
	TokenDo                  = "do"
	TokenElse                = "else"
	TokenElseif              = "elseif"
	TokenEnd                 = "end"
	TokenFalse               = "false"
	TokenIf                  = "if"
	TokenLocal               = "local"
	TokenNil                 = "nil"
	TokenNot                 = "not"
	TokenOr                  = "or"
	TokenThen                = "then"
	TokenTrue                = "true"
	TokenWhile               = "while"
	TokenID                  = "<id>"
	TokenString              = "<string>"
	TokenNumber              = "<number>"
	TokenAdd                 = "+"
	TokenSub                 = "-"
	TokenMul                 = "*"
	TokenDiv                 = "/"
	TokenLen                 = "#"
	TokenLeftParen           = "("
	TokenRightParen          = ")"
	TokenAssign              = "="
	TokenSemicolon           = ";"
	TokenComma               = ","
	TokenEqual               = "=="
	TokenNotEqual            = "~="
	TokenLess                = "<"
	TokenLessEqual           = "<="
	TokenGreater             = ">"
	TokenGreaterEqual        = ">="
	TokenConcat              = ".."
	TokenEOF                 = "<eof>"
)

type Token struct {
	val interface{}
	typ string
	line int
	col int
}

func NewToken() *Token {
	token := new(Token)
	token.typ = TokenEOF
	return token
}

//debug
func (t *Token) String() string {
	var res string
	if t.typ == TokenNumber || t.typ == TokenID ||
		t.typ == TokenString {
		res = fmt.Sprintf("%v", t.val)
	} else {
		res = t.typ
	}
	return res
}

func (t *Token) clone() *Token {
	return &Token{
		val: t.val,
		typ: t.typ,
		line: t.line,
		col: t.col,
	}
}

func isKeyword(id string) bool {
	switch id {
	case TokenAnd, TokenDo, TokenElse, TokenElseif, TokenEnd,
		TokenFalse, TokenIf, TokenLocal, TokenNil, TokenNot,
		TokenOr, TokenThen, TokenTrue, TokenWhile:
		return true
	default:
		return false
	}
}