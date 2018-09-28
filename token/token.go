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

	// TRUE true
	TRUE = "TRUE"

	// FALSE false
	FALSE = "FALSE"

	// IF conditional
	IF = "IF"

	// ELSE conditional
	ELSE = "ELSE"

	// RETURN returns a value from a function
	RETURN = "RETURN"

	// Operators:

	// ASSIGN represents assignment into a custom identifier
	ASSIGN = "="

	// PLUS used for number addition and concatenation (strings, arrays)
	PLUS = "+"

	// MINUS used for number subtraction
	MINUS = "-"

	// BANG used for negation and flipping boolean values
	BANG = "!"

	// ASTERISK used for number multipication
	ASTERISK = "*"

	// SLASH used for number division
	SLASH = "/"

	// LT short for Less Than, used for number comparison
	LT = "<"

	// GT short for Greater Than, used for number comparison
	GT = ">"

	// EQ returns boolean if left && right arguments are equal
	EQ = "=="

	// NOTEQ returns boolean if left && right arguments are not equal
	NOTEQ = "!="
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"true":   TRUE,
	"false":  FALSE,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent checks if a token is a known keyword or an identifier
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
