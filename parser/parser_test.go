package parser

import (
	"os"
	"testing"
	"turatti/ast"
	"turatti/lexer"
)

func TestDefStatement(t *testing.T) {

	f, err := os.Open("../test_files/test_def_stmt.trt")
	if err != nil {
		t.Error("couldn't open def statement test file")
		return
	}

	lex := lexer.FromFile(f)
	parser := New(lex)

	program := parser.Parse()
	if program == nil {
		t.Error("couldn't parse program")
		return
	}
	if len(program.Statements) != 3 {
		t.Errorf("expected 3 statements, got %d", len(program.Statements))
		return
	}

	tests := []struct {
		expectedIdentifier string
	}{{"x"}, {"y"}, {"z"}}

	for i, test := range tests {
		stmt := program.Statements[i]

		if stmt.TokenLiteral() != "def" {
			t.Errorf("wrong token. expected def got %q", stmt.TokenLiteral())
			return
		}

		defStmt, ok := stmt.(*ast.DefStatement)
		if !ok {
			t.Errorf("%T is not assignable from ast.DefStatement", stmt)
			return
		}

		if defStmt.Name.Value != test.expectedIdentifier {
			t.Errorf("expected %s for stmt identifier, got %s", test.expectedIdentifier, defStmt.Name.Value)
			return
		}

		if defStmt.Name.TokenLiteral() != test.expectedIdentifier {
			t.Errorf("expected %s for token literal, got %s", test.expectedIdentifier, defStmt.TokenLiteral())
			return
		}
	}

}
