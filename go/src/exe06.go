package src

import "strings"

func IsPalindromo(original_list []string) bool {
	reversed := ReverseList(original_list)

	if strings.Join(reversed, "") == strings.Join(original_list, "") {
		return true
	}

	return false
}
