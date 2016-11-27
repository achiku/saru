package saru

// Lexer lexar
type Lexer struct {
	input   string
	currPos int
	nextPos int
	ch      byte
}

func (l *Lexer) readChar() {
	if l.nextPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPos]
	}
	l.currPos = l.nextPos
	l.nextPos++
}

// NextToken returns next token
func (l *Lexer) NextToken() Token {
	var tok Token

	switch l.ch {
	case '=':
		tok = newToken(ASSIGN, l.ch)
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case ',':
		tok = newToken(COMMA, l.ch)
	case '+':
		tok = newToken(PLUS, l.ch)
	case '{':
		tok = newToken(LBRACE, l.ch)
	case '}':
		tok = newToken(RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

// NewLexer new lexer
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}
