package kvstore

import "database/sql"

type repository struct {
	DB *sql.DB
}

// MakeRepository creates a key-value repository
func MakeRepository(DB *sql.DB) Service {
	r := &repository{DB}

	return r
}

func (r *repository) Get(key string) (string, error) {
	row := r.DB.QueryRow(`SELECT value FROM env WHERE name = ?`, key)

	var value string
	err := row.Scan(&value)
	if err != nil {
		return "", err
	}

	return value, nil
}

func (r *repository) Set(key string, value string) error {
	_, err := r.DB.Exec(`REPLACE INTO env (name, value) VALUES (?, ?)`, key, value)
	if err != nil {
		return err
	}

	return nil
}
