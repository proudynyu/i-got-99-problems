package src_test

import (
	"strings"
	"testing"

	"example.com/99_problems/src"
)

func TestExe08_ListWithDuplicates(t *testing.T) {
	list := []string{"a", "a", "a", "a", "b", "c", "c", "a", "a", "d", "e", "e", "e", "e"}
	c := src.VerifyConsecutivesDuplicates(list)
	e := []string{"a", "b", "c", "a", "d", "e"}

	if strings.Join(c, "") != strings.Join(e, "") {
		t.Fatal("should return the list unduplicated")
	}
}

func TestExe08_ListWithoutDuplicates(t *testing.T) {
	list := []string{"a", "b", "c", "a", "d", "e"}
	c := src.VerifyConsecutivesDuplicates(list)
	e := []string{"a", "b", "c", "a", "d", "e"}

	if strings.Join(c, "") != strings.Join(e, "") {
		t.Fatal("should return the list passed with no duplications was in the list")
	}

}
