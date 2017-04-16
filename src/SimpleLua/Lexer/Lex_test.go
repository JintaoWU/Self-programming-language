package Lexer

import (
	"testing"
	"fmt"
	"strings"
)

func TestScanner_Scan(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	code := strings.NewReader("local x = 'Hello, 世界'")
	sc := New(code)
	tok := sc.Scan()
	for tok.typ != TokenEOF {
		fmt.Println(tok," ", tok.typ)
		tok = sc.Scan()
	}
}
