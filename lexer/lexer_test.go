package lexer

import (
	"testing"
	"../token"
)

func TestNextToken(t *testing.T) {
	input := `=+(),;`
	tests := []struct {
		exceptedType 	token.TokenType
		exceptedLiteral string
	}{
		{token.ASSIGN, "="},
	}

	l := New(input)

	for i, tt := range tests {

		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("type")
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("literal")
		}
	}
}