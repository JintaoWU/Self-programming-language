package Calc

import (
	"strconv"
	"fmt"
)

type Exp struct {
	tokens []Token
	curr int
	size int
}

func InitExp(tk []Token, size int) Exp {
	res := Exp{tk, 0, size}
	return res
}

func (e *Exp) check() bool {
	return e.curr >= e.size
}

func (e *Exp) nextToken() {
	e.curr ++
}

func (e *Exp) getToken() Token {
	return e.tokens[e.curr]
}

func (e *Exp) evalue() float64 {
	t1 := e.term()
	for !e.check() && (e.getToken().typ == ADD || e.getToken().typ == SUB) {
		tok := e.getToken()
		e.nextToken()
		if tok.typ == ADD {
			t1 += e.term()
		}
		if tok.typ == SUB {
			t1 -= e.term()
		}
	}
	return t1
}

func (e *Exp) term() float64 {
	res := e.factor()
	for !e.check() && (e.getToken().typ == MUL || e.getToken().typ == DIV) {
		tok := e.getToken()
		e.nextToken()
		if tok.typ == MUL {
			res *= e.factor()
		}
		if tok.typ == DIV {
			res /= e.factor()
		}
	}
	return res
}

func (e *Exp) factor() float64 {
	if e.check() {
		return 0
	}

	if e.getToken().typ == SUB {
		e.nextToken()
		return -e.factor()
	}else if e.getToken().typ == LB {
		e.nextToken()
		val := e.evalue()
		e.nextToken()
		return val
	}else {
		tok := e.getToken()
		res, _ := strconv.ParseFloat(tok.tok, 64)
		e.nextToken()
		return res
	}
}

func Calc(tk []Token) float64 {
	e := InitExp(tk, len(tk))
	fmt.Println(e.size,"  ", e.curr)
	return e.evalue()
}
