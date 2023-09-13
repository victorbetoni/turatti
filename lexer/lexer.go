package lexer

import (
	"io"
	"os"
	"turatti/token"
)

type Lexer struct {
	position      int
	readPosition  int
	currentLine   int
	currentColumn int
	current       rune
	input         string
	fileName      string
}

func New(file *os.File) *Lexer {
	if result, err := io.ReadAll(file); err == nil {
		lexer := &Lexer{
			input:         string(result),
			fileName:      file.Name(),
			position:      -1,
			readPosition:  0,
			currentLine:   1,
			currentColumn: 1,
		}
		lexer.read()
		return lexer
	}
	return &Lexer{}
}

func (lexer *Lexer) getRuneAt(pos int) rune {
	if pos >= int(len(lexer.input)) {
		return 0
	}
	return []rune(lexer.input)[pos]
}

func (lexer *Lexer) read() {
	if lexer.current == '\\' && lexer.getRuneAt(lexer.readPosition) == 'n' {
		lexer.currentLine++
		lexer.currentColumn = 1
	}
	if lexer.readPosition >= len(lexer.input) {
		lexer.current = 0
	} else {
		lexer.current = []rune(lexer.input)[lexer.readPosition]
		lexer.currentColumn++
	}
	lexer.position = lexer.readPosition
	lexer.readPosition++
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	switch lexer.current {
	case '=':
		tok = token.NewToken(token.ASSIGN, lexer.current, lexer.currentLine, lexer.currentColumn)
	case ';':
		tok = token.NewToken(token.SEMICOLON, lexer.current, lexer.currentLine, lexer.currentColumn)
	case '(':
		tok = token.NewToken(token.LPAREN, lexer.current, lexer.currentLine, lexer.currentColumn)
	case ')':
		tok = token.NewToken(token.RPAREN, lexer.current, lexer.currentLine, lexer.currentColumn)
	case ',':
		tok = token.NewToken(token.COMMA, lexer.current, lexer.currentLine, lexer.currentColumn)
	case '+':
		tok = token.NewToken(token.PLUS, lexer.current, lexer.currentLine, lexer.currentColumn)
	case '{':
		tok = token.NewToken(token.LBRACE, lexer.current, lexer.currentLine, lexer.currentColumn)
	case '}':
		tok = token.NewToken(token.RBRACE, lexer.current, lexer.currentLine, lexer.currentColumn)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Column = lexer.currentColumn
		tok.Line = lexer.currentLine
	}

	lexer.read()
	return tok
}
