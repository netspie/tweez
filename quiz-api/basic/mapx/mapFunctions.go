package mapx

func Map[TKey comparable, TValue, TOut any](m map[TKey]TValue, mapper func(TKey, TValue) TOut) []TOut {
	res := make([]TOut, len(m))
	i := 0
	for key, value := range m {
		res[i] = mapper(key, value)
		i++
	}
	return res
}
