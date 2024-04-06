package src_test

import (
	"testing"

	"example.com/99_problems/src"
)

func TestExe06_NotPalindromo(t *testing.T) {
	list := []string{"i", "g", "o", "r"}
	c := src.IsPalindromo(list)

	if c {
		t.Fatal("should return false if the word is not palindromo")
	}
}

func TestExe06_IsPalindromoWord(t *testing.T) {
	list := []string{"x", "a", "m", "a", "x"}
	c := src.IsPalindromo(list)
	if !c {
		t.Fatal("should return true if the word is palindromo")
	}

	list_2 := []string{"a", "n", "a"}
	c2 := src.IsPalindromo(list_2)
	if !c2 {
		t.Fatal("should return true if the word is palindromo")
	}
}
