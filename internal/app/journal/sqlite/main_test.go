package sqlite

import (
	"testing"

	"github.com/rolandmarg/jou/internal/pkg/fixture"
)

func TestGet(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	name := "test"
	j, err := r.Get(name)
	if err != nil {
		t.Fatal(err)
	}
	if j == nil {
		t.Fatal("Expected to exist")
	}
	if j.ID < 1 {
		t.Fatal("Expected ID to be positive")
	}
	if j.Name[:len("test")] != "test" {
		t.Fatalf("Expected name %v received %v", "test", j.Name)
	}
}

func TestSetDefault(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	err := r.SetDefault("test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetDefault(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	err := r.SetDefault("test")
	if err != nil {
		t.Fatal(err)
	}

	j, err := r.GetDefault()
	if err != nil {
		t.Fatal(err)
	}
	if j.Name != "test" {
		t.Fatalf("Expected default journal %v received %v", "test", j.Name)
	}
}

func TestUpdate(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	_, err := r.Create("xutu")
	if err != nil {
		t.Fatal(err)
	}

	err = r.Update("xutu", "butu")
	if err != nil {
		t.Fatal(err)
	}

	j, err := r.Get("butu")
	if err != nil {
		t.Fatal(err)
	}
	j, err = r.Get("xutu")
	if err != nil {
		t.Fatal(err)
	}
	if j != nil {
		t.Fatal("Expected journal xutu not to exist")
	}
}
func TestCreate(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	_, err := r.Create("testJ")
	if err != nil {
		t.Fatal(err)
	}

	j, err := r.Get("testJ")
	if err != nil {
		t.Fatal(err)
	}

	if j == nil {
		t.Fatal("Expected to exist")
	}
	if j.Name != "testJ" {
		t.Fatalf("Expected name %v received %v", "testJ", j.Name)
	}
}

func TestCreateDuplicate(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	_, err := r.Create("dup")
	if err != nil {
		t.Fatal(err)
	}

	_, err = r.Create("dup")
	if err == nil {
		t.Fatal("Expected journal create to fail on duplicate name")
	}
}

func TestRemove(t *testing.T) {
	DB, Teardown := fixture.Setup(t)
	defer Teardown()

	r := MakeRepository(DB)

	_, err := r.Create("testJ")
	if err != nil {
		t.Fatal(err)
	}

	_, err = r.Get("testJ")
	if err != nil {
		t.Fatal(err)
	}

	err = r.Remove("testJ")
	if err != nil {
		t.Fatal(err)
	}

	j, err := r.Get("testJ")
	if err != nil {
		t.Fatal(err)
	}
	if j != nil {
		t.Fatal("Expected journal testJ not to exist")
	}
}
