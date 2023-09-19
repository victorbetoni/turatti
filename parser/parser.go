package parser

import (
	"fmt"
	"turatti/ast"
	"turatti/lexer"
	"turatti/token"
)

const (
	_      int = iota
	LOWEST     // lowest precedence
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL // highest precedence
)

type (
	prefixParser func() ast.Expression
	infixParser  func(ast.Expression) ast.Expression
)

type Parser struct {
	lex           *lexer.Lexer
	currentToken  token.Token
	peekToken     token.Token
	errors        []string
	prefixParsers map[token.TokenType]prefixParser
	infixParsers  map[token.TokenType]infixParser
}

func (p *Parser) registerPrefixParser(token token.TokenType, parser prefixParser) {
	p.prefixParsers[token] = parser
}

func (p *Parser) registerInfixParser(token token.TokenType, parser infixParser) {
	p.infixParsers[token] = parser
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
		lex:           lex,
		errors:        []string{},
		prefixParsers: make(map[token.TokenType]prefixParser),
		infixParsers:  make(map[token.TokenType]infixParser),
	}
	p.nextToken()
	p.nextToken()
	p.registerPrefixParser(token.IDENT, p.parseIdentifier)
	return p
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
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
		return p.parseDefStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	smt := &ast.ExpressionStatement{Token: p.currentToken}
	smt.Expression = p.parseExpression(LOWEST)
	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}
	return smt
}

func (p *Parser) parseDefStatement() ast.Statement {

	stmt := &ast.DefStatement{Token: p.currentToken}

	if !p.expectToken(token.IDENT) {
		p.peekError(token.IDENT, p.peekToken, p.lex.FileName)
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectToken(token.ASSIGN) {
		p.peekError(token.ASSIGN, p.peekToken, p.lex.FileName)
		return nil
	}

	p.nextToken()

	expression := []token.Token{}

	for p.currentToken.Type != token.SEMICOLON {
		expression = append(expression, p.currentToken)
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.DefStatement{Token: p.currentToken}

	p.nextToken()

	expression := []token.Token{}
	for p.currentToken.Type != token.SEMICOLON {
		expression = append(expression, p.currentToken)
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefixParser := p.prefixParsers[p.currentToken.Type]
	if prefixParser == nil {
		return nil
	}
	leftExpression := prefixParser()
	return leftExpression
}

func (p *Parser) expectToken(tok token.TokenType) bool {

	if p.peekToken.Type == tok {
		p.nextToken()
		return true
	}
	return false

}

func (p *Parser) peekError(tok token.TokenType, token token.Token, file string) {
	p.errors = append(p.errors, fmt.Sprintf("%s: unexpected token %s at: line %d column %d. expected %s instead.",
		p.lex.FileName, p.currentToken.Type, p.currentToken.Line, p.currentToken.Column, tok))
}
