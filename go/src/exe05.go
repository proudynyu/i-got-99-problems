package src

func ReverseList(list []string) []string {
	if len(list) <= 1 {
		return list
	}

	var reversed []string

	for i := len(list) - 1; i >= 0; i-- {
		reversed = append(reversed, list[i])
	}

	return reversed
}
