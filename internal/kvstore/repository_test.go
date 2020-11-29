package kvstore

import (
	"testing"

	"github.com/rolandmarg/jou/internal/pkg/fixture"
)

func TestGet(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	val, err := r.Get("testName")
	if err != nil {
		t.Fatal(err)
	}
	if val != "testValue" {
		t.Fatalf("Expected value %v received %v", "testValue", val)
	}
}

func TestSet(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

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

func TestUpdate(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

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
