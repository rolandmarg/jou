package main

import (
	"fmt"
	"time"
)

// Journal structure
type Journal struct {
	ID        int64
	name      string
	entries   []Entry
	createdAt time.Time
}

// JournalService provides operations on journal
type JournalService interface {
	Get(ID int64) (*Journal, error)
	GetByName(name string) (*Journal, error)
	GetDefault() (*Journal, error)
	GetAll() ([]Journal, error)
	SetDefault(name string) error
	Create(name string) (int64, error)
	Update(ID int64, name string) error
	Remove(ID int64) error
}

func (journal *Journal) String() string {
	str := fmt.Sprintln("journal:")

	for _, e := range journal.entries {
		str = fmt.Sprintln(str, " entry:")
		str = fmt.Sprintln(str, "   id:", e.ID)
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

// Entry structure
type Entry struct {
	ID        int64
	journalID int64
	title     string
	body      string
	mood      string
	tags      []string
	createdAt time.Time
}

// EntryService provides operations on journal
type EntryService interface {
	Get(ID int64) (*Entry, error)
	GetByJournalID(ID int64) ([]Entry, error)
	Create(e *Entry) (int64, error)
	Remove(id int64) error
}

// KVService provides operations on key-value store
type KVService interface {
	Get(name string) (string, error)
	Set(name string, value string) error
}
