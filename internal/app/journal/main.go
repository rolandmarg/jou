package journal

import (
	"fmt"
	"os"

	"github.com/rolandmarg/jou/internal/app/journal/note"
	"github.com/rolandmarg/jou/internal/platform/sqlite"
)

// Connect to database and return journal service
func Connect() *Service {
	db, err := sqlite.OpenDB()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open database: ", err)
		os.Exit(1)
	}
	n := note.MakeRepository(db)
	j := MakeRepository(db)

	s := MakeService(j, n)

	return s
}
