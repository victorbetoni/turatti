package lexer

import (
	"os"
	"testing"
	"turatti/token"
)

func TestNextToken(t *testing.T) {
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.DEF, "def"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.DEF, "def"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.DEF, "def"},
		{token.IDENT, "sum"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fun"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.DEF, "def"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "sum"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.NOT_EQ, "!="},
		{token.INT, "4"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.DEF, "def"},
		{token.IDENT, "val"},
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.INT, "10"},
		{token.SLASH, "/"},
		{token.INT, "4"},
		{token.RPAREN, ")"},
		{token.GREATERTHAN, ">"},
		{token.INT, "6"},
		{token.SEMICOLON, ";"},

		{token.DEF, "def"},
		{token.IDENT, "vale"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.LESSTHAN, "<"},
		{token.INT, "8"},
		{token.SEMICOLON, ";"},

		{token.DEF, "def"},
		{token.IDENT, "str"},
		{token.ASSIGN, "="},
		{token.STRING, "hello world"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	file, err := os.Open("../test_files/lexer_test.trt")
	if err != nil {
		t.Fatalf("couldn't open lexer_test.trt file wrong")
	}

	lexer := FromFile(file)

	for i, tt := range tests {
		tok := lexer.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

}
