package journal

import (
	"testing"

	"github.com/rolandmarg/jou/internal/pkg/fixture"
)

func TestGet(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	name := "test"
	j, err := r.Get(name)
	if err != nil {
		t.Fatal(err)
	}
	if j == nil {
		t.Fatal("Expected to exist")
	}
	if j.ID < 1 {
		t.Fatal("Expected ID to be positive")
	}
	if j.Name[:len("test")] != "test" {
		t.Fatalf("Expected name %v received %v", "test", j.Name)
	}
	if len(j.Notes) < 2 {
		t.Fatal("Expected to have notes")
	}
	for _, n := range j.Notes {
		if n.ID < 1 {
			t.Fatal("Expected note ID to be positive")
		}
		if n.JournalID != j.ID {
			t.Fatalf("Expected journalID %v received %v ", j.ID, n.JournalID)
		}
		if n.Title[:len("testTitle")] != "testTitle" {
			t.Fatalf("Expected title to start with %v received %v", "testTitle", n.Title)
		}
		if n.Body[:len("testBody")] != "testBody" {
			t.Fatalf("Expected body to start with %v received %v", "testBody", n.Body)
		}
		if n.Mood[:len("testMood")] != "testMood" {
			t.Fatalf("Expected mood to start with %v received %v", "testMood", n.Mood)
		}
		if n.CreatedAt.IsZero() == true {
			t.Fatal("Expected createdAt to be set")
		}
		for _, tag := range n.Tags {
			if tag[:len("testTag")] != "testTag" {
				t.Fatalf("Expected tag name %v to start with %v %v", n.Tags[0], "testTag", n.Title)
			}
		}
	}
}

func TestSetDefault(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	err := r.SetDefault("test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetDefault(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	err := r.SetDefault("test")
	if err != nil {
		t.Fatal(err)
	}

	j, err := r.GetDefault()
	if err != nil {
		t.Fatal(err)
	}
	if j.Name != "test" {
		t.Fatalf("Expected default journal %v received %v", "test", j.Name)
	}
}

func TestUpdate(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	_, err := r.Create("xutu")
	if err != nil {
		t.Fatal(err)
	}

	err = r.Update("xutu", "butu")
	if err != nil {
		t.Fatal(err)
	}

	j, err := r.Get("butu")
	if err != nil {
		t.Fatal(err)
	}
	j, err = r.Get("xutu")
	if err != nil {
		t.Fatal(err)
	}
	if j != nil {
		t.Fatal("Expected journal xutu not to exist")
	}
}
func TestCreate(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	_, err := r.Create("testJ")
	if err != nil {
		t.Fatal(err)
	}

	j, err := r.Get("testJ")
	if err != nil {
		t.Fatal(err)
	}

	if j == nil {
		t.Fatal("Expected to exist")
	}
	if j.Name != "testJ" {
		t.Fatalf("Expected name %v received %v", "testJ", j.Name)
	}
}

func TestCreateDuplicate(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	_, err := r.Create("dup")
	if err != nil {
		t.Fatal(err)
	}

	_, err = r.Create("dup")
	if err == nil {
		t.Fatal("Expected journal create to fail on duplicate name")
	}
}

func TestRemove(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	_, err := r.Create("testJ")
	if err != nil {
		t.Fatal(err)
	}

	_, err = r.Get("testJ")
	if err != nil {
		t.Fatal(err)
	}

	err = r.Remove("testJ")
	if err != nil {
		t.Fatal(err)
	}

	j, err := r.Get("testJ")
	if err != nil {
		t.Fatal(err)
	}
	if j != nil {
		t.Fatal("Expected journal testJ not to exist")
	}
}
