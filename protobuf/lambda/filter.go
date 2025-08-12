package lambda

func CountSatiesfied[T any](data []T, satFn func(T) bool) int {
	count := 0
	for i := 0; i < len(data); i++ {
		if satFn(data[i]) {
			count++
		}
	}

	return count
}

func FilterList[T any](data []T, satFn func(T) bool) []T {
	length := len(data)
	newData := make([]T, length)
	count := 0

	for i := 0; i < length; i++ {
		if satFn(data[i]) {
			newData[count] = data[i]
			count++
		}
	}

	newData = newData[0:count]

	return newData
}
