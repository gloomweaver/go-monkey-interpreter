package lexer

import (
	"github.com/gloomweaver/go-monkey-interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func mapToken(ch byte) token.Token {
	switch ch {
	case '=':
		return token.Token{Type: token.ASSIGN, Literal: string(ch)}
	case ';':
		return token.Token{Type: token.SEMICOLON, Literal: string(ch)}
	case '(':
		return token.Token{Type: token.LPAREN, Literal: string(ch)}
	case ')':
		return token.Token{Type: token.RPAREN, Literal: string(ch)}
	case ',':
		return token.Token{Type: token.COMMA, Literal: string(ch)}
	case '+':
		return token.Token{Type: token.PLUS, Literal: string(ch)}
	case '{':
		return token.Token{Type: token.LBRACE, Literal: string(ch)}
	case '}':
		return token.Token{Type: token.RBRACE, Literal: string(ch)}
	case 0:
		return token.Token{Type: token.EOF, Literal: ""}
	default:
		return token.Token{Type: token.ILLEGAL, Literal: string(ch)}
	}
}

func (l *Lexer) NextToken() token.Token {
	tok := mapToken(l.ch)
	l.readChar()
	return tok
}
