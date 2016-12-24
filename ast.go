package main

import "bytes"

// Node ast node
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement ast statement
type Statement interface {
	Node
	statementNode()
}

// Expression ast expression
type Expression interface {
	Node
	expressionNode()
}

// Program program
type Program struct {
	Statements []Statement
}

// TokenLiteral token literal
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// Identifier ident
type Identifier struct {
	Token Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral for ident
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Token.Literal
}

// LetStatement let statement
type LetStatement struct {
	Token Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral for let statement
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// ReturnStatement return statement
type ReturnStatement struct {
	Token       Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral token literal
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	return out.String()
}

// ExpressionStatement expression statement
type ExpressionStatement struct {
	Token      Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral token literal
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
