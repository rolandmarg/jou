package entry

import "time"

// Entry structure
type Entry struct {
	ID        int64
	JournalID int64
	Title     string
	Body      string
	Mood      string
	Tags      []string
	CreatedAt time.Time
}

// Service provides operations on entry
type Service interface {
	Get(ID int64) (*Entry, error)
	GetByJournalID(ID int64) ([]Entry, error)
	Create(e *Entry) (int64, error)
	Remove(id int64) error
}
