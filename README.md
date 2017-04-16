Self-programming-language
=
>Calc is expression evaluator with LL(1).
It supports a variety of operations(+, -, *, /, ()).

>>* LL(1) Grammar

>>>* expr ::= expr `+` term | expr `-` term | term
>>>* term ::= term `*` factor | term `/` factor | factor
>>>* factor ::= `(`expr`)` | NUM

