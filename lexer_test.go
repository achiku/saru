package saru

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

func TestNextTokenSimple(t *testing.T) {
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

func TestNextTokenComplex(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
    return true;
} else {
    return false;
}
`
	cases := []struct {
		expType    TokenType
		expLiteral string
	}{
		{LET, "let"},
		{IDENT, "five"},
		{ASSIGN, "="},
		{INT, "5"},
		{SEMICOLON, ";"},
		{LET, "let"},
		{IDENT, "ten"},
		{ASSIGN, "="},
		{INT, "10"},
		{SEMICOLON, ";"},
		{LET, "let"},
		{IDENT, "add"},
		{ASSIGN, "="},
		{FUNCTION, "fn"},
		{LPAREN, "("},
		{IDENT, "x"},
		{COMMA, ","},
		{IDENT, "y"},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{IDENT, "x"},
		{PLUS, "+"},
		{IDENT, "y"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},
		{SEMICOLON, ";"},
		{LET, "let"},
		{IDENT, "result"},
		{ASSIGN, "="},
		{IDENT, "add"},
		{LPAREN, "("},
		{IDENT, "five"},
		{COMMA, ","},
		{IDENT, "ten"},
		{RPAREN, ")"},
		{SEMICOLON, ";"},
		{BANG, "!"},
		{MINUS, "-"},
		{SLASH, "/"},
		{ASTERISK, "*"},
		{INT, "5"},
		{SEMICOLON, ";"},
		{INT, "5"},
		{LT, "<"},
		{INT, "10"},
		{GT, ">"},
		{INT, "5"},
		{SEMICOLON, ";"},
		{IF, "if"},
		{LPAREN, "("},
		{INT, "5"},
		{LT, "<"},
		{INT, "10"},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{RETURN, "return"},
		{TRUE, "true"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},
		{ELSE, "else"},
		{LBRACE, "{"},
		{RETURN, "return"},
		{FALSE, "false"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},
		{EOF, ""},
	}

	l := NewLexer(input)
	for i, c := range cases {
		tok := l.NextToken()
		if tok.Literal != c.expLiteral {
			t.Errorf("[%d] want %s got %s", i, c.expLiteral, tok.Literal)
		}
		if tok.Type != c.expType {
			t.Errorf("[%d] want %s got %s", i, c.expType, tok.Type)
		}
	}
}
