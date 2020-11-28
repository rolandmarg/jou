package fixture

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/rolandmarg/jou/internal/platform/sqlite"
)

// Setup initializes database with fixture data and returns DB, teardown func
func Setup(t *testing.T) (*sql.DB, func()) {
	name := fmt.Sprintf("file:%v.db?cache=shared&mode=memory", t.Name())
	DB, err := sqlite.Open(name)
	if err != nil {
		t.Fatal(err)
	}

	// TODO maybe generate random data
	_, err = DB.Exec(`
		INSERT INTO journal (name, created_at) VALUES ("test", "2020-01-01");
		INSERT INTO journal (name, created_at) VALUES ("test2", "2020-01-02");
		
		INSERT INTO default_journal (name, j_id) VALUES ("default", 1);

		INSERT INTO note (j_id, title, body, mood, created_at) 
			VALUES (1, "testTitle", "testBody", "testMood", "2020-01-01");
		INSERT INTO note (j_id, title, body, mood, created_at) 
			VALUES (1, "testTitle2", "testBody2", "testMood2", "2020-01-02");
		INSERT INTO note (j_id, title, body, mood, created_at) 
			VALUES (1, "testTitle3", "testBody3", "testMood3", "2020-01-03");
		INSERT INTO note (j_id, title, body, mood, created_at) 
			VALUES (1, "testTitle4", "testBody4", "testMood4", "2020-01-04");
		INSERT INTO note (j_id, title, body, mood, created_at) 
			VALUES (1, "testTitle5", "testBody5", "testMood5", "2020-01-05");
		INSERT INTO note (j_id, title, body, mood, created_at) 
			VALUES (2, "testTitle", "testBody", "testMood", "2020-01-01");
		INSERT INTO note (j_id, title, body, mood, created_at) 
			VALUES (2, "testTitle2", "testBody2", "testMood2", "2020-01-02");
		INSERT INTO note (j_id, title, body, mood, created_at) 
			VALUES (2, "testTitle3", "testBody3", "testMood3", "2020-01-03");

		INSERT INTO tag (name, note_id) VALUES ("testTag", 1);
		INSERT INTO tag (name, note_id) VALUES ("testTag2", 1);
		INSERT INTO tag (name, note_id) VALUES ("testTag3", 1);
		INSERT INTO tag (name, note_id) VALUES ("testTag", 2);
		INSERT INTO tag (name, note_id) VALUES ("testTag2", 2);
		INSERT INTO tag (name, note_id) VALUES ("testTag3", 3);
		INSERT INTO tag (name, note_id) VALUES ("testTag4", 4);

		INSERT INTO env (key, value) VALUES ("xutkunchula", "shivdkunchula");
		INSERT INTO env (key, value) VALUES ("testName", "testValue");
	`)
	if err != nil {
		t.Error(err)
	}

	return DB, func() {
		DB.Close()
	}
}
