package sqlite

import (
	"database/sql"
	"strings"
	"time"

	"github.com/rolandmarg/jou/internal/pkg/note"
)

type repository struct {
	db *sql.DB
}

// MakeRepository creates note repository
func MakeRepository(db *sql.DB) note.Repository {
	r := &repository{db}

	return r
}

func (r *repository) Get(id int64) (*note.Note, error) {
	row := r.db.QueryRow(`
		SELECT n.id, n.j_id, n.title, n.body, n.mood, n.created_at, GROUP_CONCAT(t.name)
		FROM note n
		LEFT JOIN tag t ON n.id = t.note_id
		WHERE n.id=?
			AND n.deleted_at IS NULL
		GROUP BY n.id`,
		id)

	n := &note.Note{}
	var body, mood, tags sql.NullString
	err := row.Scan(&n.ID, &n.JournalID, &n.Title, &body, &mood, &n.CreatedAt, &tags)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	if tags.Valid == true {
		n.Tags = strings.Split(tags.String, ",")
	}
	if body.Valid == true {
		n.Body = body.String
	}
	if mood.Valid == true {
		n.Mood = mood.String
	}

	return n, nil
}

func (r *repository) GetByJournalID(id int64) ([]note.Note, error) {
	rows, err := r.db.Query(`
		SELECT n.id, n.title, n.body, n.mood, n.created_at, GROUP_CONCAT(t.name)
		FROM note n
		LEFT JOIN tag t ON n.id = t.note_id
		WHERE n.j_id=?
			AND n.deleted_at IS NULL
		GROUP BY n.id`,
		id)
	if err != nil {
		return nil, err
	}

	n := note.Note{JournalID: id}
	var body, mood, tags sql.NullString
	var notes []note.Note
	for rows.Next() {
		err := rows.Scan(&n.ID, &n.Title, &body, &mood, &n.CreatedAt, &tags)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		if tags.Valid == true {
			n.Tags = strings.Split(tags.String, ",")
		}
		if body.Valid == true {
			n.Body = body.String
		}
		if mood.Valid == true {
			n.Mood = mood.String
		}
		notes = append(notes, n)
	}
	return notes, nil
}

func (r *repository) Create(journalID int64, title, body, mood string, tags []string) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	// TODO maybe handle rollback error
	defer tx.Rollback()

	res, err := tx.Exec(
		`INSERT INTO note (j_id, title, body, mood, created_at) VALUES (?, ?, ?, ?, ?)`,
		journalID, title, body, mood, time.Now())
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	for _, tag := range tags {
		// TODO maybe create tags in goroutine
		_, err = tx.Exec(`INSERT INTO tag (name, note_id) VALUES (?, ?)`, tag, id)
		if err != nil {
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) Remove(id int64) error {
	_, err := r.db.Exec(`UPDATE note SET deleted_at=? WHERE id=?`, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}
