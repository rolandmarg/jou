package main

import (
	"database/sql"
	"testing"
)

func Setup(t *testing.T) (*sql.DB, func()) {
	db, err := OpenTestDB(t.Name())
	if err != nil {
		t.Fatal(err)
	}

	// TODO maybe generate random data
	_, err = db.Exec(`
		INSERT INTO journal (name, created_at) values("test", "2020-01-01");
		INSERT INTO journal (name, created_at) values("test2", "2020-01-02");
		
		INSERT INTO entry (journal_id, title, body, mood, created_at) 
			values(1, "testTitle", "testBody", "testMood", "2020-01-01");
		INSERT INTO entry (journal_id, title, body, mood, created_at) 
			values(1, "testTitle2", "testBody2", "testMood2", "2020-01-02");
		INSERT INTO entry (journal_id, title, body, mood, created_at) 
			values(1, "testTitl3", "testBody3", "testMood3", "2020-01-03");
		INSERT INTO entry (journal_id, title, body, mood, created_at) 
			values(1, "testTitle4", "testBody4", "testMood4", "2020-01-04");
		INSERT INTO entry (journal_id, title, body, mood, created_at) 
			values(1, "testTitle5", "testBody5", "testMood5", "2020-01-05");
		INSERT INTO entry (journal_id, title, body, mood, created_at) 
			values(2, "testTitle", "testBody", "testMood", "2020-01-01");
		INSERT INTO entry (journal_id, title, body, mood, created_at) 
			values(2, "testTitle2", "testBody2", "testMood2", "2020-01-02");
		INSERT INTO entry (journal_id, title, body, mood, created_at) 
			values(2, "testTitle3", "testBody3", "testMood3", "2020-01-03");

		INSERT INTO tag (name, entry_id) values("testTag", 1);
		INSERT INTO tag (name, entry_id) values("testTag2", 1);
		INSERT INTO tag (name, entry_id) values("testTag3", 1);
		INSERT INTO tag (name, entry_id) values("testTag", 2);
		INSERT INTO tag (name, entry_id) values("testTag2", 2);
		INSERT INTO tag (name, entry_id) values("testTag3", 3);
		INSERT INTO tag (name, entry_id) values("testTag4", 4);
	`)
	if err != nil {
		t.Error(err)
	}

	return db, func() {
		db.Close()
	}
}
