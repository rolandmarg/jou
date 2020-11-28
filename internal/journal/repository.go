package journal

import (
	"database/sql"
	"time"

	"github.com/rolandmarg/jou/internal/journal/entry"
	"github.com/rolandmarg/jou/internal/pkg/kvstore"
)

type repository struct {
	DB *sql.DB
	es entry.Service
	ks kvstore.Service
}

// MakeRepository is a database bridge for journal
func MakeRepository(DB *sql.DB) Service {
	es := entry.MakeRepository(DB)
	ks := kvstore.MakeRepository(DB)
	r := &repository{DB, es, ks}

	return r
}

func (r *repository) Get(id int64) (*Journal, error) {
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

	j.entries, err = r.es.GetByJournalID(j.ID)
	if err != nil {
		return j, err
	}

	return j, nil
}

func (r *repository) GetAll() ([]Journal, error) {
	rows, err := r.DB.Query(`SELECT id, name, created_at FROM journal
	WHERE deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}

	j := Journal{}
	var journals []Journal

	for rows.Next() {
		err := rows.Scan(&j.ID, &j.name, &j.createdAt)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		// TODO fix n+1
		j.entries, err = r.es.GetByJournalID(j.ID)
		if err != nil {
			return journals, err
		}
		journals = append(journals, j)
	}

	return journals, nil
}

func (r *repository) GetByName(name string) (*Journal, error) {
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

	j.entries, err = r.es.GetByJournalID(j.ID)
	if err != nil {
		return j, err
	}

	return j, nil
}

func (r *repository) GetDefault() (*Journal, error) {
	name, err := r.ks.Get("default_journal")
	if err != nil {
		return nil, err
	}

	j, err := r.GetByName(name)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (r *repository) SetDefault(name string) error {
	err := r.ks.Set("default_journal", name)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Create(name string) (int64, error) {
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

func (r *repository) Update(id int64, name string) error {
	_, err := r.DB.Exec(`UPDATE journal SET name=? WHERE id=?`, name, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Remove(id int64) error {
	_, err := r.DB.Exec(`UPDATE journal SET deleted_at=? WHERE id=?`, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}
