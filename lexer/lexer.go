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
	currentRune   rune
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
		lexer.readRune()
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

func (lexer *Lexer) readRune() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.currentRune = 0
	} else {
		lexer.currentRune = []rune(lexer.input)[lexer.readPosition]
		lexer.currentColumn++
	}
	lexer.position = lexer.readPosition
	lexer.readPosition++
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.eatWhitespace()

	switch lexer.currentRune {
	case '=':
		if lexer.peekChar() == '=' {
			current := lexer.currentRune
			lexer.readRune()
			tok = token.NewComposableToken(token.EQ, current, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
		} else {
			tok = token.NewToken(token.ASSIGN, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
		}
	case ';':
		tok = token.NewToken(token.SEMICOLON, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
	case '(':
		tok = token.NewToken(token.LPAREN, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
	case ')':
		tok = token.NewToken(token.RPAREN, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
	case ',':
		tok = token.NewToken(token.COMMA, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
	case '+':
		if lexer.peekChar() == '=' {
			current := lexer.currentRune
			lexer.readRune()
			tok = token.NewComposableToken(token.PLUS_EQ, current, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
		} else {
			tok = token.NewToken(token.PLUS, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
		}
	case '{':
		tok = token.NewToken(token.LBRACE, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
	case '}':
		tok = token.NewToken(token.RBRACE, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
	case '*':
		if lexer.peekChar() == '=' {
			current := lexer.currentRune
			lexer.readRune()
			tok = token.NewComposableToken(token.ASTERISK_EQ, current, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
		} else {
			tok = token.NewToken(token.ASTERISK, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
		}
	case '<':
		if lexer.peekChar() == '=' {
			current := lexer.currentRune
			lexer.readRune()
			tok = token.NewComposableToken(token.LESSEQTHAN, current, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
		} else {
			tok = token.NewToken(token.LESSTHAN, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
		}
	case '>':
		if lexer.peekChar() == '=' {
			current := lexer.currentRune
			lexer.readRune()
			tok = token.NewComposableToken(token.GREATEREQTHAN, current, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
		} else {
			tok = token.NewToken(token.GREATERTHAN, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
		}
	case '-':
		if lexer.peekChar() == '=' {
			current := lexer.currentRune
			lexer.readRune()
			tok = token.NewComposableToken(token.MINUS_EQ, current, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
		} else {
			tok = token.NewToken(token.MINUS, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
		}
	case '/':
		if lexer.peekChar() == '=' {
			current := lexer.currentRune
			lexer.readRune()
			tok = token.NewComposableToken(token.SLASH_EQ, current, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
		} else {
			tok = token.NewToken(token.SLASH, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
		}
	case '!':
		if lexer.peekChar() == '=' {
			current := lexer.currentRune
			lexer.readRune()
			tok = token.NewComposableToken(token.NOT_EQ, current, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
		} else {
			tok = token.NewToken(token.BANG, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Column = lexer.currentColumn
		tok.Line = lexer.currentLine
	default:
		if isLetter(lexer.currentRune) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.FindKeywordOrIdent(tok.Literal)
			return tok
		} else if isDigit(lexer.currentRune) {
			tok.Literal = lexer.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, lexer.currentRune, lexer.currentLine, lexer.currentColumn)
		}
	}

	lexer.readRune()
	return tok
}

func (lexer *Lexer) peekChar() rune {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.getRuneAt(lexer.readPosition)
	}
}

func (lexer *Lexer) readIdentifier() string {
	currentPosition := lexer.position
	for isLetter(lexer.currentRune) {
		lexer.readRune()
	}
	return lexer.input[currentPosition:lexer.position]
}

func (lexer *Lexer) readNumber() string {
	currentPosition := lexer.position
	for isDigit(lexer.currentRune) {
		lexer.readRune()
	}
	return lexer.input[currentPosition:lexer.position]
}

func isLetter(ch rune) bool {
	return ('a' <= ch && 'z' >= ch) || ('A' <= ch && 'Z' >= ch) || ch == '_'
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func (lexer *Lexer) eatWhitespace() {
	for lexer.currentRune == ' ' || lexer.currentRune == '\t' || lexer.currentRune == '\r' || lexer.currentRune == '\n' {
		if lexer.currentRune == '\r' || lexer.currentRune == '\n' {
			lexer.currentLine++
			lexer.currentColumn = 1
		} else {
			lexer.currentColumn++
		}
		lexer.readRune()
	}

}
