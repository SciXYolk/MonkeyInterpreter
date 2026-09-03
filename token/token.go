// token/token.go

package token

//Token type is defined as string
type TokenType string

type Token struct {
	Type	TokenType
	Literal string
}

// ILLEGAL means tokens or characters you dont know about
// EOF means end of file which means it tells our parser it can stop
const (
	ILLEGAL = "ILLEGAL"
	EOF		= "EOF"

	// Identifier + literals
	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"   // 1343456


	// Operators
	ASSIGN	= "="
	PLUS	= "+"

	// Delimiters
	COMMA	= ","
	SEMICOLON	= ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET		 = "LET"
)

var keywords = map[string] TokenType {
	"fn": FUNCTION,
	"let": LET,
}

// checks  the keyword table to see if the given identifier is a keyword
func LookupIdent(ident string) TokenType{
	if tok, ok := keywords[ident]; ok {
		return tok 
	}
	return IDENT
}