package journal

import (
	"fmt"
	"time"

	"github.com/rolandmarg/jou/internal/pkg/note"
)

// Journal structure
type Journal struct {
	ID        int64
	Name      string
	Notes     []note.Note
	CreatedAt time.Time
}

// Repository provides operations on journal DAL
type Repository interface {
	Get(name string) (*Journal, error)
	GetDefault() (*Journal, error)
	GetAll() ([]Journal, error)
	SetDefault(name string) error
	Create(name string) (int64, error)
	Update(oldName string, newName string) error
	Remove(name string) error
}

func (j Journal) String() string {
	str := fmt.Sprintf(`journal "%v" notes: [`, j.Name)
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
