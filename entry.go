package main

import (
	"database/sql"
	"time"
)

// Entry contains entry user input and other metadata
type Entry struct {
	id        int64
	deletedAt time.Time
	createdAt time.Time
	*EntryInput
}

// EntryInput contains all fields user can provide for journal entry creation
type EntryInput struct {
	title string
	body  string
	tags  []string
	mood  string
}

// EntryRepo provides entry CRUD methods
type EntryRepo struct {
	db *sql.DB // TODO don't leak db variable
}

// MakeEntryRepo creates a new entry repository
func MakeEntryRepo(db *sql.DB) *EntryRepo {
	e := &EntryRepo{db: db}

	return e
}

// Create a new journal entry
func (repo *EntryRepo) Create(journalID int64, i *EntryInput) (e *Entry, err error) {
	e = &Entry{createdAt: time.Now(), EntryInput: i}

	res, err := repo.db.Exec(
		`INSERT INTO entry (journalID, title, body, mood, created_at) values(?, ?, ?, ?, ?)`,
		journalID, e.title, e.body, e.mood, e.createdAt)
	if err != nil {
		return
	}

	e.id, err = res.LastInsertId()
	if err != nil {
		return
	}

	return
}

// Get a journal entry
func (repo *EntryRepo) Get(entryID int64) (e *Entry, err error) {
	e = &Entry{EntryInput: &EntryInput{}}

	row := repo.db.QueryRow(`
		SELECT id, title, body, mood, created_at, T.name FROM ENTRY E 
		WHERE E.id=? AND E.deleted_at IS NULL
		INNER JOIN TAG T
			ON E.id = T.entry_id
		`, entryID)
	err = row.Scan(&e.id, &e.title, &e.body, &e.mood, &e.createdAt)
	// TODO add goroutine to select tags too
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return
	}
	return
}

// Remove a journal entry
func (repo *EntryRepo) Remove(entryID int64) (affected bool, err error) {
	res, err := repo.db.Exec(`UPDATE entry SET deleted_at=? WHERE id=?`, time.Now(), entryID)
	if err != nil {
		return
	}

	numAffected, err := res.RowsAffected()
	if err != nil || numAffected == 0 {
		return
	}

	return true, nil
}
