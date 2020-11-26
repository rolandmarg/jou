package main

import (
	"testing"
)

func TestGetEnv(t *testing.T) {
	db, Teardown := Setup(t)
	defer Teardown()

	repo := MakeEnvRepo(db)

	val, err := repo.Get("testName")
	if err != nil {
		t.Fatal(err)
	}
	if val != "testValue" {
		t.Fatalf("Expected value %v received %v", "testValue", val)
	}
}

func TestSetEnv(t *testing.T) {
	db, Teardown := Setup(t)
	defer Teardown()

	repo := MakeEnvRepo(db)

	err := repo.Set("rnd", "123")
	if err != nil {
		t.Fatal(err)
	}

	val, err := repo.Get("rnd")
	if err != nil {
		t.Fatal(err)
	}
	if val != "123" {
		t.Fatalf("Expected value %v received %v", "123", val)
	}
}

func TestUpdateEnv(t *testing.T) {
	db, Teardown := Setup(t)
	defer Teardown()

	repo := MakeEnvRepo(db)

	err := repo.Set("rnd", "123")
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Set("rnd", "321")
	if err != nil {
		t.Fatal(err)
	}

	val, err := repo.Get("rnd")
	if err != nil {
		t.Fatal(err)
	}
	if val != "321" {
		t.Fatalf("Expected value %v received %v", "321", val)
	}
}
