package lexer

import (
	"fmt"
	"ippei/token"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	rn           rune // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readRune() {
	var size int
	if l.readPosition >= len(l.input) {
		l.rn = 0
	} else {
		l.rn, size = utf8.DecodeRuneInString(l.input[l.readPosition:])
	}
	l.position = l.readPosition
	l.readPosition += size
}

func (l *Lexer) readIdentifier() (string, error) {
	position := l.position

	if unicode.IsLetter(l.rn) {
		l.readRune()
	} else {
		return "", fmt.Errorf("invalid identifier")
	}

	for unicode.IsLetter(l.rn) || unicode.IsDigit(l.rn) {
		l.readRune()
	}

	return l.input[position:l.position], nil
}

func (l *Lexer) readString() string {
	if l.rn == '"' {
		l.readRune()
	}

	position := l.position
	for l.rn != '"' && l.rn != 0 {
		l.readRune()
	}

	str := l.input[position:l.position]

	if l.rn == '"' {
		l.readRune()
	}

	return str
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	var err error

	switch l.rn {
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: string(l.rn)}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: string(l.rn)}
	case '"':
		tok = token.Token{Type: token.STRING, Literal: l.readString()}
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		if unicode.IsLetter(l.rn) {
			tok.Literal, err = l.readIdentifier()
			if err != nil {
				tok.Type = token.ILLEGAL
				return tok
			}
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			tok = token.Token{Type: token.ILLEGAL, Literal: string(l.rn)}
		}
	}

	return tok
}
