package main

import "testing"

func TestSyntaxError(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let 838383;
	`
	p := NewParser(NewLexer(input))
	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}
	if len(p.errors) != 1 {
		t.Errorf("want 1 got %d", len(p.errors))
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	if len(p.Errors()) == 0 {
		return
	}
	for _, msg := range p.Errors() {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	p := NewParser(NewLexer(input))
	checkParserErrors(t, p)

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

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 8383;
	`
	p := NewParser(NewLexer(input))
	checkParserErrors(t, p)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParserProgram() returned nil")
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ReturnStatement)
		if !ok {
			t.Errorf("stmt not ReturnStatement got=%T", stmt)
		}
		if actual := returnStmt.TokenLiteral(); actual != "return" {
			t.Errorf("returnStmt.TokenLiteral not `return` got=%q", actual)
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	p := NewParser(NewLexer(input))
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if actual := len(program.Statements); actual != 1 {
		t.Fatalf("program.Statements needs to be 1 got %d ", actual)
	}
	stmt, ok := program.Statements[0].(*ExpressionStatement)
	if !ok {
		t.Fatalf("want ExpressionStatement got %T", program.Statements[0])
	}
	ident, ok := stmt.Expression.(*Identifier)
	if !ok {
		t.Fatalf("want Identifier got %T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("want foobar got %s", ident.Value)
	}
	if actual := ident.TokenLiteral(); actual != "foobar" {
		t.Errorf("want foobar got %s", ident.Value)
	}
}
