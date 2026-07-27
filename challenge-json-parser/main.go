package main

import (
	"json-parser/lexer"
	"os"
)

func main() {
	// filename := os.Args[1]

	// 2. Read the entire file content into memory
	// content, err := os.ReadFile(filename)
	// if err != nil {
	// 	fmt.Printf("Error reading file '%s': %v\n", filename, err)
	// 	os.Exit(1)
	// }

	// lex := lexer.NewLexer(string(content))
	lex := lexer.NewLexer("  {}  ")
	result := parse(lex)
	if result {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func parse(l *lexer.Lexer) bool {
	token1 := l.NextToken()
	if token1.Type != lexer.LBRACE {
		return false
	}

	token2 := l.NextToken()
	if token2.Type != lexer.RBRACE {
		return false
	}

	token3 := l.NextToken()
	if token3.Type != lexer.EOF {
		return false
	}

	return true
}