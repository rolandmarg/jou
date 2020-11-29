package journal

import (
	"database/sql"
	"time"

	"github.com/rolandmarg/jou/internal/journal/note"
)

type repository struct {
	db *sql.DB
	es note.Service
}

// MakeRepository is a database bridge for journal
func MakeRepository(db *sql.DB) Service {
	es := note.MakeRepository(db)
	r := &repository{db, es}

	return r
}

func (r *repository) GetAll() ([]Journal, error) {
	// TODO add pagination
	rows, err := r.db.Query(`SELECT id, name, created_at FROM journal
	WHERE deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}

	j := Journal{}
	var journals []Journal

	for rows.Next() {
		err := rows.Scan(&j.ID, &j.Name, &j.CreatedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		// TODO fix n+1
		j.Notes, err = r.es.GetByJournalID(j.ID)
		if err != nil {
			return journals, err
		}
		journals = append(journals, j)
	}

	return journals, nil
}

func (r *repository) Get(name string) (*Journal, error) {
	// TODO possibly get journal and notes in 1 sql statement
	// or use goroutines
	row := r.db.QueryRow(`SELECT id, created_at FROM journal
		WHERE name = ? AND deleted_at IS NULL`, name)

	j := &Journal{Name: name}
	err := row.Scan(&j.ID, &j.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	j.Notes, err = r.es.GetByJournalID(j.ID)
	if err != nil {
		return j, err
	}

	return j, nil
}

func (r *repository) GetDefault() (*Journal, error) {
	row := r.db.QueryRow(`
		SELECT j.id, j.name, j.created_at
		FROM journal j
		JOIN default_journal d ON j.id = d.j_id
		WHERE j.deleted_at IS NULL
		`)

	j := &Journal{}
	err := row.Scan(&j.ID, &j.Name, &j.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	j.Notes, err = r.es.GetByJournalID(j.ID)
	if err != nil {
		return j, err
	}

	return j, nil
}

func (r *repository) SetDefault(name string) error {
	_, err := r.db.Exec(`
		REPLACE INTO default_journal (name, j_id)
		SELECT ?, id FROM journal
		WHERE name = ?`,
		"default", name)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Create(name string) (int64, error) {
	res, err := r.db.Exec(`INSERT INTO journal (name, created_at) VALUES (?, ?)`,
		name, time.Now())
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, err
}

func (r *repository) Update(oldName string, newName string) error {
	// TODO return not found error if deleted_at is set or name not found
	_, err := r.db.Exec(`UPDATE journal SET name=? WHERE name=?`, newName, oldName)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Remove(name string) error {
	_, err := r.db.Exec(`UPDATE journal SET deleted_at=? WHERE name=?`, time.Now(), name)
	if err != nil {
		return err
	}
	return nil
}
