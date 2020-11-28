package note

import (
	"testing"

	"github.com/rolandmarg/jou/internal/pkg/fixture"
)

func TestGet(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	n, err := r.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	if n == nil {
		t.Fatal("Expected note to exist")
	}
	if n.ID < 1 {
		t.Fatal("Expected note ID to be positive")
	}
	if n.JournalID < 1 {
		t.Fatal("Expected note JournalID to be positive")
	}
	if n.Title != "testTitle" {
		t.Fatalf("Expected Title %v received %v", "testTitle", n.Title)
	}
	if n.Body != "testBody" {
		t.Fatalf("Expected Body %v received %v", "testBody", n.Body)
	}
	if n.Mood != "testMood" {
		t.Fatalf("Expected Mood %v received %v", "testMood", n.Mood)
	}
	if n.CreatedAt.IsZero() == true {
		t.Fatal("Expected CreatedAt to be set")
	}
	if len(n.Tags) < 2 {
		t.Fatal("Expected Tags to have elements")
	}
	for _, tag := range n.Tags {
		if tag[:len("testTag")] != "testTag" {
			t.Fatalf("Expected tag name %v to start with %v", n.Tags[0], "testTag")
		}
	}
}

func TestGetByJournalID(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	JournalID := int64(1)
	notes, err := r.GetByJournalID(JournalID)
	if err != nil {
		t.Fatal(err)
	}
	if notes == nil {
		t.Fatal("Expected note to exist")
	}
	if len(notes) < 2 {
		t.Fatal("Expected notes to have lements")
	}
	for _, n := range notes {
		if n.ID < 1 {
			t.Fatal("Expected note ID to be positive")
		}
		if n.JournalID != JournalID {
			t.Fatalf("Expected JournalID %v received %v ", JournalID, n.JournalID)
		}
		if n.Title[:len("testTitle")] != "testTitle" {
			t.Fatalf("Expected Title to start with %v received %v", "testTitle", n.Title)
		}
		if n.Body[:len("testBody")] != "testBody" {
			t.Fatalf("Expected Body to start with %v received %v", "testBody", n.Body)
		}
		if n.Mood[:len("testMood")] != "testMood" {
			t.Fatalf("Expected Mood to start with %v received %v", "testMood", n.Mood)
		}
		if n.CreatedAt.IsZero() == true {
			t.Fatal("Expected CreatedAt to be set")
		}
		for _, tag := range n.Tags {
			if tag[:len("testTag")] != "testTag" {
				t.Fatalf("Expected tag name %v to start with %v %v", n.Tags[0], "testTag", n.Title)
			}
		}
	}
}

func TestCreate(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	JournalID := int64(1)
	note := &Note{JournalID: JournalID, Title: "absolutelyRandomTitle", Body: "absolutely random Body",
		Mood: "notsomuch", Tags: []string{"randomtag1", "randomtag2"}}

	noteID, err := r.Create(note)
	if err != nil {
		t.Fatal(err)
	}
	if noteID < 1 {
		t.Fatal("Expected note ID to be positive")
	}

	n, err := r.Get(noteID)
	if err != nil {
		t.Fatal(err)
	}
	if n == nil {
		t.Fatal("Expected note to exist")
	}
	if n.ID < 1 {
		t.Fatal("Expected note ID to be positive")
	}
	if n.JournalID != JournalID {
		t.Fatalf("Expected JournalID %v received %v", JournalID, n.JournalID)
	}
	if n.Title != note.Title {
		t.Fatalf("Expected Title %v received %v", note.Title, n.Title)
	}
	if n.Body != note.Body {
		t.Fatalf("Expected Body %v received %v", note.Body, n.Body)
	}
	if n.Mood != note.Mood {
		t.Fatalf("Expected Mood %v received %v", note.Mood, n.Mood)
	}
	if n.CreatedAt.IsZero() == true {
		t.Fatal("Expected CreatedAt to be set")
	}
	if len(n.Tags) != len(note.Tags) {
		t.Fatalf("Expected Tags array len %v received %v", len(note.Tags), len(n.Tags))
	}
	for IDx := range note.Tags {
		if n.Tags[IDx] != note.Tags[IDx] {
			t.Fatalf("Expected tag name %v received %v", note.Tags[IDx], n.Tags[IDx])
		}
	}
}

func TestRemove(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	JournalID := int64(1)
	note := &Note{JournalID: JournalID, Title: "absolutelyRandomTitle", Body: "absolutely random Body",
		Mood: "notsomuch", Tags: []string{"randomtag1", "randomtag2"}}

	ID, err := r.Create(note)
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

	n, err := r.Get(ID)
	if err != nil {
		t.Fatal(err)
	}
	if n != nil {
		t.Fatal("Expected note not to exist")
	}
}
