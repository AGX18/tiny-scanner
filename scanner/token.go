// Package scanner contains all the logic for scanning and tokenizing.
package scanner

import "fmt"

// TokenType represents the type of a token (e.g., IDENTIFIER, NUMBER, IF).
type TokenType int

// This is an 'enum' of all possible token types in the TINY language.
const (
	// Special tokens
	ILLEGAL TokenType = iota // Represents a token we don't recognize
	EOF                      // End of File

	// Literals
	IDENTIFIER // x, abc, my_var
	NUMBER     // 123, 289

	// Operators
	ASSIGN   // :=
	EQUAL    // =
	LESSTHAN // <
	PLUS     // +
	MINUS    // -
	MULT     // *
	DIV      // /

	// Delimiters
	OPENBRACKET   // (
	CLOSEDBRACKET // )
	SEMICOLON     // ; (From your PDF table [cite: 13])

	// Keywords
	IF     // if
	THEN   // then
	END    // end
	REPEAT // repeat
	UNTIL  // until
	READ   // read
	WRITE  // write
)

// keywords maps the string representation of keywords (reserved words) to their TokenType.
var keywords = map[string]TokenType{
	"if":     IF,
	"then":   THEN,
	"end":    END,
	"repeat": REPEAT,
	"until":  UNTIL,
	"read":   READ,
	"write":  WRITE,
}

// Token represents a single token scanned from the source code.
type Token struct {
	Type  TokenType // The type of token
	Value string    // The actual text of the token (e.g., "if", "x", "123")
	Line  int       // The line number it appeared on (good for errors)
}

// String returns a human-readable string representation of the TokenType.
// This is essential for the output file!
func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case IDENTIFIER:
		return "IDENTIFIER"
	case NUMBER:
		return "NUMBER"
	case ASSIGN:
		return "ASSIGN"
	case EQUAL:
		return "EQUAL"
	case LESSTHAN:
		return "LESSTHAN"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case MULT:
		return "MULT"
	case DIV:
		return "DIV"
	case OPENBRACKET:
		return "OPENBRACKET"
	case CLOSEDBRACKET:
		return "CLOSEDBRACKET"
	case SEMICOLON:
		return "SEMICOLON"
	case IF:
		return "IF"
	case THEN:
		return "THEN"
	case END:
		return "END"
	case REPEAT:
		return "REPEAT"
	case UNTIL:
		return "UNTIL"
	case READ:
		return "READ"
	case WRITE:
		return "WRITE"
	default:
		return fmt.Sprintf("Unknown(%d)", t)
	}
}
