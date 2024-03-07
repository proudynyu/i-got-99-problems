package main

import (
	"fmt"
	"strings"
)

/*********** Ex 01 **********/
func getLastFromArrayString(arr []string) string {
	if len(arr) == 0 {
		return ""
	}

	return arr[len(arr)-1]
}

func Ex01() {
	test_1 := []string{}
	test_2 := []string{"a", "b", "c", "d"}
	test_3 := []string{"a"}

	c := getLastFromArrayString(test_1)
	s := getLastFromArrayString(test_2)
	t := getLastFromArrayString(test_3)

	fmt.Println(c)
	fmt.Println(s)
	fmt.Println(t)
}

/*********** Ex 01 **********/

/*********** Ex 02 **********/
func getLastButOneFromStringArray(arr []string) (string, string) {
	if len(arr) == 0 {
		return "", ""
	}

	if len(arr) == 1 {
		return "", ""
	}

	return arr[len(arr)-2], arr[len(arr)-1]
}

func Ex02() {
	test_1 := []string{}
	test_2 := []string{"a", "b", "c", "d"}
	test_3 := []string{"a"}

	c1, c2 := getLastButOneFromStringArray(test_1)
	s1, s2 := getLastButOneFromStringArray(test_2)
	t1, t2 := getLastButOneFromStringArray(test_3)

	fmt.Println(c1, c2)
	fmt.Println(s1, s2)
	fmt.Println(t1, t2)
}

/*********** Ex 02 **********/

/*********** Ex 03 **********/
func getKFromArrayString(arr []string, element int) string {
	if len(arr) < element || len(arr) == 0 {
		return ""
	}

	return arr[element-1]
}

func Ex03() {
	arr := []string{"a", "b", "c", "d", "e"}
	fail := []string{"a"}

	c := getKFromArrayString(arr, 3)
	s := getKFromArrayString(fail, 3)

	fmt.Println(c)
	fmt.Println(s)
}

/*********** Ex 03 **********/

/*********** Ex 04 **********/
// ideal would be to recreate the idea of len()
func Ex04() {
	list := []string{"a", "b", "c"}
	list_2 := []string{}

	fmt.Println(len(list), cap(list), list)
	fmt.Println(len(list_2), cap(list_2), list_2)
}

/*********** Ex 04 **********/

/*********** Ex 05 **********/
func ReverseList(list []string) []string {
	var reversed []string

	for i := len(list) - 1; i >= 0; i-- {
		reversed = append(reversed, list[i])
	}

	return reversed
}

func Ex05() {
	list := []string{"a", "b", "c"}
	reversed_list := ReverseList(list)

	fmt.Println(list)
	fmt.Println(reversed_list)
}

/*********** Ex 05 **********/

/*********** Ex 06 **********/
func isPalindromo(original_list []string) bool {
	reversed := ReverseList(original_list)

	if strings.Join(reversed, "") == strings.Join(original_list, "") {
		return true
	}

	return false
}

func Ex06() {
	list := []string{"x", "a", "m", "a", "x"}
	list_2 := []string{"a", "n", "a"}
	list_3 := []string{"i", "g", "o", "r"}

	fmt.Println(list, " is palindromo: ", isPalindromo(list))
	fmt.Println(list_2, " is palindromo: ", isPalindromo(list_2))
	fmt.Println(list_3, " is palindromo: ", isPalindromo(list_3))
}

/*********** Ex 06 **********/

/*********** Ex 07 **********/
// just one level
func flatten(list [][]string) []string {
	var l []string

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i]); j++ {
			l = append(l, list[i][j])
		}
	}

	return l
}

func Ex07() {
	list := [][]string{{"a"}, {"b", "c"}}
	flatted := flatten(list)

	fmt.Println(list, " flatted: ", flatted)
}

/*********** Ex 07 **********/

/*********** Ex 07 **********/
func verifyConsecutivesDuplicates(list []string) []string {
	var l []string

	var s string
	for i := 0; i <= len(list); i++ {
		if i == len(list)-1 {
			l = append(l, list[i])
			break
		}

		s = list[i]

		if list[i+1] != s {
			l = append(l, s)
		}
	}

	return l
}
func Ex08() {
	list := []string{"a", "a", "a", "a", "b", "c", "c", "a", "a", "d", "e", "e", "e", "e"}

	fmt.Println(list, " un-duplicated: ", verifyConsecutivesDuplicates(list))
}

/*********** Ex 07 **********/
func main() {
	Ex08()
}
