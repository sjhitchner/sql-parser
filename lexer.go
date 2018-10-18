package main

import (
	"fmt"
	//"strings"

	. "github.com/sjhitchner/lexer"
)

const (
	TokenEOF TokenType = iota
	TokenCreate
	TokenTable
	TokenIndex
	TokenIfNotExists
	TokenLeftPara
	TokenRightPara
	TokenTableName
	TokenFieldName
	TokenFieldType
	TokenText

	sqlCreate        = "CREATE"
	sqlTable         = "TABLE"
	sqlIndex         = "INDEX"
	sqlIfNotExists   = "IF NOT EXISTS"
	sqlLeftPara      = "("
	sqlRightPara     = ")"
	sqlComma         = ","
	sqlSemiColon     = ";"
	sqlInteger       = "INTEGER"
	sqlVarChar       = "VARCHAR"
	sqlFloat         = "FLOAT"
	sqlPrimaryKey    = "PRIMARY KEY"
	sqlAutoIncrement = "AUTOINCREMENT"
	sqlText          = "TEXT"
	sqlNotNull       = "NOT NULL"
	sqlNull          = "NULL"
	sqlUnique        = "UNIQUE"
	sqlForeignKey    = "FOREIGN KEY"
	sqlReferences    = "REFERENCES"
	sqlOnUpdate      = "ON UPDATE"
	sqlCascade       = "CASCADE"
	sqlOnDelete      = "ON DELETE"
)

/*
func NewLexer(input string, initialState StateFunc) *Lexer
func (l *Lexer) Emit(tokenType TokenType)
func (l *Lexer) Errorf(format string, args ...interface{}) StateFunc
func (l *Lexer) Ignore() rune
func (l *Lexer) IgnoreUpTo(predicate RunePredicate) rune
func (l *Lexer) Next() rune
func (l *Lexer) NextToken() Token
func (l *Lexer) NextUpTo(predicate RunePredicate) rune
func (l *Lexer) Peek() rune
func (l *Lexer) Previous() rune
func (l *Lexer) PreviousToken() Token
*/

func lexCreate(l *Lexer) StateFunc {
	for {
		if l.Matches(sqlCreate) {
			fmt.Println("CREATE")
			l.Emit(TokenCreate)
			return lexCreateType
		}

		if l.Next() == EOF {
			break
		}
		l.Ignore()
	}

	// Correctly reached EOF.
	//if l.CurrentPosition > start {
	//	l.Emit(TokenText)
	//}
	l.Emit(TokenEOF) // Useful to make EOF a token.
	l.Emit(TokenEOF) // Useful to make EOF a token.
	l.Emit(TokenEOF) // Useful to make EOF a token.
	l.Emit(TokenEOF) // Useful to make EOF a token.
	return nil       // Stop the run
}

/*

	if FindToken(l, sqlCreate) {
		l.Emit(TokenCreate)
		return lexCreateType(l)
	}

	l.Emit(TokenEOF)
	return nil
}
*/

func lexCreateType(l *Lexer) StateFunc {
	l.Emit(TokenEOF) // Useful to make EOF a token.
	return nil
}

func lexIndex(l *Lexer) StateFunc {
	l.Emit(TokenEOF)
	return nil
}

func lexTable(l *Lexer) StateFunc {
	l.Emit(TokenEOF)
	return nil
}

func FindToken(l *Lexer, str string) bool {
	var position int
	for i := 0; i < 100; i++ {
		r := l.Next()

		fmt.Printf("%s = %s\n", string(r), string(str[position]))

		if r == rune(str[position]) {
			position++
		} else {
			position = 0
		}

		if position == len(str) {
			return true
		}
	}
	return false
}

/*
	var position int
	for i := 0; i < 100; i++ {
		r := l.Next()

		fmt.Printf("%s = %s\n", string(r), string(sqlCreate[position]))

		if r == rune(sqlCreate[position]) {
			position++
		} else {
			position = 0
		}

		if position == len(sqlCreate) {
			fmt.Println(sqlCreate)
		}
	}
	return nil
}

func lexTable(l *Lexer) StateFunc {
	return nil
}

/*
	for {
		if strings.HasPrefix(l.Input[l.CurrentPosition:], sqlCreate) {
			if l.CurrentPosition > l.start {
				l.Emit(TokenCreate)
			}
			return lexTable
		}
		if l.next() == eof {
			break
		}
	}

	// Correctly reached EOF.
	if l.pos > l.start {
		l.Emit(TokenText)
	}

	l.Emit(TokenEOF) // Useful to make EOF a token.
	return nil       // Stop the run loop.
}

func lexLeftMeta(l *Lexer) stateFn {
	l.pos += len(leftMeta)
	l.emit(TokenLeftMeta)
	return lexInsideAction // Now inside {{ }}.
}
*/

/*

// State Functions
//
// States
//  Alphanumeric String (possible Name)
//  Non Alphanumeric String (ignored text)
//  Text between <a..a> (hyperlink)
//

// Parses out alphanumeric names
func lexName(l *Lexer) stateFunction {
	for {
		if strings.HasPrefix(l.input[l.pos:], LEFT_ANGLE) {
			if l.pos > l.start {
				l.emit(TokenName)
			}
			return lexLeftHTML
		}

		if strings.HasPrefix(l.input[l.pos:], RIGHT_ANGLE) {
			return l.errorf("Invalid HTML Snippet")
		}

		if !IsAlphaNumeric(l.peek()) {
			if l.pos > l.start {
				l.emit(TokenName)
			}
			return lexText
		}

		r := l.next()
		if r == '\n' || r == '\r' {
			return l.errorf("Invalid HTML Snippet - has line break")
		} else if r == eof {
			break
		}
	}

	if l.pos > l.start {
		l.emit(TokenName)
	}

	l.emit(TokenEOF)
	return nil
}

// Parsers out non-name text
func lexText(l *Lexer) stateFunction {
	for {
		if strings.HasPrefix(l.input[l.pos:], LEFT_ANGLE) {
			if l.pos > l.start {
				l.emit(TokenText)
			}
			return lexLeftHTML
		}

		if IsAlphaNumeric(l.peek()) {
			if l.pos > l.start {
				l.emit(TokenText)
			}
			return lexName
		}

		r := l.next()
		if r == '\n' || r == '\r' {
			return l.errorf("Invalid HTML Snippet - has line break")
		} else if r == eof {
			break
		}
	}

	if l.pos > l.start {
		l.emit(TokenText)
	}

	l.emit(TokenEOF)
	return nil
}

func lexLeftHTML(l *Lexer) stateFunction {
	if strings.HasPrefix(l.input[l.pos:], LEFT_HYPER) {
		l.pos += len(LEFT_HYPER)
		return lexInsideHyperlink
	}
	l.pos += len(LEFT_ANGLE)
	return lexInsideHTML
}

func lexRightHTML(l *Lexer) stateFunction {
	l.pos += len(RIGHT_ANGLE)
	l.emit(TokenHTML)
	return lexName
}

func lexLeftHyperlink(l *Lexer) stateFunction {
	l.pos += len(LEFT_HYPER)
	return lexInsideHyperlink
}

func lexRightHyperlink(l *Lexer) stateFunction {
	l.pos += len(RIGHT_HYPER)
	l.emit(TokenHyperlink)
	return lexName
}

// Parsers out hyperlinks
func lexInsideHyperlink(l *Lexer) stateFunction {
	for {
		if strings.HasPrefix(l.input[l.pos:], LEFT_HYPER) {
			return l.errorf("Invalid Hyperlink Snippet")
		}

		if strings.HasPrefix(l.input[l.pos:], RIGHT_HYPER) {
			return lexRightHyperlink
		}

		r := l.next()
		if r == '\n' || r == '\r' {
			return l.errorf("Invalid HTML Snippet - has line break")
		} else if r == eof {
			return l.errorf("Invalid HTML Snippet - unclosed hyperlink")
			break
		}
	}
	return nil
}

func lexInsideHTML(l *Lexer) stateFunction {
	for {
		if strings.HasPrefix(l.input[l.pos:], LEFT_ANGLE) {
			return l.errorf("Invalid HTML Snippet")
		}

		if strings.HasPrefix(l.input[l.pos:], RIGHT_ANGLE) {
			return lexRightHTML
		}

		r := l.next()
		if r == '\n' || r == '\r' {
			return l.errorf("Invalid HTML Snippet - has line break")
		} else if r == eof {
			return l.errorf("Invalid HTML Snippet - unclosed tag")
			break
		}
	}
	return nil
}

*/
