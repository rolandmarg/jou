package note

import (
	"time"

	"github.com/rolandmarg/jou/internal/pkg/random"
)

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

// Repository provideDAL
type Repository interface {
	Get(id int64) (*Note, error)
	GetByJournalID(id int64) ([]Note, error)
	Create(journalID int64, title, body, mood string, tags []string) (int64, error)
	Remove(id int64) error
}

// MockRepository is mock implementation of note MockRepository
type MockRepository struct {
	GetFn                 func(id int64) (*Note, error)
	GetInvoked            bool
	GetByJournalIDFn      func(id int64) ([]Note, error)
	GetByJournalIDInvoked bool
	CreateFn              func(journalID int64, title, body, mood string, tags []string) (int64, error)
	CreateInvoked         bool
	RemoveFn              func(id int64) error
	RemoveInvoked         bool
}

// Generate random note
func (r *MockRepository) Generate() *Note {
	n := &Note{}
	n.ID = random.Int64()
	n.JournalID = random.Int64()
	n.Title = random.String(128)
	n.Body = random.String(65535)
	n.Mood = random.String(64)
	n.Tags = random.Strings(64, 12)
	n.CreatedAt = random.Time()

	return n
}

// Get is mock implementation of note.MockRepository.Get
func (r *MockRepository) Get(id int64) (*Note, error) {
	r.GetInvoked = true
	return r.GetFn(id)
}

// GetByJournalID is mock implementation of journal.MockRepository.GetByJournalID
func (r *MockRepository) GetByJournalID(id int64) ([]Note, error) {
	r.GetByJournalIDInvoked = true
	return r.GetByJournalIDFn(id)
}

// Create is mock implementation of journal.MockRepository.Create
func (r *MockRepository) Create(journalID int64, title, body, mood string, tags []string) (int64, error) {
	r.CreateInvoked = true
	return r.CreateFn(journalID, title, body, mood, tags)
}

// Remove is mock implementation of note.MockRepository.Remove
func (r *MockRepository) Remove(id int64) error {
	r.RemoveInvoked = true
	return r.RemoveFn(id)
}
