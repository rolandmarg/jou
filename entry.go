package main

import (
	"database/sql"
	"strings"
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
	tags      []string
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
	tagRepo := MakeTagRepo(db)
	entryRepo := &EntryRepo{db, tagRepo}

	return entryRepo
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
		`INSERT INTO entry (journal_id, title, body, mood, created_at) VALUES (?, ?, ?, ?, ?)`,
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

// Get an entry
func (repo *EntryRepo) Get(id int64) (*Entry, error) {
	row := repo.db.QueryRow(`
		SELECT e.id, e.journal_id, e.title, e.body, e.mood, e.created_at, GROUP_CONCAT(t.name)
		FROM entry e
		LEFT JOIN tag t ON e.id = t.entry_id
		WHERE e.id=?
			AND e.deleted_at IS NULL
		GROUP BY e.id`,
		id)

	e := &Entry{}
	var body, mood, tags sql.NullString
	err := row.Scan(&e.id, &e.journalID, &e.title, &body, &mood, &e.createdAt, &tags)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	if tags.Valid == true {
		e.tags = strings.Split(tags.String, ",")
	}
	if body.Valid == true {
		e.body = body.String
	}
	if mood.Valid == true {
		e.mood = mood.String
	}

	return e, nil
}

// GetByJournalID returns entries by journal id
func (repo *EntryRepo) GetByJournalID(id int64) ([]Entry, error) {
	rows, err := repo.db.Query(`
		SELECT e.id, e.title, e.body, e.mood, e.created_at, GROUP_CONCAT(t.name)
		FROM entry e
		LEFT JOIN tag t ON e.id = t.entry_id
		WHERE e.journal_id=?
			AND e.deleted_at IS NULL
		GROUP BY e.id`,
		id)
	if err != nil {
		return nil, err
	}

	e := Entry{journalID: id}
	var body, mood, tags sql.NullString
	var entries []Entry
	for rows.Next() {
		err := rows.Scan(&e.id, &e.title, &body, &mood, &e.createdAt, &tags)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		if tags.Valid == true {
			e.tags = strings.Split(tags.String, ",")
		}
		if body.Valid == true {
			e.body = body.String
		}
		if mood.Valid == true {
			e.mood = mood.String
		}
		entries = append(entries, e)
	}
	return entries, nil
}

// Remove journal entry
func (repo *EntryRepo) Remove(id int64) error {
	_, err := repo.db.Exec(`UPDATE entry SET deleted_at=? WHERE id=?`, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}
