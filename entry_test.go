package main

import (
	"testing"
)

func TestGetEntry(t *testing.T) {
	db, Teardown := Setup(t)
	defer Teardown()

	repo := MakeEntryRepo(db)

	e, err := repo.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	if e == nil {
		t.Fatal("Expected entry to exist")
	}
	if e.id < 1 {
		t.Fatal("Expected entry id to be positive")
	}
	if e.journalID < 1 {
		t.Fatal("Expected entry journalID to be positive")
	}
	if e.title != "testTitle" {
		t.Fatalf("Expected title %v received %v", "testTitle", e.title)
	}
	if e.body != "testBody" {
		t.Fatalf("Expected body %v received %v", "testBody", e.body)
	}
	if e.mood != "testMood" {
		t.Fatalf("Expected mood %v received %v", "testMood", e.mood)
	}
	if e.createdAt.IsZero() == true {
		t.Fatal("Expected createdAt to be set")
	}
	if len(e.tags) == 0 {
		t.Fatal("Expected tags array to have elements")
	}
	for _, tag := range e.tags {
		if tag.id < 1 {
			t.Fatal("Expected tag id to be positive")
		}
		if tag.entryID != e.id {
			t.Fatalf("Expected tag id %v received %v", e.id, tag.entryID)
		}
		if tag.name[:len("testTag")] != "testTag" {
			t.Fatalf("Expected tag name %v to start with %v", e.tags[0], "testTag")
		}
	}
}

func TestCreateEntry(t *testing.T) {
	db, Teardown := Setup(t)
	defer Teardown()

	repo := MakeEntryRepo(db)

	journalID := int64(1)
	i := &EntryInput{title: "absolutelyRandomTitle", body: "absolutely random body",
		mood: "notsomuch", tags: []string{"randomtag1", "randomtag2"}}

	entryID, err := repo.Create(journalID, i)
	if err != nil {
		t.Fatal(err)
	}
	if entryID < 1 {
		t.Fatal("Expected entry id to be positive")
	}

	e, err := repo.Get(entryID)
	if err != nil {
		t.Fatal(err)
	}
	if e == nil {
		t.Fatal("Expected entry to exist")
	}
	if e.id < 1 {
		t.Fatal("Expected entry id to be positive")
	}
	if e.journalID != journalID {
		t.Fatalf("Expected journalID %v received %v", journalID, e.journalID)
	}
	if e.title != i.title {
		t.Fatalf("Expected title %v received %v", i.title, e.title)
	}
	if e.body != i.body {
		t.Fatalf("Expected body %v received %v", i.body, e.body)
	}
	if e.mood != i.mood {
		t.Fatalf("Expected mood %v received %v", i.mood, e.mood)
	}
	if e.createdAt.IsZero() == true {
		t.Fatal("Expected createdAt to be set")
	}
	if len(e.tags) != len(i.tags) {
		t.Fatalf("Expected tags array len %v received %v", len(i.tags), len(e.tags))
	}
	for idx := range i.tags {
		if e.tags[idx].id < 1 {
			t.Fatal("Expected tag id to be positive")
		}
		if e.tags[idx].entryID != e.id {
			t.Fatalf("Expected tag id %v received %v", e.id, e.tags[idx].entryID)
		}
		if e.tags[idx].name != i.tags[idx] {
			t.Fatalf("Expected tag name %v received %v", i.tags[idx], e.tags[idx].name)
		}
	}
}
