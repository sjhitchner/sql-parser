package main

import (
	"fmt"
	"github.com/sjhitchner/lexer"
)

func main() {

	l := lexer.NewLexer(schema, lexCreate)

	for {
		token := l.NextToken()
		if token.Type == lexer.TokenEOF || token.Type == lexer.TokenError {
			break
		}
		fmt.Println(token)
	}
}

const schema = `
CREATE TABLE IF NOT EXISTS user (
	id INTEGER PRIMARY KEY AUTOINCREMENT
	, username TEXT NOT NULL
	, email TEXT NOT NULL
	, password TEXT NOT NULL
	, UNIQUE(username)
	, UNIQUE(email)
);
CREATE TABLE IF NOT EXISTS team (
	id INTEGER PRIMARY KEY AUTOINCREMENT
	, owner_id INTEGER NOT NULL
	, name TEXT NOT NULL 
	, FOREIGN KEY (owner_id) REFERENCES user(id)
	    ON UPDATE CASCADE ON DELETE CASCADE
   , UNIQUE(name)
);
CREATE TABLE IF NOT EXISTS channel (
	id INTEGER PRIMARY KEY AUTOINCREMENT
	, team_id INTEGER NOT NULL
	, owner_id INTEGER NOT NULL
	, name TEXT NOT NULL
	, is_public BOOLEAN NOT NULL
	, FOREIGN KEY (team_id) REFERENCES team(id)
	    ON UPDATE CASCADE ON DELETE CASCADE
	, FOREIGN KEY (owner_id) REFERENCES user(id)
	    ON UPDATE CASCADE ON DELETE CASCADE
   , UNIQUE(team_id, name)
);
CREATE TABLE IF NOT EXISTS team_member (
	id INTEGER PRIMARY KEY AUTOINCREMENT
	, team_id INTEGER NOT NULL
	, user_id INTEGER NOT NULL
	, FOREIGN KEY (team_id) REFERENCES team(id)
	    ON UPDATE CASCADE ON DELETE CASCADE
	, FOREIGN KEY (user_id) REFERENCES user(id)
	    ON UPDATE CASCADE ON DELETE CASCADE
   , UNIQUE(team_id, user_id)
);
CREATE TABLE IF NOT EXISTS channel_member (
	id INTEGER PRIMARY KEY AUTOINCREMENT
	, channel_id INTEGER NOT NULL
	, user_id INTEGER NOT NULL
	, FOREIGN KEY (channel_id) REFERENCES channel(id)
	    ON UPDATE CASCADE ON DELETE CASCADE
	, FOREIGN KEY (user_id) REFERENCES user(id)
	    ON UPDATE CASCADE ON DELETE CASCADE
   , UNIQUE(channel_id, user_id)
);
CREATE TABLE IF NOT EXISTS message (
	id INTEGER PRIMARY KEY AUTOINCREMENT
	, user_id INTEGER NOT NULL
	, channel_id INTEGER NOT NULL
	, text TEXT NOT NULL
	, timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	, FOREIGN KEY (user_id) REFERENCES user(id)
	    ON UPDATE CASCADE ON DELETE CASCADE
	, FOREIGN KEY (channel_id) REFERENCES channel(id)
	    ON UPDATE CASCADE ON DELETE CASCADE
);
`
