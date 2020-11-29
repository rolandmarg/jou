package journal

import (
	"fmt"
	"os"

	// TODO THIS IS UGLY
	jr "github.com/rolandmarg/jou/internal/app/journal/sqlite"
	nr "github.com/rolandmarg/jou/internal/app/note/sqlite"
	"github.com/rolandmarg/jou/internal/platform/sqlite"
)

// Connect to database and return journal service
func Connect() *Service {
	db, err := sqlite.OpenDB()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not open database: ", err)
		os.Exit(1)
	}
	n := nr.MakeRepository(db)
	j := jr.MakeRepository(db)

	s := MakeService(j, n)

	return s
}
