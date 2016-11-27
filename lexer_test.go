package monkey

import "testing"

func TestReadChar(t *testing.T) {
	cases := []struct {
		input string
	}{
		{input: "123456789011"},
		{input: "abcdefghijxx"},
		{input: "0"},
	}

	for _, c := range cases {
		l := &Lexer{input: c.input}
		for i := 0; i <= len(c.input)-1; i++ {
			l.readChar()
			if expected := string(c.input[i]); string(l.ch) != expected {
				t.Errorf("want %s got %s", expected, string(l.ch))
			}
		}
		l.readChar()
		if l.ch != 0 {
			t.Errorf("want %d got %s", 0, string(l.ch))
		}
	}
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	cases := []struct {
		expType    TokenType
		expLiteral string
	}{
		{ASSIGN, "="},
		{PLUS, "+"},
		{LPAREN, "("},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{RBRACE, "}"},
		{COMMA, ","},
		{SEMICOLON, ";"},
		{EOF, ""},
	}

	l := NewLexer(input)
	for i, c := range cases {
		tok := l.NextToken()
		if tok.Literal != c.expLiteral {
			t.Errorf("%d: want %s got %s", i, c.expLiteral, tok.Literal)
		}
		if tok.Type != c.expType {
			t.Errorf("%d: want %s got %s", i, c.expType, tok.Type)
		}
	}
}
