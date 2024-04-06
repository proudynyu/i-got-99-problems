package src_test

import (
	"testing"

	"example.com/99_problems/src"
)

func TestEx02_EmptyList(t *testing.T) {
	test_1 := []string{}
	c1, c2 := src.GetLastButOneFromStringArray(test_1)
	e1, e2 := "", ""

	if c1 != e1 {
		t.Fatalf("%s is not equal %s", c1, e1)
	}

	if c2 != e2 {
		t.Fatalf("%s is not equal %s", c2, e2)
	}
}

func TestEx02_OneItemList(t *testing.T) {
	test_2 := []string{"a", "b", "c", "d"}
	c1, c2 := src.GetLastButOneFromStringArray(test_2)
	e1, e2 := "c", "d"

	if c1 != e1 {
		t.Fatalf("%s is not equal %s", c1, e1)
	}
	if c2 != e2 {
		t.Fatalf("%s is not equal %s", c2, e2)
	}
}

func TestEx02_MoreItemsInList(t *testing.T) {
	test_3 := []string{"a"}
	c1, c2 := src.GetLastButOneFromStringArray(test_3)
	e1, e2 := "", ""

	if c1 != e1 {
		t.Fatalf("%s is not equal %s", c1, e1)
	}
	if c2 != e2 {
		t.Fatalf("%s is not equal %s", c2, e2)
	}
}
