package src_test

import (
	"testing"

	"example.com/99_problems/src"
)

func TestExe01_EmptyList(t *testing.T) {
	list := []string{}
	expected := ""

	test := src.GetLastFromArrayString(list)

	if test != expected {
		t.Fatalf("%s is not equal to %s", test, expected)
	}
}

func TestExe01_ListWithOneElement(t *testing.T) {
	list := []string{"a", "b", "c", "d"}
	expected := "d"
	test := src.GetLastFromArrayString(list)

	if test != expected {
		t.Fatalf("%s is not equal to %s", test, expected)
	}
}

func TestExe01_ListWithMoreElements(t *testing.T) {
	list := []string{"a"}
	expected := "a"
	test := src.GetLastFromArrayString(list)

	if test != expected {
		t.Fatalf("%s is not equal to %s", test, expected)
	}
}
