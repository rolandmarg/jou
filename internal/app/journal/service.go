package journal

import (
	"errors"
	"fmt"

	"github.com/rolandmarg/jou/internal/pkg/journal"
	"github.com/rolandmarg/jou/internal/pkg/note"
)

// Service provides journal functions
type Service struct {
	j journal.Repository
	n note.Repository
	// TODO add custom logging for errors info and so on
}

// MakeService creates journal service
func MakeService(j journal.Repository, n note.Repository) *Service {
	s := &Service{j, n}

	return s
}

// Get journal by name
func (s *Service) Get(name string) (*journal.Journal, error) {
	j, err := s.j.Get(name)
	if err != nil {
		return nil, err
	}

	if j == nil {
		return nil, errors.New("journal not found")
	}

	j.Notes, err = s.n.GetByJournalID(j.ID)
	if err != nil {
		return j, err
	}

	return j, nil
}

// GetDefault journal
func (s *Service) GetDefault() (*journal.Journal, error) {
	j, err := s.j.GetDefault()
	if err != nil {
		return nil, err
	}

	if j == nil {
		return nil, errors.New("default journal not found")
	}

	j.Notes, err = s.n.GetByJournalID(j.ID)
	if err != nil {
		return j, err
	}

	return j, nil
}

// GetAll journals
func (s *Service) GetAll() ([]journal.Journal, error) {
	journals, err := s.j.GetAll()
	if err != nil {
		return nil, err
	}

	if journals == nil || len(journals) == 0 {
		return nil, errors.New("no journals found")
	}

	for i := range journals {
		journals[i].Notes, err = s.n.GetByJournalID(journals[i].ID)
		if err != nil {
			// TODO maybe try getting other journal notes instead of return
			return journals, err
		}
	}

	return journals, nil
}

// Create journal
func (s *Service) Create(name string, isDefault bool) error {
	j, err := s.j.Get(name)
	if err != nil {
		return err
	}

	if j != nil {
		return errors.New("journal already exists")
	}

	_, err = s.j.Create(name)
	if err != nil {
		return fmt.Errorf(`journal not created: %w`, err)
	}

	if isDefault {
		err = s.j.SetDefault(name)
		if err != nil {
			return fmt.Errorf(`journal created, but not set default: %w`, err)
		}
	}

	return nil
}

// Remove a journal
func (s *Service) Remove(name string) error {
	// TODO seems like we need goroutine and sqlite optimized for mthread read
	j, err := s.j.Get(name)
	if err != nil {
		return err
	}

	if j == nil {
		return errors.New("journal not found")
	}

	j, err = s.j.GetDefault()
	if err != nil {
		return err
	}

	if j.Name == name {
		return errors.New("deleting default journal is prohibited by law")
	}

	err = s.j.Remove(name)
	if err != nil {
		return err
	}

	return nil
}

// SetDefault journal
func (s *Service) SetDefault(name string) error {
	// TODO seems like we need goroutine and sqlite optimized for mthread read
	j, err := s.j.Get(name)
	if err != nil {
		return err
	}

	if j == nil {
		return errors.New("journal not found")
	}

	err = s.j.SetDefault(name)
	if err != nil {
		return err
	}

	return nil
}

// Rename a journal
func (s *Service) Rename(old string, new string) error {
	j, err := s.j.Get(old)
	if err != nil {
		return err
	}

	if j == nil {
		return errors.New("journal not found")
	}

	err = s.j.Update(old, new)
	if err != nil {
		return err
	}

	return nil
}
