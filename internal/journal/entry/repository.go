package entry

import (
	"database/sql"
	"strings"
	"time"
)

type repository struct {
	db *sql.DB
}

// MakeRepository is a database bridge for entry
func MakeRepository(db *sql.DB) Service {
	r := &repository{db}

	return r
}

func (r *repository) Get(id int64) (*Entry, error) {
	row := r.db.QueryRow(`
		SELECT e.ID, e.journal_id, e.title, e.body, e.mood, e.created_at, GROUP_CONCAT(t.name)
		FROM entry e
		LEFT JOIN tag t ON e.ID = t.entry_id
		WHERE e.ID=?
			AND e.deleted_at IS NULL
		GROUP BY e.ID`,
		id)

	e := &Entry{}
	var body, mood, tags sql.NullString
	err := row.Scan(&e.ID, &e.JournalID, &e.Title, &body, &mood, &e.CreatedAt, &tags)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	if tags.Valid == true {
		e.Tags = strings.Split(tags.String, ",")
	}
	if body.Valid == true {
		e.Body = body.String
	}
	if mood.Valid == true {
		e.Mood = mood.String
	}

	return e, nil
}

func (r *repository) GetByJournalID(id int64) ([]Entry, error) {
	rows, err := r.db.Query(`
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

	e := Entry{JournalID: id}
	var body, mood, tags sql.NullString
	var entries []Entry
	for rows.Next() {
		err := rows.Scan(&e.ID, &e.Title, &body, &mood, &e.CreatedAt, &tags)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		if tags.Valid == true {
			e.Tags = strings.Split(tags.String, ",")
		}
		if body.Valid == true {
			e.Body = body.String
		}
		if mood.Valid == true {
			e.Mood = mood.String
		}
		entries = append(entries, e)
	}
	return entries, nil
}

func (r *repository) Create(e *Entry) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	// TODO maybe handle rollback error
	defer tx.Rollback()

	e.CreatedAt = time.Now()

	res, err := tx.Exec(
		`INSERT INTO entry (journal_id, title, body, mood, created_at) VALUES (?, ?, ?, ?, ?)`,
		e.JournalID, e.Title, e.Body, e.Mood, e.CreatedAt)
	if err != nil {
		return 0, err
	}

	e.ID, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	for _, tag := range e.Tags {
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

func (r *repository) Remove(id int64) error {
	_, err := r.db.Exec(`UPDATE entry SET deleted_at=? WHERE id=?`, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}
