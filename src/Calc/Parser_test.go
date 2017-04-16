package Calc

import (
	"testing"
)

func TestCalc(t *testing.T) {
	code := "(-1 - 3) * 6 / 2 "
	lex := Init()
	_, res := lex.Lex(code)

	e := InitExp(res, len(res))

	//fmt.Println(e.evalue())

	if e.evalue() != -12 {
		t.Errorf("Wrong answer\n")
	}
}
