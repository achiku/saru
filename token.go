package saru

// TokenType token type
type TokenType string

// Token token
type Token struct {
	Type    TokenType
	Literal string
}
