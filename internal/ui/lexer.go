package ui

import (
	"strings"
	"unicode"
)

type TokenKind int
const (
	TokenIdent TokenKind = iota
	TokenBraceOpen
	TokenBraceClose
	TokenBracketOpen
	TokenBracketClose
	TokenAssign
	TokenDot
	TokenComma
	TokenTrue
	TokenFalse
	TokenString
	TokenInt
	// TokenFloat
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

func (l *Lexer) NextToken() Token {
	l.SkipWhitespace();

	var t Token;
	t.Line = l.Line;
	t.Col = l.Col;

	c := l.Next();
	switch c {
	case 0:
		t.Kind = TokenEof;
	case '{':
		t.Kind = TokenBraceOpen;
	case '}':
		t.Kind = TokenBraceClose;
	case '[':
		t.Kind = TokenBracketOpen;
	case ']':
		t.Kind = TokenBracketClose;
	case '=':
		t.Kind = TokenAssign;
	case ',':
		t.Kind = TokenComma;
	case '.':
		t.Kind = TokenDot;
	case '"':
		var tstr strings.Builder;
		for l.Pos <= len(l.Content) {
			c := l.Next()
			if c == '"' {
				break;
			}

			tstr.WriteByte(c);
		}
		t.Kind = TokenString;
		t.Value = tstr.String();
	default:
		if unicode.IsDigit(rune(c)) {
			return l.Int(c);
		}

		if l.IsValidIdent(c) {
			return l.Ident(c);
		}
	}
	return t;
}

func (l *Lexer) Int(begin byte) Token {
	var num strings.Builder;
	num.WriteByte(begin);

	for unicode.IsDigit(rune(l.Peek())) {
		num.WriteByte(l.Next());
	}

	var t Token;
	t.Kind = TokenInt;

	return t;
}

func (l *Lexer) IsValidIdent(c byte) bool {
	if unicode.IsLetter(rune(c)) || c == '_' || c == '-' {
		return true;
	}
	return false;
}

func (l *Lexer) Ident(begin byte) Token {
	var ident strings.Builder;
	ident.WriteByte(begin);
	for {
		c := l.Peek();

		if l.IsValidIdent(c) || unicode.IsDigit(rune(c)) {
			ident.WriteByte(l.Next());
			continue;
		}
		break;
	}
	value := ident.String();

	var t Token;
	t.Kind = TokenIdent;
	t.Value = value;
	t.Line = l.Line;
	t.Col = l.Col;

	switch value {
	case "true":
		t.Kind = TokenTrue;
	case "false":
		t.Kind = TokenFalse;
	}

	return t;
}

func (l *Lexer) Next() byte {
	c := l.Peek();

	l.Pos++;
	if c == '\n' {
		l.Line++;
		l.Col = 1;
	} else {
		l.Col++
	}

	return c;
}

func (l *Lexer) Peek() byte {
	if l.Pos >= len(l.Content) {
		return 0;
	}
	return l.Content[l.Pos];
}

func (l *Lexer) SkipWhitespace() {
	for {
		switch l.Peek() {
		case ' ', '\t', '\r', '\n':
			l.Next()
		default:
			return
		}
	}
}

