package lexer

type TokenType int

const (
	LBRACE TokenType = iota // {
	RBRACE                  // }
	EOF
	ILLEGAL
)

type Token struct {
	Type    TokenType
	Literal string
}

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{
		input:        input,
		position:     0,
		readPosition: 0,
	}
	lexer.readChar()

	return lexer
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition = l.readPosition + 1
}

func (l *Lexer) NextToken() Token {
	var token Token
	switch l.ch {
	case '{':
		token.Type = LBRACE
		token.Literal = "{"
		l.readChar()
	case '}':
		token.Type = RBRACE
		token.Literal = "}"
		l.readChar()
	case 0:
		token.Type = EOF
		token.Literal = ""
	default:
		token.Type = ILLEGAL
		token.Literal = string(l.ch)
		l.readChar()
	}

	return token
}