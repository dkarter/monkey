package token

// TokenType represents the token type
type TokenType string

// Token represents a part of parsed code
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL represents illegal character
	ILLEGAL = "ILLEGAL"
	// EOF represents the end of file
	EOF = "EOF"

	// Identifiers + literals:

	// IDENT represents code indentation
	IDENT = "IDENT"
	// INT represents integer type numbers
	INT = "INT"

	// Operators:

	// ASSIGN represents assignment
	ASSIGN = "="
	// PLUS represents addition of two numbers
	PLUS = "+"

	// Delimiters:

	// COMMA represents a separator between two values
	COMMA = ","

	// SEMICOLON represents the end of a code block
	SEMICOLON = ";"

	// Closures:

	// LPAREN represents a left parenthesis closure
	LPAREN = "("
	// RPAREN represents a right parenthesis closure
	RPAREN = ")"
	// LBRACE represents a right curly brace closure
	LBRACE = "{"
	// RBRACE represents a right curly brace closure
	RBRACE = "}"

	// Keywords:

	// FUNCTION represents a function
	FUNCTION = "FUNCTION"
	// LET reprensts variable declaration
	LET = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent checks if a token is a known keyword or an identifier
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
