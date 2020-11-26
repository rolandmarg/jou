package main

import (
	"testing"
)

func TestGetJournal(t *testing.T) {
	db, Teardown := Setup(t)
	defer Teardown()

	repo := MakeJournalRepo(db)

	journalID := int64(1)
	j, err := repo.Get(journalID)
	if err != nil {
		t.Fatal(err)
	}
	if j == nil {
		t.Fatal("Expected to exist")
	}
	if j.id != journalID {
		t.Fatalf("Expected id %v received %v", journalID, j.id)
	}
	if j.name[:len("test")] != "test" {
		t.Fatalf("Expected name %v received %v", "test", j.name)
	}
	if len(j.entries) < 2 {
		t.Fatal("Expected to have entries")
	}
	for _, e := range j.entries {
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

func TestCreateJournal(t *testing.T) {
	db, Teardown := Setup(t)
	defer Teardown()

	repo := MakeJournalRepo(db)

	id, err := repo.Create("testJ")
	if err != nil {
		t.Fatal(err)
	}

	j, err := repo.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	if j == nil {
		t.Fatal("Expected to exist")
	}
	if j.id != id {
		t.Fatalf("Expected id %v received %v", id, j.id)
	}
	if j.name[:len("test")] != "test" {
		t.Fatalf("Expected name %v received %v", "test", j.name)
	}
}

func TestRemoveJournal(t *testing.T) {
	db, Teardown := Setup(t)
	defer Teardown()

	repo := MakeJournalRepo(db)

	id, err := repo.Create("testJ")
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

	j, err := repo.Get(id)
	if err != nil {
		t.Fatal(err)
	}
	if j != nil {
		t.Fatal("Expected journal not to exist")
	}
}
