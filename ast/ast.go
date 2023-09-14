package ast

import "turatti/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type DefStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *DefStatement) statementNode() {}

func (ls *DefStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type IfStatement struct {
	Token token.Token
	Value Expression
}

func (ls *IfStatement) statementNode() {}
