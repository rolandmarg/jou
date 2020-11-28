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
	if len(j.Entries) < 2 {
		t.Fatal("Expected to have entries")
	}
	for _, e := range j.Entries {
		if e.ID < 1 {
			t.Fatal("Expected entry ID to be positive")
		}
		if e.JournalID != j.ID {
			t.Fatalf("Expected journalID %v received %v ", j.ID, e.JournalID)
		}
		if e.Title[:len("testTitle")] != "testTitle" {
			t.Fatalf("Expected title to start with %v received %v", "testTitle", e.Title)
		}
		if e.Body[:len("testBody")] != "testBody" {
			t.Fatalf("Expected body to start with %v received %v", "testBody", e.Body)
		}
		if e.Mood[:len("testMood")] != "testMood" {
			t.Fatalf("Expected mood to start with %v received %v", "testMood", e.Mood)
		}
		if e.CreatedAt.IsZero() == true {
			t.Fatal("Expected createdAt to be set")
		}
		for _, tag := range e.Tags {
			if tag[:len("testTag")] != "testTag" {
				t.Fatalf("Expected tag name %v to start with %v %v", e.Tags[0], "testTag", e.Title)
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

	err := r.Create("xutu")
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

	err := r.Create("testJ")
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

func TestRemove(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	err := r.Create("testJ")
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
