package src

// // just one level
func Flatten(list [][]string) []string {
	var l []string

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i]); j++ {
			l = append(l, list[i][j])
		}
	}

	return l
}
