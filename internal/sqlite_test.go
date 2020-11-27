package main

import (
	"testing"
)

func TestJournalGet(t *testing.T) {
	DB, Teardown := Setup(t)
	defer Teardown()

	entryService := CreateEntryRepository(DB)
	keyvalueService := CreateKVRepository(DB)
	r := CreateJournalRepository(DB, entryService, keyvalueService)

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
func TestJournalGetByName(t *testing.T) {
	DB, Teardown := Setup(t)
	defer Teardown()

	entryService := CreateEntryRepository(DB)
	keyvalueService := CreateKVRepository(DB)
	r := CreateJournalRepository(DB, entryService, keyvalueService)

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

func TestJournalSetDefault(t *testing.T) {
	DB, Teardown := Setup(t)
	defer Teardown()

	entryService := CreateEntryRepository(DB)
	keyvalueService := CreateKVRepository(DB)
	r := CreateJournalRepository(DB, entryService, keyvalueService)

	err := r.SetDefault("test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestJournalGetDefault(t *testing.T) {
	DB, Teardown := Setup(t)
	defer Teardown()

	entryService := CreateEntryRepository(DB)
	keyvalueService := CreateKVRepository(DB)
	r := CreateJournalRepository(DB, entryService, keyvalueService)

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

func TestJournalUpdate(t *testing.T) {
	DB, Teardown := Setup(t)
	defer Teardown()

	entryService := CreateEntryRepository(DB)
	keyvalueService := CreateKVRepository(DB)
	r := CreateJournalRepository(DB, entryService, keyvalueService)

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
func TestJournalCreate(t *testing.T) {
	DB, Teardown := Setup(t)
	defer Teardown()

	entryService := CreateEntryRepository(DB)
	keyvalueService := CreateKVRepository(DB)
	r := CreateJournalRepository(DB, entryService, keyvalueService)

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

func TestJournalRemove(t *testing.T) {
	DB, Teardown := Setup(t)
	defer Teardown()

	entryService := CreateEntryRepository(DB)
	keyvalueService := CreateKVRepository(DB)
	r := CreateJournalRepository(DB, entryService, keyvalueService)

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

func TestEntryGet(t *testing.T) {
	DB, Teardown := Setup(t)
	defer Teardown()

	r := CreateEntryRepository(DB)

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

func TestEntryGetByJournalID(t *testing.T) {
	DB, Teardown := Setup(t)
	defer Teardown()

	r := CreateEntryRepository(DB)

	journalID := int64(1)
	entries, err := r.GetByJournalID(journalID)
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

func TestEntryCreate(t *testing.T) {
	DB, Teardown := Setup(t)
	defer Teardown()

	r := CreateEntryRepository(DB)

	journalID := int64(1)
	entry := &Entry{journalID: journalID, title: "absolutelyRandomTitle", body: "absolutely random body",
		mood: "notsomuch", tags: []string{"randomtag1", "randomtag2"}}

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
	if e.journalID != journalID {
		t.Fatalf("Expected journalID %v received %v", journalID, e.journalID)
	}
	if e.title != entry.title {
		t.Fatalf("Expected title %v received %v", entry.title, e.title)
	}
	if e.body != entry.body {
		t.Fatalf("Expected body %v received %v", entry.body, e.body)
	}
	if e.mood != entry.mood {
		t.Fatalf("Expected mood %v received %v", entry.mood, e.mood)
	}
	if e.createdAt.IsZero() == true {
		t.Fatal("Expected createdAt to be set")
	}
	if len(e.tags) != len(entry.tags) {
		t.Fatalf("Expected tags array len %v received %v", len(entry.tags), len(e.tags))
	}
	for IDx := range entry.tags {
		if e.tags[IDx] != entry.tags[IDx] {
			t.Fatalf("Expected tag name %v received %v", entry.tags[IDx], e.tags[IDx])
		}
	}
}

func TestEntryRemove(t *testing.T) {
	DB, Teardown := Setup(t)
	defer Teardown()

	r := CreateEntryRepository(DB)

	journalID := int64(1)
	entry := &Entry{journalID: journalID, title: "absolutelyRandomTitle", body: "absolutely random body",
		mood: "notsomuch", tags: []string{"randomtag1", "randomtag2"}}

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

func TestKVGet(t *testing.T) {
	DB, Teardown := Setup(t)
	defer Teardown()

	r := CreateKVRepository(DB)

	val, err := r.Get("testName")
	if err != nil {
		t.Fatal(err)
	}
	if val != "testValue" {
		t.Fatalf("Expected value %v received %v", "testValue", val)
	}
}

func TestKVSet(t *testing.T) {
	DB, Teardown := Setup(t)
	defer Teardown()

	r := CreateKVRepository(DB)

	err := r.Set("rnd", "123")
	if err != nil {
		t.Fatal(err)
	}

	val, err := r.Get("rnd")
	if err != nil {
		t.Fatal(err)
	}
	if val != "123" {
		t.Fatalf("Expected value %v received %v", "123", val)
	}
}

func TestKVUpdate(t *testing.T) {
	DB, Teardown := Setup(t)
	defer Teardown()

	r := CreateKVRepository(DB)

	err := r.Set("rnd", "123")
	if err != nil {
		t.Fatal(err)
	}

	err = r.Set("rnd", "321")
	if err != nil {
		t.Fatal(err)
	}

	val, err := r.Get("rnd")
	if err != nil {
		t.Fatal(err)
	}
	if val != "321" {
		t.Fatalf("Expected value %v received %v", "321", val)
	}
}
