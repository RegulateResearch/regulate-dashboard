package querying

import (
	"strings"
)

func UpdatePairsQuerystr[T any](data T, startIdx int, pairGenFn func(elem T) []Pair) (res string, args []any, nextIdx int) {
	pairs := pairGenFn(data)
	length := len(pairs)
	pairStrings := make([]string, length)
	args = make([]any, 0)
	currentIdx := startIdx

	for i := 0; i < length; i++ {
		pair := pairs[i]
		token, arg, nextIdx := pair.updateQueryStrWithArg(currentIdx)
		currentIdx = nextIdx
		pairStrings[i] = token

		if arg != nil {
			args = append(args, arg)
		}
	}

	nextIdx = currentIdx

	return strings.Join(pairStrings, ",\n"), args, nextIdx
}
