package note

import "time"

// Note structure
type Note struct {
	ID        int64
	JournalID int64
	Title     string
	Body      string
	Mood      string
	Tags      []string
	CreatedAt time.Time
}

// Service provides operations on note
type Service interface {
	Get(ID int64) (*Note, error)
	GetByJournalID(ID int64) ([]Note, error)
	Create(e *Note) (int64, error)
	Remove(id int64) error
}
