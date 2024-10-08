package ast

import "phpgo/src/token"

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
	}
	return ""
}

type Identifier struct {
	Token token.Token // IDENT
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type VarStatement struct {
	Token token.Token // $
	Name  *Identifier
	Value Expression
}

func (s *VarStatement) TokenLiteral() string {
	return s.Token.Literal
}
func (s *VarStatement) statementNode() {}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (s *ReturnStatement) TokenLiteral() string {
	return s.Token.Literal
}
func (s *ReturnStatement) statementNode() {}
