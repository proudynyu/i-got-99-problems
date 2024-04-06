package src

func PackConsecutivesDuplicates(list []string) [][]string {
	var l [][]string
	var s []string

	for i := 0; i <= len(list); i++ {
		if i == len(list)-1 {
			s = append(s, list[i])
			break
		}
		if list[i+1] == list[i] {
			s = append(s, list[i])
			continue
		}
		l = append(l, s)
		s = []string{}
	}
	return l
}
