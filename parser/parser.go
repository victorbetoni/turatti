package parser

import (
	"turatti/ast"
	"turatti/lexer"
	"turatti/token"
)

type Parser struct {
	lex          *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

func New(lex *lexer.Lexer) *Parser {
	p := &Parser{lex: lex}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Parse() *ast.Program {
	return nil
}
