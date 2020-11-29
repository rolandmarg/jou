package sqlite

import (
	"fmt"
	"testing"

	"github.com/rolandmarg/jou/internal/pkg/journal"
	"github.com/rolandmarg/jou/internal/platform/sqlite"
)

func setup(t *testing.T) (journal.Repository, func()) {
	name := fmt.Sprintf("file:%v.db?cache=shared&mode=memory", t.Name())
	db, e := sqlite.Open(name)
	her(t, e)

	_, e = db.Exec(`
		INSERT INTO journal (name, created_at) VALUES ("test", "2020-01-01");
		INSERT INTO journal (name, created_at) VALUES ("test2", "2020-01-02");
	`)
	her(t, e)

	r := MakeRepository(db)

	return r, func() {
		db.Close()
	}
}

func her(t *testing.T, args ...interface{}) {
	if len(args) != 0 && args[0] != nil {
		t.Fatal(args...)
	}
}

func TestGet(t *testing.T) {
	r, teardown := setup(t)
	defer teardown()

	name := "test"
	j, e := r.Get(name)
	her(t, e)

	if j == nil {
		her(t, "Expected to exist")
	}
	if j.ID < 1 {
		her(t, "Expected ID to be positive")
	}
	if j.Name[:len("test")] != "test" {
		her(t, "Expected name 'test' received ", j.Name)
	}
}

func TestSetDefault(t *testing.T) {
	r, teardown := setup(t)
	defer teardown()

	e := r.SetDefault("test")
	her(t, e)
}

func TestGetDefault(t *testing.T) {
	r, teardown := setup(t)
	defer teardown()

	e := r.SetDefault("test")
	her(t, e)

	j, e := r.GetDefault()
	her(t, e)

	if j.Name != "test" {
		her(t, "Expected default journal 'test' received ", j.Name)
	}
}

func TestUpdate(t *testing.T) {
	r, teardown := setup(t)
	defer teardown()

	_, e := r.Create("xutu")
	her(t, e)

	e = r.Update("xutu", "butu")
	her(t, e)

	j, e := r.Get("butu")
	her(t, e)

	j, e = r.Get("xutu")
	her(t, e)

	if j != nil {
		her(t, "Expected journal xutu not to exist")
	}
}
func TestCreate(t *testing.T) {
	r, teardown := setup(t)
	defer teardown()

	_, e := r.Create("testJ")
	her(t, e)

	j, e := r.Get("testJ")
	her(t, e)

	if j == nil {
		her(t, "Expected to exist")
	}
	if j.Name != "testJ" {
		her(t, "Expected name 'testJ' received ", j.Name)
	}
}

func TestCreateDuplicate(t *testing.T) {
	r, teardown := setup(t)
	defer teardown()

	_, e := r.Create("dup")
	her(t, e)

	_, e = r.Create("dup")
	if e == nil {
		her(t, "Expected journal create to fail on duplicate name")
	}
}

func TestRemove(t *testing.T) {
	r, teardown := setup(t)
	defer teardown()

	_, e := r.Create("testJ")
	her(t, e)

	_, e = r.Get("testJ")
	her(t, e)

	e = r.Remove("testJ")
	her(t, e)

	j, e := r.Get("testJ")
	her(t, e)

	if j != nil {
		her(t, "Expected journal testJ not to exist")
	}
}
