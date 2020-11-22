package main

import (
	"testing"
)

func beforeEach(t *testing.T) *EntryRepo {
	db, err := OpenTestDB(t.Name())
	if err != nil {
		t.Error(err)
	}

	_, err = db.Exec(`
		INSERT INTO journal (name, created_at) values("test", "2020-01-02");
		INSERT INTO entry (journal_id, title, body, mood, created_at) 
			values(1, "testTitle", "testBody", "testMood", "2020-01-02");
		INSERT INTO tag (name, entry_id) values("testTag", 1);
	`)
	if err != nil {
		t.Error(err)
	}

	repo := MakeEntryRepo(db)

	return repo
}

func TestGet(t *testing.T) {
	repo := beforeEach(t)
	defer repo.db.Close()

	e, err := repo.Get(1)
	if err != nil {
		t.Error(err)
	}
	if e == nil {
		t.Error("nil entry")
	}
	if e.title != "testTitle" {
		t.Errorf("wrong title, expected %v received %v", "testTitle", e.title)
	}
	if e.body != "testBody" {
		t.Errorf("wrong body, expected %v received %v", "testBody", e.body)
	}
	if e.mood != "testMood" {
		t.Errorf("wrong mood, expected %v received %v", "testMood", e.mood)
	}
	if e.createdAt.Format("2006-01-02") != "2020-01-02" {
		t.Errorf("wrong createdAt, expected %v received %v", "2006-01-02", e.createdAt)
	}
	if len(e.tags) != 1 {
		t.Errorf("wrong tag length, expected %v received %v", 1, len(e.tags))
	}
	if e.tags[0] != "testTag" {
		t.Errorf("wrong tag, expected %v received %v", "testTag", e.tags[0])
	}
}
