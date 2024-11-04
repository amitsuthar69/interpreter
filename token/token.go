package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LoopUpIdent(ident string) TokenType {
	/*
		LookupIdent checks the keywords table to see whether the given identifier is in fact a keyword.
		If it is, it returns the keyword’s TokenType constant. If it isn’t, we just get back token.IDENT,
		which is the TokenType for all user-defined identifiers.
	*/
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, etc
	INT   = "INT"   // 12345

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	MUL      = "*"
	DIV      = "/"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	EQ     = "=="
	NOT_EQ = "!="

	GT = ">"
	LT = "<"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "fn"
	LET      = "let"
	TRUE     = "true"
	FALSE    = "false"
	IF       = "if"
	ELSE     = "else"
	RETURN   = "return"
)
