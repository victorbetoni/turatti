package token

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	FUNCTION  = "FUNCTION"
	DEF       = "DEF"
)

var keywords = map[string]TokenType{
	"fun": FUNCTION,
	"def": DEF,
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

func NewToken(tokenType TokenType, ch rune, line, column int) Token {
	return Token{
		Type:    tokenType,
		Literal: string(ch),
		Line:    line,
		Column:  column,
	}
}

func FindKeywordOrIdent(literal string) TokenType {
	if tok, ok := keywords[literal]; ok {
		return tok
	}
	return IDENT
}
