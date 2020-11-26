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
	if len(e.tags) < 2 {
		t.Fatal("Expected tags to have elements")
	}
	for _, tag := range e.tags {
		if tag[:len("testTag")] != "testTag" {
			t.Fatalf("Expected tag name %v to start with %v", e.tags[0], "testTag")
		}
	}
}

func TestGetEntryByJournalID(t *testing.T) {
	db, Teardown := Setup(t)
	defer Teardown()

	repo := MakeEntryRepo(db)

	journalID := int64(1)
	entries, err := repo.GetByJournalID(journalID)
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
		if e.id < 1 {
			t.Fatal("Expected entry id to be positive")
		}
		if e.journalID != journalID {
			t.Fatalf("Expected journalID %v received %v ", journalID, e.journalID)
		}
		if e.title[:len("testTitle")] != "testTitle" {
			t.Fatalf("Expected title to start with %v received %v", "testTitle", e.title)
		}
		if e.body[:len("testBody")] != "testBody" {
			t.Fatalf("Expected body to start with %v received %v", "testBody", e.body)
		}
		if e.mood[:len("testMood")] != "testMood" {
			t.Fatalf("Expected mood to start with %v received %v", "testMood", e.mood)
		}
		if e.createdAt.IsZero() == true {
			t.Fatal("Expected createdAt to be set")
		}
		for _, tag := range e.tags {
			if tag[:len("testTag")] != "testTag" {
				t.Fatalf("Expected tag name %v to start with %v %v", e.tags[0], "testTag", e.title)
			}
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
		if e.tags[idx] != i.tags[idx] {
			t.Fatalf("Expected tag name %v received %v", i.tags[idx], e.tags[idx])
		}
	}
}

func TestRemoveEntry(t *testing.T) {
	db, Teardown := Setup(t)
	defer Teardown()

	repo := MakeEntryRepo(db)

	journalID := int64(1)
	i := &EntryInput{title: "absolutelyRandomTitle", body: "absolutely random body",
		mood: "notsomuch", tags: []string{"randomtag1", "randomtag2"}}

	id, err := repo.Create(journalID, i)
	if err != nil {
		t.Fatal(err)
	}

	_, err = repo.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Remove(id)
	if err != nil {
		t.Fatal(err)
	}

	entry, err := repo.Get(id)
	if err != nil {
		t.Fatal(err)
	}
	if entry != nil {
		t.Fatal("Expected entry not to exist")
	}
}
