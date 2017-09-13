package lexer

import (
	"testing"
	"../token"
)

func TestNextToken(t *testing.T) {
	input := `=+(),;{}`
	tests := []struct {
		expectedType 	token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {

		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] tokentype - expected=%q actual=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] literal - expected=%q actual=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}