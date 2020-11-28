package entry

import (
	"testing"

	"github.com/rolandmarg/jou/internal/pkg/fixture"
)

func TestGet(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	e, err := r.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	if e == nil {
		t.Fatal("Expected entry to exist")
	}
	if e.ID < 1 {
		t.Fatal("Expected entry ID to be positive")
	}
	if e.JournalID < 1 {
		t.Fatal("Expected entry JournalID to be positive")
	}
	if e.Title != "testTitle" {
		t.Fatalf("Expected Title %v received %v", "testTitle", e.Title)
	}
	if e.Body != "testBody" {
		t.Fatalf("Expected Body %v received %v", "testBody", e.Body)
	}
	if e.Mood != "testMood" {
		t.Fatalf("Expected Mood %v received %v", "testMood", e.Mood)
	}
	if e.CreatedAt.IsZero() == true {
		t.Fatal("Expected CreatedAt to be set")
	}
	if len(e.Tags) < 2 {
		t.Fatal("Expected Tags to have elements")
	}
	for _, tag := range e.Tags {
		if tag[:len("testTag")] != "testTag" {
			t.Fatalf("Expected tag name %v to start with %v", e.Tags[0], "testTag")
		}
	}
}

func TestGetByJournalID(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	JournalID := int64(1)
	entries, err := r.GetByJournalID(JournalID)
	if err != nil {
		t.Fatal(err)
	}
	if entries == nil {
		t.Fatal("Expected entry to exist")
	}
	if len(entries) < 2 {
		t.Fatal("Expected entries to have lements")
	}
	for _, e := range entries {
		if e.ID < 1 {
			t.Fatal("Expected entry ID to be positive")
		}
		if e.JournalID != JournalID {
			t.Fatalf("Expected JournalID %v received %v ", JournalID, e.JournalID)
		}
		if e.Title[:len("testTitle")] != "testTitle" {
			t.Fatalf("Expected Title to start with %v received %v", "testTitle", e.Title)
		}
		if e.Body[:len("testBody")] != "testBody" {
			t.Fatalf("Expected Body to start with %v received %v", "testBody", e.Body)
		}
		if e.Mood[:len("testMood")] != "testMood" {
			t.Fatalf("Expected Mood to start with %v received %v", "testMood", e.Mood)
		}
		if e.CreatedAt.IsZero() == true {
			t.Fatal("Expected CreatedAt to be set")
		}
		for _, tag := range e.Tags {
			if tag[:len("testTag")] != "testTag" {
				t.Fatalf("Expected tag name %v to start with %v %v", e.Tags[0], "testTag", e.Title)
			}
		}
	}
}

func TestCreate(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	JournalID := int64(1)
	entry := &Entry{JournalID: JournalID, Title: "absolutelyRandomTitle", Body: "absolutely random Body",
		Mood: "notsomuch", Tags: []string{"randomtag1", "randomtag2"}}

	entryID, err := r.Create(entry)
	if err != nil {
		t.Fatal(err)
	}
	if entryID < 1 {
		t.Fatal("Expected entry ID to be positive")
	}

	e, err := r.Get(entryID)
	if err != nil {
		t.Fatal(err)
	}
	if e == nil {
		t.Fatal("Expected entry to exist")
	}
	if e.ID < 1 {
		t.Fatal("Expected entry ID to be positive")
	}
	if e.JournalID != JournalID {
		t.Fatalf("Expected JournalID %v received %v", JournalID, e.JournalID)
	}
	if e.Title != entry.Title {
		t.Fatalf("Expected Title %v received %v", entry.Title, e.Title)
	}
	if e.Body != entry.Body {
		t.Fatalf("Expected Body %v received %v", entry.Body, e.Body)
	}
	if e.Mood != entry.Mood {
		t.Fatalf("Expected Mood %v received %v", entry.Mood, e.Mood)
	}
	if e.CreatedAt.IsZero() == true {
		t.Fatal("Expected CreatedAt to be set")
	}
	if len(e.Tags) != len(entry.Tags) {
		t.Fatalf("Expected Tags array len %v received %v", len(entry.Tags), len(e.Tags))
	}
	for IDx := range entry.Tags {
		if e.Tags[IDx] != entry.Tags[IDx] {
			t.Fatalf("Expected tag name %v received %v", entry.Tags[IDx], e.Tags[IDx])
		}
	}
}

func TestRemove(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	JournalID := int64(1)
	entry := &Entry{JournalID: JournalID, Title: "absolutelyRandomTitle", Body: "absolutely random Body",
		Mood: "notsomuch", Tags: []string{"randomtag1", "randomtag2"}}

	ID, err := r.Create(entry)
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

	e, err := r.Get(ID)
	if err != nil {
		t.Fatal(err)
	}
	if e != nil {
		t.Fatal("Expected entry not to exist")
	}
}
