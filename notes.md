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

---

2. Parsing (Syntactic Analysis) : Tokens --> AST

Wikipedia: A parser is a software component that takes input data (frequently text) and builds a data structure – often some kind of parse tree,
abstract syntax tree or other hierarchical structure – giving a structural representation of the input, checking for correct syntax in the process.
Thus the process of parsing is also called syntactic analysis.

The word “abstract” is based on the fact that certain details visible in the source code are omitted in the AST.
Semicolons, newlines, whitespace, comments, braces, bracket and parentheses – depending on the language and the parser these details
are not represented in the AST, but merely guide the parser when constructing it.

-> There are two main strategies when parsing a programming language:

Top-down parsing or Bottom-up parsing.

A lot of slightly different forms of each strategy exist. For example, “recursive descent parsing”, “Early parsing” or “predictive parsing” are all variations of top down parsing.

The parser we are going to write is a recursive descent parser. And in particular, it’s a “top down operator precedence” parser, sometimes called “Pratt parser”, after its inventor Vaughan Pratt.

The difference between top down and bottom up parsers is that:
The former starts with constructing root node of the AST and then descends while the latter does it the other way around.
