package tools

func MapItemExists[T comparable, U any](m map[T]U, item T) bool {
	_, ok := m[item]
	return ok
}
