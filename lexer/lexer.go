package lexer

import (
	"github.com/amitsuthar69/interpreter/token"
)

type Lexer struct {
	input        string
	ch           byte // current char under examination
	position     int  // current position in input (points to current char 'ch') [index of ch in input]
	readPosition int  // current reading position in input (after current char) [see what comes up next]
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	/*
		check whether we have reached the end of input.
		If that’s the case it sets l.ch to 0, which is the ASCII code for the "NUL" character
		and signifies either “we haven’t read anything yet” or “end of file” for us.
	*/

	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	/*
		recognize whether the current character is a letter
		and if so, it needs to read the rest of the identifier/keyword
		until it encounters a non-letter-character.
	*/

	position := l.position

	// reads in an identifier and advances our lexer’s positions until it encounters a non-letter-character.
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	// helper function to check whether the given argument is a letter.
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	// helper function to check whether the given argument is a number.
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhiteSpaces() {
	// whitespace only acts as a separator of tokens and doesn’t have meaning, so we need to skip over it entirely
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	// Look at the current character under examination (l.ch) and return a token depending on which character it is.
	var tok token.Token

	l.skipWhiteSpaces()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier() // check for identifiers whenever the l.ch is not one of the recognized characters.
			tok.Type = token.LoopUpIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	/*
		Before returning the token we advance our pointers into the input
		so when we call NextToken() again the l.ch field is already updated.
	*/
	l.readChar()

	return tok
}
