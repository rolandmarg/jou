package kvstore

import "database/sql"

type repository struct {
	db *sql.DB
}

// MakeRepository creates a key-value repository
func MakeRepository(db *sql.DB) Service {
	r := &repository{db}

	return r
}

func (r *repository) Get(key string) (string, error) {
	row := r.db.QueryRow(`SELECT value FROM env WHERE key = ?`, key)

	var value string
	err := row.Scan(&value)
	if err != nil {
		return "", err
	}

	return value, nil
}

func (r *repository) Set(key string, value string) error {
	_, err := r.db.Exec(`REPLACE INTO env (key, value) VALUES (?, ?)`, key, value)
	if err != nil {
		return err
	}

	return nil
}
