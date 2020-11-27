package main

import (
	"database/sql"
	"strings"
	"time"
)

// Schema is an sql tables string
const Schema = `
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
		name VARCHAR(128) NOT NULL,
		value TEXT NOT NULL,
		UNIQUE(name)
	);
`

// JournalRepository is a database bridge for journal
type JournalRepository struct {
	DB           *sql.DB
	entryService EntryService
	kService     KVService
}

// CreateJournalRepository creates a new journal repository
func CreateJournalRepository(DB *sql.DB, entryService EntryService, kService KVService) *JournalRepository {
	journalRepository := &JournalRepository{DB, entryService, kService}

	return journalRepository
}

// Get journal
func (r *JournalRepository) Get(id int64) (*Journal, error) {
	// TODO possibly get journal and entries in 1 sql statement
	// or use goroutines. what if we get all journals, n+1 problem?
	row := r.DB.QueryRow(`SELECT id, name, created_at FROM journal
		WHERE id=? AND deleted_at IS NULL`, id)

	j := &Journal{}
	err := row.Scan(&j.ID, &j.name, &j.createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	j.entries, err = r.entryService.GetByJournalID(j.ID)
	if err != nil {
		return j, err
	}

	return j, nil
}

// GetByName journal
func (r *JournalRepository) GetByName(name string) (*Journal, error) {
	// TODO possibly get journal and entries in 1 sql statement
	// or use goroutines
	row := r.DB.QueryRow(`SELECT id, created_at FROM journal
		WHERE name = ? AND deleted_at IS NULL`, name)

	j := &Journal{name: name}
	err := row.Scan(&j.ID, &j.createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	j.entries, err = r.entryService.GetByJournalID(j.ID)
	if err != nil {
		return j, err
	}

	return j, nil
}

// GetDefault journal
func (r *JournalRepository) GetDefault() (*Journal, error) {
	name, err := r.kService.Get("default_journal")
	if err != nil {
		return nil, err
	}

	j, err := r.GetByName(name)
	if err != nil {
		return nil, err
	}

	return j, nil
}

// SetDefault journal by name
func (r *JournalRepository) SetDefault(name string) error {
	err := r.kService.Set("default_journal", name)
	if err != nil {
		return err
	}

	return nil
}

// Create a new journal
func (r *JournalRepository) Create(name string) (int64, error) {
	res, err := r.DB.Exec(`INSERT INTO journal (name, created_at) VALUES (?, ?)`,
		name, time.Now())
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Update journal name
func (r *JournalRepository) Update(id int64, name string) error {
	_, err := r.DB.Exec(`UPDATE journal SET name=? WHERE id=?`, name, id)
	if err != nil {
		return err
	}

	return nil
}

// Remove journal
func (r *JournalRepository) Remove(id int64) error {
	_, err := r.DB.Exec(`UPDATE journal SET deleted_at=? WHERE id=?`, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

// EntryRepository is a database bridge for entry
type EntryRepository struct {
	DB *sql.DB
}

// CreateEntryRepository creates a new entry repository
func CreateEntryRepository(DB *sql.DB) *EntryRepository {
	entryRepository := &EntryRepository{DB}

	return entryRepository
}

// Get an entry
func (r *EntryRepository) Get(id int64) (*Entry, error) {
	row := r.DB.QueryRow(`
		SELECT e.ID, e.journal_id, e.title, e.body, e.mood, e.created_at, GROUP_CONCAT(t.name)
		FROM entry e
		LEFT JOIN tag t ON e.ID = t.entry_id
		WHERE e.ID=?
			AND e.deleted_at IS NULL
		GROUP BY e.ID`,
		id)

	e := &Entry{}
	var body, mood, tags sql.NullString
	err := row.Scan(&e.ID, &e.journalID, &e.title, &body, &mood, &e.createdAt, &tags)
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
func (r *EntryRepository) GetByJournalID(id int64) ([]Entry, error) {
	rows, err := r.DB.Query(`
		SELECT e.ID, e.title, e.body, e.mood, e.created_at, GROUP_CONCAT(t.name)
		FROM entry e
		LEFT JOIN tag t ON e.ID = t.entry_id
		WHERE e.journal_id=?
			AND e.deleted_at IS NULL
		GROUP BY e.ID`,
		id)
	if err != nil {
		return nil, err
	}

	e := Entry{journalID: id}
	var body, mood, tags sql.NullString
	var entries []Entry
	for rows.Next() {
		err := rows.Scan(&e.ID, &e.title, &body, &mood, &e.createdAt, &tags)
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

// Create a new journal entry and returns entry id
func (r *EntryRepository) Create(e *Entry) (int64, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return 0, err
	}
	// TODO maybe handle rollback error
	defer tx.Rollback()

	e.createdAt = time.Now()

	res, err := tx.Exec(
		`INSERT INTO entry (journal_id, title, body, mood, created_at) VALUES (?, ?, ?, ?, ?)`,
		e.journalID, e.title, e.body, e.mood, e.createdAt)
	if err != nil {
		return 0, err
	}

	e.ID, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	for _, tag := range e.tags {
		// TODO maybe create tags in goroutine
		_, err = tx.Exec(`INSERT INTO tag (name, entry_id) VALUES (?, ?)`, tag, e.ID)
		if err != nil {
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return e.ID, nil
}

// Remove journal entry
func (r *EntryRepository) Remove(id int64) error {
	_, err := r.DB.Exec(`UPDATE entry SET deleted_at=? WHERE id=?`, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

// KVRepository is a database bridge for key-value store
type KVRepository struct {
	DB *sql.DB
}

// CreateKVRepository creates a key-value repository
func CreateKVRepository(DB *sql.DB) *KVRepository {
	r := &KVRepository{DB}

	return r
}

// Get key value
func (r *KVRepository) Get(name string) (string, error) {
	row := r.DB.QueryRow(`SELECT value FROM env WHERE name = ?`, name)

	var value string
	err := row.Scan(&value)
	if err != nil {
		return "", err
	}

	return value, nil
}

// Set key value
func (r *KVRepository) Set(name string, value string) error {
	_, err := r.DB.Exec(`REPLACE INTO env (name, value) VALUES (?, ?)`, name, value)
	if err != nil {
		return err
	}

	return nil
}
