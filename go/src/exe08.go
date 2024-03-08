package src

func VerifyConsecutivesDuplicates(list []string) []string {
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
