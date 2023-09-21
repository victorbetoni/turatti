package parser

import (
	"os"
	"strconv"
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
	checkParserErrors(t, parser)
	if program == nil {
		t.Error("couldn't parse program")
		return
	}
	if len(program.Statements) != 3 {
		t.Errorf("expected 3 statements, got %d\n", len(program.Statements))
		return
	}

	tests := []struct {
		expectedIdentifier string
	}{{"x"}, {"y"}, {"z"}}

	for i, test := range tests {
		stmt := program.Statements[i]

		if stmt.TokenLiteral() != "def" {
			t.Errorf("wrong token. expected def got %q\n", stmt.TokenLiteral())
			return
		}

		defStmt, ok := stmt.(*ast.DefStatement)
		if !ok {
			t.Errorf("%T is not assignable from ast.DefStatement\n", stmt)
			return
		}

		if defStmt.Name.Value != test.expectedIdentifier {
			t.Errorf("expected %s for stmt identifier, got %s\n", test.expectedIdentifier, defStmt.Name.Value)
			return
		}

		if defStmt.Name.TokenLiteral() != test.expectedIdentifier {
			t.Errorf("expected %s for token literal, got %s\n", test.expectedIdentifier, defStmt.TokenLiteral())
			return
		}
	}

}

func TestReturnStatement(t *testing.T) {

	f, err := os.Open("../test_files/test_def_stmt.trt")
	if err != nil {
		t.Error("couldn't open def statement test file")
		return
	}

	lex := lexer.FromFile(f)
	parser := New(lex)

	program := parser.Parse()
	checkParserErrors(t, parser)
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d\n",
			len(program.Statements))
	}
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("%T is not assignable from ast.ReturnStatement\n", stmt)
			return
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("TokenLiteral not 'return', got %q\n", returnStmt.TokenLiteral())
			return
		}
	}
}

func TestIdentifier(t *testing.T) {
	psr := New(lexer.New("foo;"))
	program := psr.Parse()
	checkParserErrors(t, psr)

	if len(program.Statements) != 1 {
		t.Fatalf("expected 1 statement, got %d\n", len(program.Statements))
	}

	smt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("first statement is not assignable from ExpressionStatement. got %T\n", program.Statements[0])
	}

	identifier, ok := smt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression is not assignable from Identifier. got %T\n", smt)
	}

	if identifier.Value != "foo" {
		t.Errorf("identifier value is not %s. got %s", "foo\n", identifier.Value)
	}

	if identifier.TokenLiteral() != "foo" {
		t.Errorf("identifier token literal is not %s. got %s\n", "foo", identifier.TokenLiteral())
	}

}

func TestIntegerLiteralExpression(t *testing.T) {
	l := lexer.New("5;")
	parser := New(l)
	program := parser.Parse()
	checkParserErrors(t, parser)

	if len(program.Statements) != 1 {
		t.Fatalf("expected 1 statement, got %d\n", len(program.Statements))
	}

	smt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("the first statement is not assignable from ExpressionStatement, got %T\n", program.Statements[0])
	}

	literal, ok := smt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("expression not assignable from IntegerLiteral, got %T\n", smt.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("expected value %d, got %d\n", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("expected token literal %s, got %s\n", "5", literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	tests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"-20;", "-", 20},
		{"!34;", "!", 34},
	}

	for _, tt := range tests {
		parser := New(lexer.New(tt.input))
		program := parser.Parse()
		checkParserErrors(t, parser)

		if len(program.Statements) != 1 {
			t.Fatalf("expected %d statements, got %d\n", 1, len(program.Statements))
		}

		smt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("statement is not assignable from ExpressionStatement. got %T", program.Statements[0])
		}

		exp, ok := smt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("expression is not assignable from PrefixExpression. got %T", smt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("wrong operator. expected '%s' got '%s'", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp, tt.integerValue) {
			return
		}

	}
}

func testIntegerLiteral(t *testing.T, exp ast.Expression, expected int64) bool {
	il, ok := exp.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("the expression '%s' does not return a integer literal.", exp.String())
		return false
	}

	if il.Value != expected {
		t.Errorf("expected integer value %d, got %d", expected, il.Value)
		return false
	}

	if il.TokenLiteral() != strconv.Itoa(int(expected)) {
		t.Errorf("integer token literal '%s' is not the value expected '%s", il.TokenLiteral(), strconv.Itoa(int(expected)))
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors\n", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q\n", msg)
	}
	t.FailNow()
}
