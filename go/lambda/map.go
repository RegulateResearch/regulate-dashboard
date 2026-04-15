package lambda

func MapList[Origin any, Transformed any](data []Origin, mapper func(Origin) Transformed) []Transformed {
	length := len(data)
	transformedData := make([]Transformed, len(data))
	for i := 0; i < length; i++ {
		transformedData[i] = mapper(data[i])
	}

	return transformedData
}

func MapListWithDependentSideEffectState[Origin any, Transformed any, SideEffectState any](
	data []Origin,
	initState SideEffectState,
	mapperSideEffectFn func(Origin, SideEffectState) (Transformed, SideEffectState),
	sideEffectApplier func(current SideEffectState, transformResidue SideEffectState) SideEffectState,
) (result []Transformed, lastState SideEffectState) {
	length := len(data)
	transformedData := make([]Transformed, len(data))
	currentState := initState
	for i := 0; i < length; i++ {
		var residue SideEffectState
		transformedData[i], residue = mapperSideEffectFn(data[i], currentState)
		currentState = sideEffectApplier(currentState, residue)
	}

	result = transformedData
	lastState = currentState
	return result, lastState
}
