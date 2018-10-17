package main

import (
	//"fmt"
	"strings"
)

const (
	sqlCreate        = "CREATE TABLE"
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

func lexFindTable(l *Lexer) stateFn {
	for {
		if strings.HasPrefix(l.input[l.pos:], sqlCreate) {
			if l.pos > l.start {
				l.Emit(itemCreate)
			}
			return lexTable
		}
		if l.next() == eof {
			break
		}
	}

	// Correctly reached EOF.
	if l.pos > l.start {
		l.Emit(itemText)
	}

	l.Emit(itemEOF) // Useful to make EOF a token.
	return nil      // Stop the run loop.
}

func lexLeftMeta(l *Lexer) stateFn {
	l.pos += len(leftMeta)
	l.emit(itemLeftMeta)
	return lexInsideAction // Now inside {{ }}.
}
