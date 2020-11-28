package journal

import (
	"fmt"
	"time"

	"github.com/rolandmarg/jou/internal/journal/note"
)

// Journal structure
type Journal struct {
	ID        int64
	Name      string
	Notes   []note.Note
	CreatedAt time.Time
}

// Service provides operations on journal
type Service interface {
	Get(name string) (*Journal, error)
	GetDefault() (*Journal, error)
	GetAll() ([]Journal, error)
	SetDefault(name string) error
	Create(name string) error
	Update(oldName string, newName string) error
	Remove(name string) error
}

func (j *Journal) String() string {
	str := fmt.Sprintln("journal", j.Name)

	str = fmt.Sprint(str, "notes: [")
	for _, e := range j.Notes {
		str = fmt.Sprintln(str, " note:")
		str = fmt.Sprintln(str, "   id:", e.ID)
		str = fmt.Sprintln(str, "   title:", e.Title)
		str = fmt.Sprintln(str, "   body:", e.Body)
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
