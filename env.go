package main

import "database/sql"

// Env contains environmental information
type Env struct {
	name  string
	value string
}

// EnvRepo for get/set access to KV store
type EnvRepo struct {
	db *sql.DB
}

// MakeEnvRepo creates getter/setter struct
func MakeEnvRepo(db *sql.DB) *EnvRepo {
	r := &EnvRepo{db}

	return r
}

// Get key value
func (repo *EnvRepo) Get(name string) (string, error) {
	row := repo.db.QueryRow(`SELECT value FROM env WHERE name = ?`, name)

	var value string
	err := row.Scan(&value)
	if err != nil {
		return "", err
	}

	return value, nil
}

// Set key value
func (repo *EnvRepo) Set(name string, value string) error {
	_, err := repo.db.Exec(`REPLACE INTO env (name, value) VALUES (?, ?)`, name, value)
	if err != nil {
		return err
	}

	return nil
}
