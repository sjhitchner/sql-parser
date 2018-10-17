package main

import (
//"fmt"
)

const (
	sqlCreate        = "CREATE TABLE"
	sqlIfNotExists   = "IF NOT EXISTS"
	sqlLeftPara      = "("
	sqlRightPara     = ")"
	sqlComma         = ","
	sqlSemiColon     = ";"
	sqlInteger       = "INTEGER"
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

func lexFindTable(l *lexer) stateFn {
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
