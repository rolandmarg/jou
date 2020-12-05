package journal

import (
	"fmt"
	"time"

	"github.com/rolandmarg/jou/internal/pkg/journal/note"
	"github.com/rolandmarg/jou/internal/pkg/random"
)

// Journal structure
type Journal struct {
	ID        int64
	Name      string
	Notes     []note.Note
	CreatedAt time.Time
}

func (j Journal) String() string {
	str := fmt.Sprintf(`journal "%v" notes: [`, j.Name)
	for _, e := range j.Notes {
		str = fmt.Sprintln(str, " note:")
		str = fmt.Sprintln(str, "   id:", e.ID)
		str = fmt.Sprintln(str, "   title:", e.Title)
		if e.Body != "" {
			str = fmt.Sprintln(str, "   body:", e.Body)
		}
		if e.Mood != "" {
			str = fmt.Sprintln(str, "   mood:", e.Mood)
		}
		if e.Tags != nil {
			str = fmt.Sprintln(str, "   tags:", e.Tags)
		}
		str = fmt.Sprintln(str, "   createdAt:", e.CreatedAt.Format("2006-01-02 15:04:05"))
	}
	str = fmt.Sprint(str, "]")

	return str
}

// Repository provides operations on journal DAL
type Repository interface {
	Get(name string) (*Journal, error)
	GetDefault() (*Journal, error)
	GetAll() ([]Journal, error)
	SetDefault(name string) error
	Create(name string) (int64, error)
	Update(oldName, newName string) error
	Remove(name string) error
}

// MockRepository is mock implementation of journal MockRepository
type MockRepository struct {
	GetFn             func(name string) (*Journal, error)
	GetInvoked        bool
	GetAllFn          func() ([]Journal, error)
	GetAllInvoked     bool
	GetDefaultFn      func() (*Journal, error)
	GetDefaultInvoked bool
	SetDefaultFn      func(name string) error
	SetDefaultInvoked bool
	CreateFn          func(name string) (int64, error)
	CreateInvoked     bool
	UpdateFn          func(oldName, newname string) error
	UpdateInvoked     bool
	RemoveFn          func(name string) error
	RemoveInvoked     bool
}

// Generate random journal
func (r *MockRepository) Generate() *Journal {
	j := &Journal{}
	j.ID = random.Int64()
	j.Name = random.String(128)
	j.CreatedAt = random.Time()

	return j
}

// Get is mock implementation of journal.MockRepository.Get
func (r *MockRepository) Get(name string) (*Journal, error) {
	r.GetInvoked = true
	return r.GetFn(name)
}

// GetAll is mock implementation of journal.MockRepository.GetAll
func (r *MockRepository) GetAll() ([]Journal, error) {
	r.GetAllInvoked = true
	return r.GetAllFn()
}

// GetDefault is mock implementation of journal.MockRepository.GetDefault
func (r *MockRepository) GetDefault() (*Journal, error) {
	r.GetDefaultInvoked = true
	return r.GetDefaultFn()
}

// SetDefault is mock implementation of journal.MockRepository.SetDefault
func (r *MockRepository) SetDefault(name string) error {
	r.SetDefaultInvoked = true
	return r.SetDefaultFn(name)
}

// Create is mock implementation of journal.MockRepository.Create
func (r *MockRepository) Create(name string) (int64, error) {
	r.CreateInvoked = true
	return r.CreateFn(name)
}

// Update is mock implementation of journal.MockRepository.Update
func (r *MockRepository) Update(oldName, newName string) error {
	r.UpdateInvoked = true
	return r.UpdateFn(oldName, newName)
}

// Remove is mock implementation of journal.MockRepository.Remove
func (r *MockRepository) Remove(name string) error {
	r.RemoveInvoked = true
	return r.RemoveFn(name)
}
