package lambda

func MapList[Origin any, Transformed any](data []Origin, mapper func(Origin) Transformed) []Transformed {
	length := len(data)
	transformedData := make([]Transformed, len(data))
	for i := 0; i < length; i++ {
		transformedData[i] = mapper(data[i])
	}

	return transformedData
}
