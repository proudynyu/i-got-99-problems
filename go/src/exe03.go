package src

func GetKFromArrayString(arr []string, element int) string {
	if len(arr) < element || len(arr) == 0 || element < 0 {
		return ""
	}

	return arr[element]
}
