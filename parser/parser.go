package parser

import (
	"fmt"
	"turatti/ast"
	"turatti/lexer"
	"turatti/token"
)

type Parser struct {
	lex          *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	errors       []string
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

func New(lex *lexer.Lexer) *Parser {
	p := &Parser{
		lex:    lex,
		errors: []string{},
	}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for p.currentToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.DEF:
		return p.parseDefStmt()
	case token.IF:
		return p.parseIfStmt()
	case token.ELSE:
		return p.parseDefStmt()
	default:
		return nil
	}
}

func (p *Parser) parseDefStmt() ast.Statement {

	stmt := &ast.DefStatement{Token: p.currentToken}

	if !p.expectToken(token.IDENT) {
		p.peekError(token.IDENT, p.peekToken, p.lex.FileName)
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectToken(token.ASSIGN) {
		p.peekError(token.ASSIGN, p.peekToken, p.lex.FileName)
	}

	p.nextToken()

	expression := []token.Token{}

	for p.currentToken.Type != token.SEMICOLON {
		expression = append(expression, p.currentToken)
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseFunctionStmt() ast.Statement {
	return nil
}

func (p *Parser) parseIfStmt() ast.Statement {
	return nil
}

func (p *Parser) parseElseStmt() ast.Statement {
	return nil
}

func (p *Parser) expectToken(tok token.TokenType) bool {

	if p.peekToken.Type == tok {
		p.nextToken()
		return true
	}
	return false

}

func (p *Parser) peekError(tok token.TokenType, token token.Token, file string) {
	p.errors = append(p.errors, fmt.Sprintf("%s: unexpected token at: line %d column %d. expected %s instead.",
		p.lex.FileName, p.currentToken.Line, p.currentToken.Column, tok))
}
