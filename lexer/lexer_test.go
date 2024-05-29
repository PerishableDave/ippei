package lexer

import (
	"ippei/token"
	"testing"
)

func TestReadRune(t *testing.T) {
	input := `(`

	l := New(input)
	l.readRune()
	if l.rn != '(' {
		t.Errorf("Expected %q, got %q", '(', l.rn)
	}
}

func TestReadIdentifier(t *testing.T) {
	input := `let`

	l := New(input)
	l.readRune()
	ident, err := l.readIdentifier()
	if ident != "let" || err != nil {
		t.Errorf("Expected %q, got %q", "foobar", ident)
	}
}

func TestReadString(t *testing.T) {
	input := `"foobar"`

	l := New(input)
	l.readRune()
	str := l.readString()
	if str != "foobar" {
		t.Errorf("Expected %q, got %q", "foobar", str)
	}
}

func TestNextToken(t *testing.T) {
	input := `let`

	l := New(input)
	l.readRune()
	tok := l.NextToken()
	if tok.Type != token.LET || tok.Literal != `let` {
		t.Errorf("Expected type: %q literal: %q, got type: %q literal: %q", token.LET, "let", tok.Type, tok.Literal)
	}
}
