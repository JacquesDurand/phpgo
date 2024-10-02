package parser

import (
	"phpgo/src/ast"
	"phpgo/src/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
$x = 5;
$y = 10;
$fooBar = 838383;
`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"fooBar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testVarStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testVarStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "$" {
		t.Errorf("s.TokenLiteral not '$'. got=%q", s.TokenLiteral())
		return false
	}
	varStatement, ok := s.(*ast.VarStatement)
	if !ok {
		t.Errorf("s not *ast.VarStatement. got=%T", s)
		return false
	}
	if varStatement.Name.Value != name {
		t.Errorf("varStatement.Name.Value not '%s'. got=%s", name, varStatement.Name.Value)
		return false
	}
	if varStatement.Name.TokenLiteral() != name {
		t.Errorf("varStatement.Name.TokenLiteral() not '%s'. got=%s",
			name, varStatement.Name.TokenLiteral())
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
