package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT            = "IDENT" // add, foobar, x, y, ...
	AND              = "AND"
	BEGIN            = "BEGIN"
	CASE             = "CASE"
	COND             = "COND"
	DEFINE           = "DEFINE"
	DELAY            = "DELAY"
	DO               = "DO"
	ELSE             = "ELSE"
	IF               = "IF"
	LAMBDA           = "LAMBDA"
	LET              = "LET"
	LETSTAR          = "LET*"
	LETREC           = "LETREC"
	OR               = "OR"
	QUASIQUOTE       = "QUASIQUOTE"
	QUOTE            = "QUOTE"
	SET              = "SET!"
	UNQUOTE          = "UNQUOTE"
	UNQUOTE_SPLICING = "UNQUOTE-SPLICING"

	INT = "INT" // 123456

	LPAREN = "("
	RPAREN = ")"

	STRING = "STRING"
	NUMBER = "NUMBER"
)

func LookupIdent(ident string) TokenType {
	switch ident {
	case "fn":
		return "FUNCTION"
	case "let":
		return "LET"
	case "true":
		return "TRUE"
	case "false":
		return "FALSE"
	case "if":
		return "IF"
	case "else":
		return "ELSE"
	case "return":
		return "RETURN"
	default:
		return "IDENT"
	}
}
