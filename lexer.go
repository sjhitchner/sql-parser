package main

import (
//"fmt"
)

const (
	itemError itemType = iota
	itemEOF
	itemCreate
	itemName
	itemType
)

type itemType int

type item struct {
	typ   itemType
	value string
}

func (i item) String() string {
	switch i.typ {
	case itemEOF:
		return "EOF"
	case itemError:
		return i.value
	}

	if len(i.value) > 10 {
		return fmt.Sprintf("%.10q...", i.value)
	}

	return fmt.Sprintf("%q", i.value)
}

type stateFn func(*lexer) stateFn

type Lexer struct {
	name  string    // used only for error reports.
	input string    // the string being scanned.
	start int       // start position of this item.
	pos   int       // current position in the input.
	width int       // width of last rune read from input.
	items chan item // channel of scanned items.
}

func run() {
	for state := startState; state != nil; {
		state = state(lexer)
	}
}

// lex creates a new scanner for the input string.
func NewLexer(name, input string) *lexer {
	l := &Lexer{
		name:  name,
		input: input,
		state: lexText,
		items: make(chan item, 2), // Two items sufficient.
	}
	return l
}

func (l *Lexer) NextItem() item {
	for {
		select {
		case item := <-l.items:
			return item
		default:
			l.state = l.state(l)
		}
	}
	panic("not reached")
}

/*
func (l *lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.items) // No more tokens will be delivered.
}
*/

func (l *lexer) Emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}
