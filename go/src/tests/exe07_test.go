package src_test

import (
	"strings"
	"testing"

	"example.com/99_problems/src"
)

func TestExe07_ShouldFlattenTheList(t *testing.T) {
	list := [][]string{{"a"}, {"b", "c"}}
	c := src.Flatten(list)
	e := []string{"a", "b", "c"}

	if strings.Join(c, "") != strings.Join(e, "") {
		t.Fatal("should flatten the list")
	}
}
