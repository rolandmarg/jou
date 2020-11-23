package main

import (
	"database/sql"
	"time"
)

// Entry contains entry user input and other metadata
type Entry struct {
	id        int64
	journalID int64
	deletedAt time.Time
	createdAt time.Time
	title     string
	body      string
	tags      []Tag
	mood      string
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
	// TODO don't leak internal variables
	db      *sql.DB
	tagRepo *TagRepo
}

// MakeEntryRepo creates a new entry repository
func MakeEntryRepo(db *sql.DB) *EntryRepo {
	tagRepo := &TagRepo{db: db}
	e := &EntryRepo{db, tagRepo}

	return e
}

// Create a new journal entry and returns entry id
func (repo *EntryRepo) Create(journalID int64, i *EntryInput) (int64, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return 0, err
	}
	// TODO maybe handle rollback error
	defer tx.Rollback()

	res, err := tx.Exec(
		`INSERT INTO entry (journal_id, title, body, mood, created_at) values(?, ?, ?, ?, ?)`,
		journalID, i.title, i.body, i.mood, time.Now())
	if err != nil {
		return 0, err
	}

	entryID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	for _, tag := range i.tags {
		// TODO maybe create tags in goroutine
		_, err = repo.tagRepo.Create(entryID, tag, tx)
		if err != nil {
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return entryID, nil
}

// Get a journal entry
func (repo *EntryRepo) Get(entryID int64) (*Entry, error) {
	type result struct {
		tags []Tag
		err  error
	}
	c := make(chan result, 1)

	go func() {
		tags, err := repo.tagRepo.Get(entryID)
		c <- result{tags, err}
	}()

	row := repo.db.QueryRow(`SELECT id, journal_id, title, body, mood, created_at
		FROM entry WHERE id=? AND deleted_at IS NULL`, entryID)

	e := Entry{}
	err := row.Scan(&e.id, &e.journalID, &e.title, &e.body, &e.mood, &e.createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	r := <-c
	if r.err != nil {
		return nil, r.err
	}

	e.tags = r.tags

	return &e, nil
}

// Remove a journal entry, returns true or false if rows are affected
func (repo *EntryRepo) Remove(entryID int64) (bool, error) {
	res, err := repo.db.Exec(`UPDATE entry SET deleted_at=? WHERE id=?`, time.Now(), entryID)
	if err != nil {
		return false, err
	}

	numAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if numAffected == 0 {
		return false, nil
	}

	return true, nil
}
