package ui

import (
	"fmt"
)

type Node interface { node() }
type Object struct { Entries []Entry }

type Entry struct {
	Path []string
	Value Node
}

type String struct { Value string }
type Integer struct { Value int }
type Bool struct { Value bool }
type Array struct { Value []Node }

func (Object) node() {}
func (String) node() {}
func (Integer) node() {}
func (Bool) node() {}
func (Array) node() {}

type Parser struct {
	Current Token
	Lexer Lexer
}

func Parse(dat string) (Object, error) {
	l := Lexer{
		Content: dat,
		Pos: 0,
		Col: 0,
		Line: 1,
	};

	p := Parser{Lexer: l};
	p.Next();
	return p.ParseObject();
}

func (p *Parser) Next() {
	p.Current = p.Lexer.NextToken();
}

func (p *Parser) ExpectKind(kind TokenKind) bool {
	if p.Current.Kind != kind {
		return false
	}
	return true;
}

func (p *Parser) ExpectKindErr(kind TokenKind) error {
	if !p.ExpectKind(kind) {
		return p.InvalidToken(kind);
	}
	return nil;
}

func (p *Parser) InvalidToken(kind TokenKind) error {
	return fmt.Errorf("invalid token at: %v:%v, expected: `%v` but found `%v`\n", p.Lexer.Line, p.Lexer.Col, kind, p.Current.Kind);
}

func (p *Parser) ParseObject() (Object, error) {
	if err := p.ExpectKindErr(TokenBraceOpen); err != nil {
		return Object{}, err;
	}

	var obj Object;
	for p.ExpectKind(TokenBraceClose) {
		entry, err := p.ParseEntry();	
		if err != nil {
			return Object{}, err;
		}
		obj.Entries = append(obj.Entries, entry);
	}
	p.Next();
	return obj, nil;
}

func (p *Parser) ParseEntry() (Entry, error) {
	// TODO: remove while loop
    // TODO: this breaks nested objects
	for !p.ExpectKind(TokenComma) {
		if err := p.ExpectKindErr(TokenIdent); err != nil {
			return Entry{}, err;
		}

		/*
			iko,
			iko = { ... },
			iko.url = ... ,
		*/

        var entry Entry;
		switch p.Current.Kind {
		case TokenComma:
			break;
		case TokenAssign:
			n, err := p.ParseValue();
			if err != nil {
				return Entry{}, err;
			}
            entry.Value = n
            return entry, nil;
		case TokenDot:
            p.Next();
            if err := p.ExpectKindErr(TokenIdent); err != nil {
				return Entry{}, err;
            }
            entry.Path = append(entry.Path, p.Current.Value.(string));
            return entry, nil;
		default:
			return Entry{}, p.InvalidToken(p.Current.Kind);
		}
	}

	return Entry{}, nil;
};

func (p *Parser) ParseValue() (Node, error) {
	switch p.Current.Kind {
	case TokenString:
		return String{Value: p.Current.Value.(string)}, nil;
	case TokenInt:
		return Integer{Value: p.Current.Value.(int)}, nil;
	case TokenFalse:
		return Bool{Value: false}, nil;
	case TokenTrue:
		return Bool{Value: true}, nil;
	case TokenBracketOpen:
		return p.ParseObject();
	default:
		return nil, fmt.Errorf("TODO: put error message here");
	}
}

