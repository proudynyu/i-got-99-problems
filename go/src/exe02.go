package src

func GetLastButOneFromStringArray(arr []string) (string, string) {
	if len(arr) == 0 {
		return "", ""
	}

	if len(arr) == 1 {
		return "", ""
	}

	return arr[len(arr)-2], arr[len(arr)-1]
}
