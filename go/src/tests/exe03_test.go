package src_test

import (
	"testing"

	"example.com/99_problems/src"
)

func TestExe03_OneItemListWithIndexHigher(t *testing.T) {
	list := []string{"a"}
	c := src.GetKFromArrayString(list, 2)
	e := ""

	if c != e {
		t.Fatal("should return empty if the index is higher than the length of the list")
	}
}

func TestExe03_EmptyList(t *testing.T) {
	list := []string{}
	c := src.GetKFromArrayString(list, 0)
	e := ""
	if c != e {
		t.Fatal("should return empty if the list is empty")
	}
}

func TestExe03_IndexIsInThenRangeOfTheList(t *testing.T) {
	list := []string{"a", "b", "c", "d", "e"}
	c := src.GetKFromArrayString(list, 0)
	e := "a"
	c2 := src.GetKFromArrayString(list, 1)
	e2 := "b"
	if c != e {
		t.Fatalf("should return the correct element: %s -> %s", c, e)
	}
	if c2 != e2 {
		t.Fatalf("should return the correct element: %s -> %s", c2, e2)
	}
}
