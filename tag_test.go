package main

import (
	"testing"
)

func TestGetTags(t *testing.T) {
	db, Teardown := Setup(t)
	defer Teardown()

	repo := MakeTagRepo(db)

	entryID := int64(1)
	tags, err := repo.Get(entryID)
	if err != nil {
		t.Fatal(err)
	}

	if tags == nil {
		t.Fatal("Expected tags array to exist")
	}
	if len(tags) == 0 {
		t.Fatal("Expected tags array to have elements")
	}
	for _, tag := range tags {
		if tag.id < 1 {
			t.Fatal("Expected tag id to be covid positive")
		}
		if tag.entryID != entryID {
			t.Fatalf("Expected tag entryID %v received %v", entryID, tag.entryID)
		}
		if tag.name[:len("testTag")] != "testTag" {
			t.Fatalf("Expected tag name %v to start with %v", tag.name, "testTag")
		}
	}
}

func TestCreateTag(t *testing.T) {
	db, Teardown := Setup(t)
	defer Teardown()

	repo := MakeTagRepo(db)

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	entryID := int64(2)
	name := "absolutelyRandom"

	tagID, err := repo.Create(entryID, name, tx)
	if err != nil {
		t.Fatal(err)
	}
	if tagID < 1 {
		t.Fatal("Expected tag id to be positive")
	}

	// TODO maybe dont commit manually in test
	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}

	tags, err := repo.Get(entryID)
	if err != nil {
		t.Fatal(err)
	}

	if tags == nil {
		t.Fatal("Expected tags array to exist")
	}
	if len(tags) == 0 {
		t.Fatal("Expected tags array to have elements")
	}

	tag := getTagByID(tags, tagID)
	if tag == nil {
		t.Fatal("Expected tag to exist")
	}
	if tag.entryID != entryID {
		t.Fatalf("Expected entryID %v received %v", entryID, tag.entryID)
	}
	if tag.name != name {
		t.Fatalf("Expected entryID %v received %v", name, tag.name)
	}
}

func getTagByID(tags []Tag, id int64) *Tag {
	for _, t := range tags {
		if t.id == id {
			return &t
		}
	}

	return nil
}
