package main

// Node ast node
type Node interface {
	TokenLiteral() string
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
