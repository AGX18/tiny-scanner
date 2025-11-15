package scanner

import "unicode"

// Scanner holds the state of our scanner.
type Scanner struct {
	source  string  // The full source code string
	tokens  []Token // The list of tokens we've scanned
	start   int     // Start of the current token being scanned
	current int     // Current position in the source string
	line    int     // Current line number
}

// NewScanner creates a new scanner instance.
func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		line:   1, // Start at line 1
	}
}

// ScanTokens is the main function that loops through the source code
// and generates all tokens.
func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		// We are at the beginning of the next token
		s.start = s.current
		s.scanToken()
	}

	// Add one final EOF token
	s.tokens = append(s.tokens, Token{Type: EOF, Value: "EOF", Line: s.line})
	return s.tokens
}

// isAtEnd checks if we've consumed all characters.
func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

// scanToken identifies and processes a single token.
func (s *Scanner) scanToken() {
	char := s.advance()

	switch char {
	// Single-character tokens
	case '(':
		s.addToken(OPENBRACKET)
	case ')':
		s.addToken(CLOSEDBRACKET)
	case ';':
		s.addToken(SEMICOLON)
	case '+':
		s.addToken(PLUS)
	case '-':
		s.addToken(MINUS)
	case '*':
		s.addToken(MULT)
	case '<':
		s.addToken(LESSTHAN)
	case '=':
		s.addToken(EQUAL)

	// Two-character (or one) tokens
	case ':':
		if s.match('=') {
			s.addToken(ASSIGN)
		} else {
			s.addToken(ILLEGAL) // Or handle as an error
		}

	// Comments or Division
	case '/':
		if s.match('*') {
			// This is a comment, scan until we find "*/"
			s.comment()
		} else {
			s.addToken(DIV)
		}

	// Whitespace (skip)
	case ' ', '\r', '\t':
		// Ignore whitespace
	case '\n':
		s.line++ // Newline, increment line counter

	// Default: Identifiers, Numbers, or Errors
	default:
		if isDigit(char) {
			s.number()
		} else if isAlpha(char) {
			s.identifier()
		} else {
			// Unrecognized character
			s.addToken(ILLEGAL)
		}
	}
}

// --- Helper Functions ---

// advance consumes one character from the source and returns it.
func (s *Scanner) advance() rune {
	s.current++
	return rune(s.source[s.current-1])
}

// addToken creates a new token from the current 'start' and 'current' pointers.
func (s *Scanner) addToken(tokenType TokenType) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{Type: tokenType, Value: text, Line: s.line})
}

// match checks if the *next* character is the one we expect.
// If it is, it consumes it and returns true.
func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}
	if rune(s.source[s.current]) != expected {
		return false
	}
	// It matches, so consume the character
	s.current++
	return true
}

// peek looks at the current character without consuming it.
func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return rune(0) // Null rune
	}
	return rune(s.source[s.current])
}

// peekNext looks at the character *after* the current one.
func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return rune(0)
	}
	return rune(s.source[s.current+1])
}

// comment consumes a multi-line comment block.
func (s *Scanner) comment() {
	for !s.isAtEnd() {
		if s.peek() == '*' && s.peekNext() == '/' {
			// End of comment
			s.advance() // Consume the '*'
			s.advance() // Consume the '/'
			return
		}
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
	// If we get here, the file ended without a closing "*/"
	// You could report an error, but for this project, we just stop.
}

// number consumes a number token.
func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}
	// Note: We are NOT supporting decimals like 12.5 based on PDF examples.
	// If you need decimals, we'd add logic here to look for a '.'
	s.addToken(NUMBER)
}

// identifier consumes an identifier or keyword token.
func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	// After consuming, check if it's a keyword or a normal identifier
	text := s.source[s.start:s.current]
	tokenType, isKeyword := keywords[text]

	if isKeyword {
		s.addToken(tokenType) // It's a keyword (if, then, etc.)
	} else {
		s.addToken(IDENTIFIER) // It's a user-defined identifier (x, abc)
	}
}

// --- Character Type Checkers ---

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func isAlpha(r rune) bool {
	return unicode.IsLetter(r)
}

func isAlphaNumeric(r rune) bool {
	return isAlpha(r) || isDigit(r)
}
