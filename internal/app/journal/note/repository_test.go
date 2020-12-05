package note

import (
	"testing"

	"github.com/rolandmarg/jou/internal/pkg/journal/note"
	"github.com/rolandmarg/jou/internal/platform/sqlite"
)

func setup(t *testing.T) (note.Repository, func()) {
	fixture := `
	INSERT INTO note (j_id, title, body, mood, created_at) 
		VALUES (1, "testTitle", "testBody", "testMood", "2020-01-01");
	INSERT INTO note (j_id, title, body, mood, created_at) 
		VALUES (1, "testTitle2", "testBody2", "testMood2", "2020-01-02");
	INSERT INTO note (j_id, title, body, mood, created_at) 
		VALUES (2, "testTitle3", "testBody3", "testMood3", "2020-01-03");

	INSERT INTO tag (name, note_id) VALUES ("testTag", 1);
	INSERT INTO tag (name, note_id) VALUES ("testTag2", 1);
	INSERT INTO tag (name, note_id) VALUES ("testTag", 2);	
`
	db, e := sqlite.OpenTestDB(t.Name(), fixture)
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

	n, e := r.Get(1)
	her(t, e)

	if n == nil {
		her(t, "Expected note to exist")
	}
	if n.ID < 1 {
		her(t, "Expected note ID to be positive")
	}
	if n.JournalID < 1 {
		her(t, "Expected note JournalID to be positive")
	}
	if n.Title != "testTitle" {
		her(t, "Expected Title 'testTitle' received ", n.Title)
	}
	if n.Body != "testBody" {
		her(t, "Expected Body 'testBody' received ", n.Body)
	}
	if n.Mood != "testMood" {
		her(t, "Expected Mood 'testMood' received ", n.Mood)
	}
	if n.CreatedAt.IsZero() == true {
		her(t, "Expected CreatedAt to be set")
	}
	if len(n.Tags) < 2 {
		her(t, "Expected Tags to have elements")
	}
	for _, tag := range n.Tags {
		if tag[:len("testTag")] != "testTag" {
			her(t, "Expected tag name to start with 'testTag' received ", n.Tags[0])
		}
	}
}

func TestGetByJournalID(t *testing.T) {
	r, teardown := setup(t)
	defer teardown()

	journalID := int64(1)
	notes, e := r.GetByJournalID(journalID)
	if e != nil {
		her(t, e)
	}
	if notes == nil {
		her(t, "Expected note to exist")
	}
	if len(notes) < 2 {
		her(t, "Expected notes to have lements")
	}
	for _, n := range notes {
		if n.ID < 1 {
			her(t, "Expected note ID to be positive")
		}
		if n.JournalID != journalID {
			her(t, "Expected JournalID", journalID, "received ", n.JournalID)
		}
		if n.Title[:len("testTitle")] != "testTitle" {
			her(t, "Expected Title to start with 'testTitle' received ", n.Title)
		}
		if n.Body[:len("testBody")] != "testBody" {
			her(t, "Expected Body to start with 'testBody' received ", n.Body)
		}
		if n.Mood[:len("testMood")] != "testMood" {
			her(t, "Expected Mood to start with 'testMood' received ", n.Mood)
		}
		if n.CreatedAt.IsZero() == true {
			her(t, "Expected CreatedAt to be set")
		}
		for _, tag := range n.Tags {
			if tag[:len("testTag")] != "testTag" {
				her(t, "Expected tag name to start with 'testTag' received ", n.Tags[0])
			}
		}
	}
}

func TestCreate(t *testing.T) {
	r, teardown := setup(t)
	defer teardown()

	journalID := int64(1)
	title := "absolutelyRandomTitle"
	body := "absolutely random body"
	mood := "not so much"
	tags := []string{"randomTag1", "randomTag2"}

	noteID, e := r.Create(journalID, title, body, mood, tags)
	if e != nil {
		her(t, e)
	}
	if noteID < 1 {
		her(t, "Expected note ID to be positive")
	}

	n, e := r.Get(noteID)
	if e != nil {
		her(t, e)
	}
	if n == nil {
		her(t, "Expected note to exist")
	}
	if n.ID < 1 {
		her(t, "Expected note ID to be positive")
	}
	if n.JournalID != journalID {
		her(t, "Expected JournalID ", journalID, "received ", n.JournalID)
	}
	if n.Title != title {
		her(t, "Expected Title ", title, "received ", n.Title)
	}
	if n.Body != body {
		her(t, "Expected Body ", body, "received ", n.Body)
	}
	if n.Mood != mood {
		her(t, "Expected Mood ", mood, "received ", n.Mood)
	}
	if n.CreatedAt.IsZero() == true {
		her(t, "Expected CreatedAt to be set")
	}
	if len(n.Tags) != len(tags) {
		her(t, "Expected Tags array len ", len(tags), "received ", len(n.Tags))
	}
	for IDx := range tags {
		if n.Tags[IDx] != tags[IDx] {
			her(t, "Expected tag name ", tags[IDx], " received ", n.Tags[IDx])
		}
	}
}

func TestRemove(t *testing.T) {
	r, teardown := setup(t)
	defer teardown()

	journalID := int64(1)
	title := "absolutelyRandomTitle"
	body := "absolutely random body"
	mood := "not so much"
	tags := []string{"randomTag1", "randomTag2"}

	id, e := r.Create(journalID, title, body, mood, tags)
	if e != nil {
		her(t, e)
	}

	_, e = r.Get(id)
	if e != nil {
		her(t, e)
	}

	e = r.Remove(id)
	if e != nil {
		her(t, e)
	}

	n, e := r.Get(id)
	if e != nil {
		her(t, e)
	}
	if n != nil {
		her(t, "Expected note not to exist")
	}
}
