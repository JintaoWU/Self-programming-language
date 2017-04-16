package Calc

import ("testing"
	"fmt"
)

func TestLexer_Lex(t *testing.T) {
	code := "(1 + 3) + 56 * 345 / 56 -456"
	lex := Init()
	err, res := lex.Lex(code)
	if err != nil {
		t.Errorf("unrecognized word")
	}
	for _, item := range res {
		fmt.Println(item.tok,": ", item.typ)
	}

	code = "sdfghjk345t 4356 3456"
	err, res = lex.Lex(code)
	if err == nil {
		t.Errorf("wrong")
	}
	fmt.Println(err)
}
