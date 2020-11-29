package note

import (
	"github.com/rolandmarg/jou/internal/pkg/note"
)

// Service provides note functions
type Service struct {
	n note.Repository
}

// MakeService creates note service
func MakeService(n note.Repository) *Service {
	s := &Service{n}

	return s
}

// Get note by id
func (s *Service) Get(id int64) (*note.Note, error) {
	return s.n.Get(id)
}

// GetByJournalID returns notes by journal id
func (s *Service) GetByJournalID(id int64) ([]note.Note, error) {
	return s.n.GetByJournalID(id)
}

// Create a note
func (s *Service) Create(journalID int64, title, body, mood string, tags []string) (int64, error) {
	return s.n.Create(journalID, title, body, mood, tags)
}

// Remove a note by id
func (s *Service) Remove(id int64) error {
	return s.n.Remove(id)
}
