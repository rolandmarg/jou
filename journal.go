package main

import (
	"database/sql"
	"fmt"
	"time"
)

// Journal contains journal entries
type Journal struct {
	id        int64
	deletedAt time.Time
	createdAt time.Time
	name      string
	entries   []Entry
}

// JournalRepo provides journal CRUD operations
type JournalRepo struct {
	db        *sql.DB
	entryRepo *EntryRepo
	envRepo   *EnvRepo
}

// MakeJournalRepo creates a new journal repository
func MakeJournalRepo(db *sql.DB) *JournalRepo {
	entryRepo := MakeEntryRepo(db)
	envRepo := MakeEnvRepo(db)
	journalRepo := &JournalRepo{db, entryRepo, envRepo}

	return journalRepo
}

// Create a new journal
func (repo *JournalRepo) Create(name string) (int64, error) {
	res, err := repo.db.Exec(`INSERT INTO journal (name, created_at) VALUES (?, ?)`,
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

// Get journal
func (repo *JournalRepo) Get(id int64) (*Journal, error) {
	// TODO possibly get journal and entries in 1 sql statement
	// or use goroutines
	row := repo.db.QueryRow(`SELECT id, name, created_at FROM journal
		WHERE id=? AND deleted_at IS NULL`, id)

	j := &Journal{}
	err := row.Scan(&j.id, &j.name, &j.createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	j.entries, err = repo.entryRepo.GetByJournalID(j.id)
	if err != nil {
		return j, err
	}

	return j, nil
}

// GetByName journal
func (repo *JournalRepo) GetByName(name string) (*Journal, error) {
	// TODO possibly get journal and entries in 1 sql statement
	// or use goroutines
	row := repo.db.QueryRow(`SELECT id, created_at FROM journal
		WHERE name = ? AND deleted_at IS NULL`, name)

	j := &Journal{name: name}
	err := row.Scan(&j.id, &j.createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	j.entries, err = repo.entryRepo.GetByJournalID(j.id)
	if err != nil {
		return j, err
	}

	return j, nil
}

// GetDefault journal
func (repo *JournalRepo) GetDefault() (*Journal, error) {
	name, err := repo.envRepo.Get("default_journal")
	if err != nil {
		return nil, err
	}

	j, err := repo.GetByName(name)
	if err != nil {
		return nil, err
	}

	return j, nil
}

// SetDefault journal by name
func (repo *JournalRepo) SetDefault(name string) error {
	err := repo.envRepo.Set("default_journal", name)
	if err != nil {
		return err
	}

	return nil
}

// Update journal name
func (repo *JournalRepo) Update(id int64, name string) error {
	_, err := repo.db.Exec(`UPDATE journal SET name=? WHERE id=?`, name, id)
	if err != nil {
		return err
	}

	return nil
}

// Remove journal
func (repo *JournalRepo) Remove(id int64) error {
	_, err := repo.db.Exec(`UPDATE journal SET deleted_at=? WHERE id=?`, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (journal *Journal) String() string {
	str := fmt.Sprintln("journal:")

	for _, e := range journal.entries {
		str = fmt.Sprintln(str, " entry:")
		str = fmt.Sprintln(str, "   id:", e.id)
		str = fmt.Sprintln(str, "   title:", e.title)
		str = fmt.Sprintln(str, "   body:", e.body)
		if e.mood != "" {
			str = fmt.Sprintln(str, "   mood:", e.mood)
		}
		if e.tags != nil {
			str = fmt.Sprintln(str, "   tags:", e.tags)
		}
		str = fmt.Sprintln(str, "   createdAt:", e.createdAt.Format("2006-01-02 15:04:05"))
	}

	return str
}
