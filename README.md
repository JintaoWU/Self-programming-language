[![Build Status](https://travis-ci.org/plutolove/Self-programming-language.svg?branch=master)](https://travis-ci.org/plutolove/Self-programming-language)
### Self Programming Language
1. _Calc is expression evaluator with LL(1).
It supports a variety of operations(+, -, *, /, ())._

>* LL(1) Grammar
>>* expr ::= expr `+` term | expr `-` term | term
>>* term ::= term `*` factor | term `/` factor | factor
>>* factor ::= `(`expr`)` | NUM

