package main

import (
	"fmt"
)

const (
	itemError itemType = iota
	itemEOF
	itemCreate
	itemIfNotExists
	itemLeftPara
	itemRightPara
	itemTableName
	itemFieldName
	itemFieldType
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

type stateFn func(*Lexer) stateFn

type Lexer struct {
	name  string    // used only for error reports.
	input string    // the string being scanned.
	start int       // start position of this item.
	pos   int       // current position in the input.
	width int       // width of last rune read from input.
	state stateFn   // current state
	items chan item // channel of scanned items.
}

// lex creates a new scanner for the input string.
func NewLexer(name, input string) *Lexer {
	l := &Lexer{
		name:  name,
		input: input,
		state: lexFindTable,
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

func (l *Lexer) Emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

// next returns the next rune in the input.
func (l *Lexer) next() (rune int) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	rune, l.width =
		utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return rune
}

// peek returns but does not consume
// the next rune in the input.
func (l *Lexer) peek() int {
	rune := l.next()
	l.backup()
	return rune
}

// ignore skips over the pending input before this point.
func (l *Lexer) ignore() {
	l.start = l.pos
}

// backup steps back one rune.
// Can be called only once per call of next.
func (l *Lexer) backup() {
	l.pos -= l.width
}

// accept consumes the next rune
// if it's from the valid set.
func (l *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *Lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}
