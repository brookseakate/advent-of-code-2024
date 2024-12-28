package util

func CountInSlice[T any](slice []T, f func(T) bool) int {
	count := 0
	for _, v := range slice {
		if f(v) {
			count++
		}
	}
	return count
}
