package ast

import "github.com/amitsuthar69/interpreter/token"

/*
Every node in our AST has to implement the Node interface,
meaning it has to provide a TokenLiteral() method that
returns the literal value of the token it’s associated with.
*/

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

/*
statementNode and expressionNode are not strictly necessary but they help us by guiding the Go
compiler and possibly causing it to throw errors when we use a Statement where an Expression
should’ve been used, and vice versa.
*/

type Program struct {
	Statements []Statement // Every valid Monkey program is a series of statements.
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	/*
		Let statments would have these fields on them :
		- Name of the variable (the identifier)
		- Value on the right side of equal sign (the expression that produces the value)
		- the token the AST is associated with (so we can implement the tokenLiteral method)
	*/
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
