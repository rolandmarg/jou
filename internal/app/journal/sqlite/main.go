package sqlite

import (
	"database/sql"
	"time"

	"github.com/rolandmarg/jou/internal/pkg/journal"
)

type repository struct {
	db *sql.DB
}

// MakeRepository creates journal repository
func MakeRepository(db *sql.DB) journal.Repository {
	r := &repository{db}

	return r
}

func (r *repository) Get(name string) (*journal.Journal, error) {
	row := r.db.QueryRow(`SELECT id, created_at FROM journal
		WHERE name = ? AND deleted_at IS NULL`, name)

	j := &journal.Journal{Name: name}
	err := row.Scan(&j.ID, &j.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return j, nil
}

func (r *repository) GetAll() ([]journal.Journal, error) {
	// TODO add pagination
	rows, err := r.db.Query(`SELECT id, name, created_at FROM journal
	WHERE deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}

	j := journal.Journal{}
	var journals []journal.Journal

	for rows.Next() {
		err := rows.Scan(&j.ID, &j.Name, &j.CreatedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		journals = append(journals, j)
	}

	return journals, nil
}

func (r *repository) GetDefault() (*journal.Journal, error) {
	row := r.db.QueryRow(`
		SELECT j.id, j.name, j.created_at
		FROM journal j
		JOIN default_journal d ON j.id = d.j_id
		WHERE j.deleted_at IS NULL
		`)

	j := &journal.Journal{}
	err := row.Scan(&j.ID, &j.Name, &j.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
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

func (r *repository) Update(oldName, newName string) error {
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
