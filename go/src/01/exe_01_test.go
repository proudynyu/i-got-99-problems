package src

import (
	"testing"
)

func TestExe01_EmptyList(t *testing.T) {
	list := []string{}
	expected := ""

	test := GetLastFromArrayString(list)

	if test != expected {
		t.Fatalf("%s is not equal to %s", test, expected)
	}
}

func TestExe01_ListWithOneElement(t *testing.T) {
	list := []string{"a", "b", "c", "d"}
	expected := "d"
	test := GetLastFromArrayString(list)

	if test != expected {
		t.Fatalf("%s is not equal to %s", test, expected)
	}
}

func TestExe01_ListWithMoreElements(t *testing.T) {
	list := []string{"a"}
	expected := "a"
	test := GetLastFromArrayString(list)

	if test != expected {
		t.Fatalf("%s is not equal to %s", test, expected)
	}
}
