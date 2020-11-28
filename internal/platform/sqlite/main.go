package sqlite

import (
	"database/sql"
	"time"

	// import sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

const schema = `
	CREATE TABLE IF NOT EXISTS journal (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(64) NOT NULL,
		created_at DATE NOT NULL,
		deleted_at DATE,
		UNIQUE(name)
	);

	CREATE TABLE IF NOT EXISTS entry (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(128) NOT NULL,
		body TEXT,
		mood VARCHAR(64),
		created_at DATE NOT NULL,
		deleted_at DATE,
		journal_id INTEGER NOT NULL REFERENCES journal(id) ON DELETE CASCADE
	);
	CREATE INDEX IF NOT EXISTS entry_journal_idx on entry (journal_id);

	CREATE TABLE IF NOT EXISTS tag (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(64) NOT NULL,
		entry_id INTEGER NOT NULL REFERENCES entry(id) ON DELETE CASCADE,
		UNIQUE(entry_id, name)
	);

	CREATE TABLE IF NOT EXISTS env (
		key VARCHAR(128) NOT NULL,
		value TEXT NOT NULL,
		UNIQUE(key)
	);
`

// Open a database connection and executes sql schema on it
func Open(name string) (*sql.DB, error) {
	DB, err := sql.Open("sqlite3", name)
	if err != nil {
		return nil, err
	}

	_, err = DB.Exec(schema)
	if err != nil {
		return nil, err
	}

	return DB, nil
}

// OpenDB is a helper function that opens and populates production database
func OpenDB() (*sql.DB, error) {
	db, err := Open("jou.db")
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return db, err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		INSERT INTO journal (name, created_at)
		SELECT ?, ?
		WHERE NOT EXISTS (SELECT * FROM journal)
	`, "default", time.Now())
	if err != nil {
		return db, err
	}

	_, err = tx.Exec(`
		INSERT INTO env (key, value)
		SELECT ?, ?
		WHERE NOT EXISTS (SELECT * FROM env WHERE key=?)
	`, "default_journal", "default", "default_journal")
	if err != nil {
		return db, err
	}

	tx.Commit()

	return db, nil
}
