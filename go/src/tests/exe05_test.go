package src_test

import (
	"strings"
	"testing"

	"example.com/99_problems/src"
)

func TestExe05_EmptyList(t *testing.T) {
	list := []string{}
	c := src.ReverseList(list)
	e := ""

	if strings.Join(c, "") != e {
		t.Logf("%s is diferent from %s", c, e)
		t.Fatal("should return the same list if it is empty")
	}
}

func TestExe05_OneElement(t *testing.T) {
	list := []string{"a"}
	c := src.ReverseList(list)
	e := "a"
	joined := strings.Join(c, "")

	if joined != e {
		t.Logf("%s is diferent from %s", joined, e)
		t.Fatal("should return the same list if it is empty")
	}
}

func TestExe05_MoreElements(t *testing.T) {
	list := []string{"a", "b"}
	c := src.ReverseList(list)
	e := "ba"

	joined := strings.Join(c, "")

	if joined != e {
		t.Logf("%s is diferent from %s", joined, e)
		t.Fatal("should return the list reversed")
	}
}
