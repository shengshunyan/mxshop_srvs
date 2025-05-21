package slice

// Map 将一个切片中的每个元素转换为另一个类型，并返回新的切片
func Map[T any, R any](slice []T, fn func(T) R) []R {
	result := make([]R, len(slice)) // 预分配容量
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}
