package note

import (
	"database/sql"
	"strings"
	"time"
)

type repository struct {
	db *sql.DB
}

// MakeRepository is a database bridge for note
func MakeRepository(db *sql.DB) Service {
	r := &repository{db}

	return r
}

func (r *repository) Get(id int64) (*Note, error) {
	row := r.db.QueryRow(`
		SELECT n.ID, n.journal_id, n.title, n.body, n.mood, n.created_at, GROUP_CONCAT(t.name)
		FROM note n
		LEFT JOIN tag t ON n.ID = t.note_id
		WHERE n.ID=?
			AND n.deleted_at IS NULL
		GROUP BY n.ID`,
		id)

	n := &Note{}
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

func (r *repository) GetByJournalID(id int64) ([]Note, error) {
	rows, err := r.db.Query(`
		SELECT n.ID, n.title, n.body, n.mood, n.created_at, GROUP_CONCAT(t.name)
		FROM note n
		LEFT JOIN tag t ON n.ID = t.note_id
		WHERE n.journal_id=?
			AND n.deleted_at IS NULL
		GROUP BY n.ID`,
		id)
	if err != nil {
		return nil, err
	}

	n := Note{JournalID: id}
	var body, mood, tags sql.NullString
	var notes []Note
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

func (r *repository) Create(n *Note) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	// TODO maybe handle rollback error
	defer tx.Rollback()

	n.CreatedAt = time.Now()

	res, err := tx.Exec(
		`INSERT INTO note (journal_id, title, body, mood, created_at) VALUES (?, ?, ?, ?, ?)`,
		n.JournalID, n.Title, n.Body, n.Mood, n.CreatedAt)
	if err != nil {
		return 0, err
	}

	n.ID, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	for _, tag := range n.Tags {
		// TODO maybe create tags in goroutine
		_, err = tx.Exec(`INSERT INTO tag (name, note_id) VALUES (?, ?)`, tag, n.ID)
		if err != nil {
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return n.ID, nil
}

func (r *repository) Remove(id int64) error {
	_, err := r.db.Exec(`UPDATE note SET deleted_at=? WHERE id=?`, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}
