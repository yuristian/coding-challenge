package lexer

type TokenType int

const (
	LBRACE TokenType = iota // {
	RBRACE                  // }
	EOF
	ILLEGAL
	COLON
	STRING
	COMMA
	NUMBER
	TRUE
	FALSE
	NULL
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

func (l *Lexer) readString() string {
	start_pos := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[start_pos:l.position]
}

func (l *Lexer) NextToken() Token {
	var token Token
	l.skipWhitespace()
	switch l.ch {
	case '{':
		token.Type = LBRACE
		token.Literal = "{"
		l.readChar()
	case '}':
		token.Type = RBRACE
		token.Literal = "}"
		l.readChar()
	case '"':
		token.Type = STRING
		token.Literal = l.readString()
		l.readChar()
	case ':':
		token.Type = COLON
		token.Literal = ":"
		l.readChar()
	case ',':
		token.Type = COMMA
		token.Literal = ","
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

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}