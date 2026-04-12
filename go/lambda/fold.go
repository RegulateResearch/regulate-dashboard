package lambda

func FoldList[Data any, Aggregate any](data []Data, initValue Aggregate, applyFn func(data Data, currentValue Aggregate) Aggregate) Aggregate {
	currentValue := initValue
	for i := 0; i < len(data); i++ {
		currentValue = applyFn(data[i], currentValue)
	}

	return currentValue
}
