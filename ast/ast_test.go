package ast

import (
	"testing"
	"turatti/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&DefStatement{
				Token: token.Token{Type: token.DEF, Literal: "def"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "x"},
					Value: "x",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "y"},
					Value: "y",
				},
			},
		},
	}
	if program.String() != "def x = y;" {
		t.Errorf("wrong statement. expected 'def x = y;' but got '%s'", program.String())
	}
}
