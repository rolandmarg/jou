package journal

import (
	"errors"
	"testing"

	"github.com/rolandmarg/jou/internal/pkg/journal"
	"github.com/rolandmarg/jou/internal/pkg/journal/note"
)

func setup() (*Service, *journal.MockRepository, *note.MockRepository) {
	j := &journal.MockRepository{}
	n := &note.MockRepository{}
	s := MakeService(j, n)
	return s, j, n
}
func TestGet(t *testing.T) {
	t.Run("Should fail on repository null", func(t *testing.T) {
		s, jr, _ := setup()
		jr.GetFn = func(name string) (*journal.Journal, error) { return nil, nil }
		j, err := s.Get("random")
		if err == nil {
			t.Fatal("No error returned")
		}
		if j != nil {
			t.Fatal("Existing journal returned from space time continuum", j)
		}
		if !jr.GetInvoked {
			t.Fatal("Repository.Get did not invoke")
		}
	})
	t.Run("Should fail on repository error", func(t *testing.T) {
		s, jr, _ := setup()
		jr.GetFn = func(name string) (*journal.Journal, error) { return nil, errors.New("error") }
		_, err := s.Get("random")
		if err == nil {
			t.Fatal("No error returned")
		}
		if !jr.GetInvoked {
			t.Fatal("Repository.Get did not invoke")
		}
	})
	t.Run("Should success on repository success", func(t *testing.T) {
		s, jr, nr := setup()
		jr.GetFn = func(name string) (*journal.Journal, error) { return jr.Generate(), nil }
		nr.GetByJournalIDFn = func(id int64) ([]note.Note, error) { return []note.Note{*nr.Generate()}, nil }
		j, err := s.Get("random")
		if err != nil {
			t.Fatal(err)
		}
		if j == nil {
			t.Fatal("Expected journal to be set")
		}
		if !jr.GetInvoked {
			t.Fatal("Repository.Get did not invoke")
		}
		if !nr.GetByJournalIDInvoked {
			t.Fatal("Repository.GetByJournalID did not invoke")
		}
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Should fail on repository null", func(t *testing.T) {
		s, jr, _ := setup()
		jr.GetAllFn = func() ([]journal.Journal, error) { return nil, nil }
		j, err := s.GetAll()
		if err == nil {
			t.Fatal("No error returned")
		}
		if j != nil {
			t.Fatal("Existing journal returned from space time continuum", j)
		}
		if !jr.GetAllInvoked {
			t.Fatal("Repository.GetAll did not invoke")
		}
	})
	t.Run("Should fail on repository empty", func(t *testing.T) {
		s, jr, _ := setup()
		jr.GetAllFn = func() ([]journal.Journal, error) { return []journal.Journal{}, nil }
		j, err := s.GetAll()
		if err == nil {
			t.Fatal("No error returned")
		}
		if j != nil {
			t.Fatal("Existing journal returned from space time continuum", j)
		}
		if !jr.GetAllInvoked {
			t.Fatal("Repository.GetAll did not invoke")
		}
	})
	t.Run("Should fail on repository error", func(t *testing.T) {
		s, jr, _ := setup()
		jr.GetAllFn = func() ([]journal.Journal, error) { return []journal.Journal{}, errors.New("error") }
		_, err := s.GetAll()
		if err == nil {
			t.Fatal("No error returned")
		}
		if !jr.GetAllInvoked {
			t.Fatal("Repository.GetAll did not invoke")
		}
	})
	t.Run("Should success on repository success", func(t *testing.T) {
		s, jr, nr := setup()
		jar := []journal.Journal{*jr.Generate(), *jr.Generate()}
		nar := []note.Note{*nr.Generate(), *nr.Generate()}
		jr.GetAllFn = func() ([]journal.Journal, error) { return jar, nil }
		nr.GetByJournalIDFn = func(id int64) ([]note.Note, error) { return nar, nil }
		j, err := s.GetAll()
		if err != nil {
			t.Fatal(err)
		}
		if len(jar) != len(j) {
			t.Fatalf("Expected length %v received %v", len(jar), len(j))
		}
		for i := range j {
			if len(j[i].Notes) != len(nar) {
				t.Fatalf("Expected notes length %v received %v", len(nar), len(j[i].Notes))
			}
		}
		if !jr.GetAllInvoked {
			t.Fatal("Repository.GetAll did not invoke")
		}
		if !nr.GetByJournalIDInvoked {
			t.Fatal("Repository.GetByJournalID did not invoke")
		}
	})
}

func TestCreate(t *testing.T) {
	t.Run("Should fail on repository get success", func(t *testing.T) {
		s, jr, _ := setup()
		rnd := jr.Generate()
		jr.GetFn = func(name string) (*journal.Journal, error) { return rnd, nil }
		jr.GetInvoked = false
		err := s.Create(rnd.Name, false)
		if err == nil {
			t.Fatal("No error returned")
		}
		if !jr.GetInvoked {
			t.Fatal("Repository.GetAll did not invoke")
		}
	})
	t.Run("Should fail on repository error", func(t *testing.T) {
		s, jr, _ := setup()
		rnd := jr.Generate()
		jr.GetFn = func(name string) (*journal.Journal, error) { return rnd, errors.New("error") }
		jr.GetInvoked = false
		err := s.Create(rnd.Name, false)
		if err == nil {
			t.Fatal("No error returned")
		}
		if !jr.GetInvoked {
			t.Fatal("Repository.GetAll did not invoke")
		}
	})
	t.Run("Should fail on repository setDefault error", func(t *testing.T) {
		s, jr, _ := setup()
		jr.GetFn = func(name string) (*journal.Journal, error) { return nil, nil }
		jr.CreateFn = func(name string) (int64, error) { return 1, nil }
		jr.SetDefaultFn = func(name string) error { return errors.New("error") }
		err := s.Create("jo", true)
		if err == nil {
			t.Fatal("No error returned")
		}
		if !jr.GetInvoked {
			t.Fatal("Repository.GetAll did not invoke")
		}
		if !jr.SetDefaultInvoked {
			t.Fatal("Repository.SetDefault did not invoke")
		}
		if !jr.CreateInvoked {
			t.Fatal("Repository.Create did not invoke")
		}
	})
	t.Run("Should not invoke repository setDefault", func(t *testing.T) {
		s, jr, _ := setup()
		jr.GetFn = func(name string) (*journal.Journal, error) { return nil, nil }
		jr.CreateFn = func(name string) (int64, error) { return 1, nil }
		jr.SetDefaultFn = func(name string) error { return errors.New("error") }
		err := s.Create("jo", false)
		if err != nil {
			t.Fatal(err)
		}
		if jr.SetDefaultInvoked {
			t.Fatal("Repository.SetDefault invoked")
		}
	})
	t.Run("Should create new journal", func(t *testing.T) {
		s, jr, _ := setup()
		jr.GetFn = func(name string) (*journal.Journal, error) { return nil, nil }
		jr.CreateFn = func(name string) (int64, error) { return 1, nil }
		jr.SetDefaultFn = func(name string) error { return nil }
		err := s.Create("jo", true)
		if err != nil {
			t.Fatal(err)
		}
		if !jr.GetInvoked {
			t.Fatal("Repository.GetAll did not invoke")
		}
		if !jr.SetDefaultInvoked {
			t.Fatal("Repository.SetDefault did not invoke")
		}
		if !jr.CreateInvoked {
			t.Fatal("Repository.Create did not invoke")
		}
	})
}