package scanner

import (
	"reflect" // We need this to compare the token slices
	"testing"
)

func TestScanTokens(t *testing.T) {
	// Here we define our "table" of test cases
	testCases := []struct {
		name     string  // A name for this specific test
		input    string  // The source code we're scanning
		expected []Token // The list of tokens we expect to get
	}{
		// --- Test Case 1: Simple Assignment ---
		{
			name:  "Simple Assignment",
			input: "x := 10;",
			expected: []Token{
				{Type: IDENTIFIER, Value: "x", Line: 1},
				{Type: ASSIGN, Value: ":=", Line: 1},
				{Type: NUMBER, Value: "10", Line: 1},
				{Type: SEMICOLON, Value: ";", Line: 1},
				{Type: EOF, Value: "EOF", Line: 1},
			},
		},

		// --- Test Case 2: Keywords, Operators, and Whitespace ---
		{
			name:  "Keywords and Operators",
			input: "if (x < 1) then write x end",
			expected: []Token{
				{Type: IF, Value: "if", Line: 1},
				{Type: OPENBRACKET, Value: "(", Line: 1},
				{Type: IDENTIFIER, Value: "x", Line: 1},
				{Type: LESSTHAN, Value: "<", Line: 1},
				{Type: NUMBER, Value: "1", Line: 1},
				{Type: CLOSEDBRACKET, Value: ")", Line: 1},
				{Type: THEN, Value: "then", Line: 1},
				{Type: WRITE, Value: "write", Line: 1},
				{Type: IDENTIFIER, Value: "x", Line: 1},
				{Type: END, Value: "end", Line: 1},
				{Type: EOF, Value: "EOF", Line: 1},
			},
		},

		// --- Test Case 3: Comments, Newlines, and Line Numbers ---
		{
			name: "Comments and Newlines",
			// We use backticks (`) for a multi-line string
			input: `
read x;
/* this is a
   multi-line comment */
write x
`,
			expected: []Token{
				// Note: Line numbers start at 2 because of the first newline
				{Type: READ, Value: "read", Line: 2},
				{Type: IDENTIFIER, Value: "x", Line: 2},
				{Type: SEMICOLON, Value: ";", Line: 2},
				// The comment is skipped, and line numbers are tracked
				{Type: WRITE, Value: "write", Line: 5},
				{Type: IDENTIFIER, Value: "x", Line: 5},
				{Type: EOF, Value: "EOF", Line: 6}, // EOF is on the last line
			},
		},

		// --- Test Case 4: All Single/Multi-Char Operators ---
		{
			name:  "All Operators",
			input: "+ - * / = < () ; :=",
			expected: []Token{
				{Type: PLUS, Value: "+", Line: 1},
				{Type: MINUS, Value: "-", Line: 1},
				{Type: MULT, Value: "*", Line: 1},
				{Type: DIV, Value: "/", Line: 1},
				{Type: EQUAL, Value: "=", Line: 1},
				{Type: LESSTHAN, Value: "<", Line: 1},
				{Type: OPENBRACKET, Value: "(", Line: 1},
				{Type: CLOSEDBRACKET, Value: ")", Line: 1},
				{Type: SEMICOLON, Value: ";", Line: 1},
				{Type: ASSIGN, Value: ":=", Line: 1},
				{Type: EOF, Value: "EOF", Line: 1},
			},
		},
	}

	// --- The Test Runner ---
	// This loop runs one sub-test for each case in our table
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. Run the scanner
			scanner := NewScanner(tc.input)
			tokens := scanner.ScanTokens()

			// 2. Compare the actual tokens to the expected tokens
			if !reflect.DeepEqual(tokens, tc.expected) {
				// 3. If they don't match, fail the test and print a helpful diff
				t.Errorf("Test failed for input:\n%s\n", tc.input)
				t.Errorf("Expected tokens:\n%+v\n", tc.expected)
				t.Errorf("Got tokens:\n%+v\n", tokens)
			}
		})
	}
}
