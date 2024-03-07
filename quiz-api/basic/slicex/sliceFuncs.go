package slicex

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

func Map[TIn any, TOut any](slice []TIn, mapper func(TIn) TOut) []TOut {
	res := make([]TOut, len(slice))
	for i, item := range slice {
		res[i] = mapper(item)
	}
	return res
}
