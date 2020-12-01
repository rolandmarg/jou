package note

import (
	"errors"
	"fmt"

	"github.com/rolandmarg/jou/internal/pkg/journal"
	"github.com/rolandmarg/jou/internal/pkg/note"
)

// Service provides note functions
type Service struct {
	j journal.Repository
	n note.Repository
}

// MakeService creates note service
func MakeService(j journal.Repository, n note.Repository) *Service {
	s := &Service{j, n}

	return s
}

// Create a note
func (s *Service) Create(journalName string, title string) error {
	j, err := s.j.Get(journalName)
	if err != nil {
		return err
	}
	if j == nil {
		return errors.New("journal not found")
	}

	_, err = s.n.Create(j.ID, title, "", "", []string{})
	if err != nil {
		return fmt.Errorf("note not created: %w", err)
	}

	return nil
}

// CreateDefault creates note in default journal
func (s *Service) CreateDefault(title string) error {
	j, err := s.j.GetDefault()
	if err != nil {
		return err
	}
	if j == nil {
		return errors.New("default journal not found")
	}

	_, err = s.n.Create(j.ID, title, "", "", []string{})
	if err != nil {
		return fmt.Errorf("note not created: %w", err)
	}

	return nil
}
