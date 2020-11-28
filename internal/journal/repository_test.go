package journal

import (
	"testing"

	"github.com/rolandmarg/jou/internal/pkg/fixture"
)

func TestGet(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	journalID := int64(1)
	j, err := r.Get(journalID)
	if err != nil {
		t.Fatal(err)
	}
	if j == nil {
		t.Fatal("Expected to exist")
	}
	if j.ID != journalID {
		t.Fatalf("Expected ID %v received %v", journalID, j.ID)
	}
	if j.name[:len("test")] != "test" {
		t.Fatalf("Expected name %v received %v", "test", j.name)
	}
	if len(j.entries) < 2 {
		t.Fatal("Expected to have entries")
	}
	for _, e := range j.entries {
		if e.ID < 1 {
			t.Fatal("Expected entry ID to be positive")
		}
		if e.JournalID != journalID {
			t.Fatalf("Expected journalID %v received %v ", journalID, e.JournalID)
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
func TestGetByName(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	journalID := int64(1)
	j, err := r.GetByName("test")
	if err != nil {
		t.Fatal(err)
	}
	if j == nil {
		t.Fatal("Expected to exist")
	}
	if j.ID != journalID {
		t.Fatalf("Expected ID %v received %v", journalID, j.ID)
	}
	if j.name != "test" {
		t.Fatalf("Expected name %v received %v", "test", j.name)
	}
	if len(j.entries) < 2 {
		t.Fatal("Expected to have entries")
	}
	for _, e := range j.entries {
		if e.ID < 1 {
			t.Fatal("Expected entry ID to be positive")
		}
		if e.JournalID != journalID {
			t.Fatalf("Expected journalID %v received %v ", journalID, e.JournalID)
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
	if j.name != "test" {
		t.Fatalf("Expected default journal %v received %v", "test", j.name)
	}
}

func TestUpdate(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	ID, err := r.Create("xutu")
	if err != nil {
		t.Fatal(err)
	}

	err = r.Update(ID, "butu")
	if err != nil {
		t.Fatal(err)
	}

	j, err := r.Get(ID)
	if err != nil {
		t.Fatal(err)
	}
	if j.name != "butu" {
		t.Fatalf("Expected name %v received %v", "butu", j.name)
	}
}
func TestCreate(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	ID, err := r.Create("testJ")
	if err != nil {
		t.Fatal(err)
	}

	j, err := r.Get(ID)
	if err != nil {
		t.Fatal(err)
	}

	if j == nil {
		t.Fatal("Expected to exist")
	}
	if j.ID != ID {
		t.Fatalf("Expected ID %v received %v", ID, j.ID)
	}
	if j.name[:len("test")] != "test" {
		t.Fatalf("Expected name %v received %v", "test", j.name)
	}
}

func TestRemove(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	ID, err := r.Create("testJ")
	if err != nil {
		t.Fatal(err)
	}

	_, err = r.Get(ID)
	if err != nil {
		t.Fatal(err)
	}

	err = r.Remove(ID)
	if err != nil {
		t.Fatal(err)
	}

	j, err := r.Get(ID)
	if err != nil {
		t.Fatal(err)
	}
	if j != nil {
		t.Fatal("Expected journal not to exist")
	}
}
