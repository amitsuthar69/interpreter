### Notes

1. Lexical Analysis (Tokenization) : Source code --> Tokens

Tokenization is the process of breaking down input text into smaller, manageable units called 'Tokens'.

Tokens are small, easily categorizable data structures that are fed to the parser, which turns the tokens into an “Abstract Syntax Tree”.

A token is represented by a TokenType (a category label) and a Literal (the actual text).

Typical token types include:

- Keywords: Reserved words like let, fn.
- Identifiers: User-defined variable names.
- Literals: Values like numbers (INT) and strings.
- Operators: Symbols like +, -, \*, /.
- Delimiters: Symbols that separate code segments, e.g. ; , {} ()
