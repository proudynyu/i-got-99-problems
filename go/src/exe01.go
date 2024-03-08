package src

func GetLastFromArrayString(arr []string) string {
	if len(arr) == 0 {
		return ""
	}

	return arr[len(arr)-1]
}
