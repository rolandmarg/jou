package journal

import (
	"fmt"
	"time"

	"github.com/rolandmarg/jou/internal/journal/entry"
)

// Journal structure
type Journal struct {
	ID        int64
	name      string
	entries   []entry.Entry
	createdAt time.Time
}

// Service provides operations on journal
type Service interface {
	Get(ID int64) (*Journal, error)
	GetByName(name string) (*Journal, error)
	GetDefault() (*Journal, error)
	GetAll() ([]Journal, error)
	SetDefault(name string) error
	Create(name string) (int64, error)
	Update(ID int64, name string) error
	Remove(ID int64) error
}

func (journal *Journal) String() string {
	str := fmt.Sprintln("journal:")

	for _, e := range journal.entries {
		str = fmt.Sprintln(str, " entry:")
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

	return str
}
