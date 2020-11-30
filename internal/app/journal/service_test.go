package journal

import (
	"testing"

	// TODO ugly
	"github.com/rolandmarg/jou/internal/pkg/journal"
	jm "github.com/rolandmarg/jou/internal/pkg/journal/mock"
	nm "github.com/rolandmarg/jou/internal/pkg/note/mock"
)

func setup() (*Service, *jm.Repository, *nm.Repository) {
	j := &jm.Repository{}
	n := &nm.Repository{}
	s := MakeService(j, n)

	return s, j, n
}
func TestGet(t *testing.T) {
	s, j, _ := setup()

	t.Run("Should fail on non existing Get", func(t *testing.T) {
		j.GetFn = func(name string) (*journal.Journal, error) { return &journal.Journal{}, nil }
		j, err := s.Get("random")
		if err == nil {
			t.Fatal("No error returned")
		}
		if j != nil {
			t.Fatal("Existing journal returned from space time continuum", j)
		}
	})
}
