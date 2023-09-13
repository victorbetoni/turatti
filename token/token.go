package token

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSING    = "="
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
)

type TokenType string

type Token struct {
	Type     TokenType
	Literal  string
	FileName string
	Line     int64
	Column   int64
}
