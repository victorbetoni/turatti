package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"
	INT     = "INT"

	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	ASSIGN      = "="
	PLUS        = "+"
	BANG        = "!"
	ASTERISK    = "*"
	SLASH       = "/"
	MINUS       = "-"
	EQ          = "=="
	NOT_EQ      = "!="
	SLASH_EQ    = "/="
	ASTERISK_EQ = "*="
	PLUS_EQ     = "+="
	MINUS_EQ    = "-="
	INCREMENT   = "++"
	DECREMENT   = "--"

	LESSTHAN      = "<"
	GREATERTHAN   = ">"
	LESSEQTHAN    = "<="
	GREATEREQTHAN = ">="

	FUNCTION = "FUNCTION"
	DEF      = "DEF"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fun":    FUNCTION,
	"def":    DEF,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
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

func NewComposableToken(tokenType TokenType, ch rune, ch2 rune, line, column int) Token {
	return Token{
		Type:    tokenType,
		Literal: string(ch) + string(ch2),
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
