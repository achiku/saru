package main

import "testing"

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	p := NewParser(NewLexer(input))
	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Errorf("want 3 got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		testLetStatement(t, stmt, tt.expectedIdentifier)
	}
}

func testLetStatement(t *testing.T, s Statement, name string) {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not `let` got=%s", s.TokenLiteral())
	}

	letStmt, ok := s.(*LetStatement)
	if !ok {
		t.Errorf("s not LetStatement. got=%T", s)
	}

	if actual := letStmt.Name.Value; actual != name {
		t.Errorf("want %s got %s", name, actual)
	}

	if actual := letStmt.Name.TokenLiteral(); actual != name {
		t.Errorf("want %s got %s", name, actual)
	}
}
