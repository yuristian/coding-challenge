package main

import (
	"fmt"
	"json-parser/lexer"
	"os"
)

func main() {
	filename := os.Args[1]
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file '%s': %v\n", filename, err)
		os.Exit(1)
	}

	lex := lexer.NewLexer(string(content))
	result := parseObject(lex)
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

func parseObject(l *lexer.Lexer) bool {
	tok := l.NextToken()
	if tok.Type != lexer.LBRACE {
		return false
	}

	// var result bool = true

	for true {
		tok = l.NextToken()
		if tok.Type != lexer.STRING {
			return false
		}

		tok = l.NextToken()
		if tok.Type != lexer.COLON {
			return false
		}

		tok = l.NextToken()
		if tok.Type != lexer.STRING {
			return false
		}

		tok = l.NextToken()
		if tok.Type == lexer.RBRACE {
			break
		} else if tok.Type == lexer.COMMA {
		} else {
			return false
		}
	}

	tok = l.NextToken()
	if tok.Type != lexer.EOF {
		return false
	}

	return true
}