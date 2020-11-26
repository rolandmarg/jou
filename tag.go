package main

import (
	"database/sql"
)

// Tag contains entry user input and other metadata
type Tag struct {
	id      int64
	entryID int64
	name    string
}

// TagRepo provides entry CRUD methods
type TagRepo struct {
	// TODO don't leak internal variables
	db *sql.DB
}

// MakeTagRepo creates a new entry repository
func MakeTagRepo(db *sql.DB) *TagRepo {
	r := &TagRepo{db: db}

	return r
}

// Create a new tag and return tag id
func (repo *TagRepo) Create(entryID int64, name string, tx *sql.Tx) (int64, error) {
	// TODO allow only alphabet characters in name
	res, err := tx.Exec(`INSERT INTO tag (name, entry_id) VALUES (?, ?)`, name, entryID)
	if err != nil {
		return 0, err
	}

	tagID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return tagID, err
}

// Get tags
func (repo *TagRepo) Get(entryID int64) ([]Tag, error) {
	rows, err := repo.db.Query(`SELECT id, name FROM tag WHERE entry_id=?`, entryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tag := Tag{entryID: entryID}
	var tags []Tag
	for rows.Next() {
		err = rows.Scan(&tag.id, &tag.name)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
