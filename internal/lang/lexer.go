package lang

import (
	"fmt"
	"os"
)

type TokenKind int
const (
	TokenIdent TokenKind = iota
	TokenBraceOpen
	TokenBraceClose
	TokenBracketOpen
	TokenBracketClose
	TokenEquals
	TokenString
	TokenDot
	TokenComma
	TokenTrue
	TokenFalse
	TokenEof
)

type Token struct {
	Kind TokenKind
	Value any
	Col int
	Line int
}

type Lexer struct {
	Content string
	Line int
	Col int
	Pos int
}


func Parse(path string) error {
	dat, err := os.ReadFile(path);
	if err != nil {
		return err;
	}

	l := Lexer{
		Content: string(dat),
		Pos: 0,
		Col: 0,
		Line: 1,
	};

	for {
		tok := l.NextToken();

		if tok.Kind == TokenEof {
			break;	
		}
	}


	return nil;
}

func (l *Lexer) NextToken() Token {

	var t Token;

	c := l.Next();
	switch c {
	case '\n':
	case '{':
		t.Kind = TokenBraceOpen;
	case '}':
		t.Kind = TokenBraceClose;
	case '[':
		t.Kind = TokenBracketOpen;
	case ']':
		t.Kind = TokenBracketClose;
	case '=':
		t.Kind = TokenEquals;
	case ',':
		t.Kind = TokenComma;
	case '.':
		t.Kind = TokenDot;
	default:

	}

	return t;
}

func (l *Lexer) Next() byte {
	if l.Pos >= len(l.Content) {
		return 0;
	}

	c := l.Content[l.Pos];
	for c == ' ' || c == '\t' || c == '\r' {
		l.Pos++;
	}
	if c == '\n' {
		l.Line++;
		l.Col = 0;
	} else {
		l.Col++;
	}
	l.Pos++;
	return l.Content[l.Pos];
}

