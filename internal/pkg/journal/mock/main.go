package mock

// TODO don't know if importing parent is good idea
import (
	"github.com/rolandmarg/jou/internal/pkg/journal"
)

// Repository is mock implementation of journal repository
type Repository struct {
	GetFn             func(name string) (*journal.Journal, error)
	GetInvoked        bool
	GetAllFn          func() ([]journal.Journal, error)
	GetAllInvoked     bool
	GetDefaultFn      func() (*journal.Journal, error)
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

// Get is mock implementation of journal.Repository.Get
func (r *Repository) Get(name string) (*journal.Journal, error) {
	r.GetInvoked = true
	return r.GetFn(name)
}

// GetAll is mock implementation of journal.Repository.GetAll
func (r *Repository) GetAll() ([]journal.Journal, error) {
	r.GetAllInvoked = true
	return r.GetAllFn()
}

// GetDefault is mock implementation of journal.Repository.GetDefault
func (r *Repository) GetDefault() (*journal.Journal, error) {
	r.GetDefaultInvoked = true
	return r.GetDefaultFn()
}

// SetDefault is mock implementation of journal.Repository.SetDefault
func (r *Repository) SetDefault(name string) error {
	r.SetDefaultInvoked = true
	return r.SetDefaultFn(name)
}

// Create is mock implementation of journal.Repository.Create
func (r *Repository) Create(name string) (int64, error) {
	r.CreateInvoked = true
	return r.CreateFn(name)
}

// Update is mock implementation of journal.Repository.Update
func (r *Repository) Update(oldName, newName string) error {
	r.UpdateInvoked = true
	return r.UpdateFn(oldName, newName)
}

// Remove is mock implementation of journal.Repository.Remove
func (r *Repository) Remove(name string) error {
	r.RemoveInvoked = true
	return r.RemoveFn(name)
}
