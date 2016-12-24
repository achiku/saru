package main

import "testing"

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: Token{Type: LET, Literal: "let"},
				Name: &Identifier{
					Token: Token{Type: IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: Token{Type: IDENT, Literal: "yourVar"},
					Value: "yourVar",
				},
			},
		},
	}
	if actual := program.String(); actual != "let myVar = yourVar;" {
		t.Error(actual)
	}
}
