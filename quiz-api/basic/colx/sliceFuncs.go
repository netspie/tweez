package colx

func HasEmptyStrings(arr []string) bool {
	for _, item := range arr {
		if len(item) == 0 {
			return true
		}
	}

	return false
}

func IsInRange[T any](arr []T, idx int) bool {
	if idx < 0 || idx > len(arr)-1 {
		return false
	}

	return true
}
