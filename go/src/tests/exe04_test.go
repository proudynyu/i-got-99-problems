package src_test

import (
	"testing"

	"example.com/99_problems/src"
)

func TestExe04_EmptyList(t *testing.T) {
	list := []string{}
	c := src.GetTheLength(list)
	e := 0

	if c != e {
		t.Fatal("should be zero if list is empty")
	}
}

func TestExe04_HasItemsInTheList(t *testing.T) {
	list := []string{"a", "b"}
	c := src.GetTheLength(list)
	e := 2

	if c != e {
		t.Fatal("should return the length of the list")
	}
}
