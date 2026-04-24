package querying

import (
	"frascati/lambda"
	"strings"
)

func BulkValues[T any](values []T, idxStart int, genFn func(value T, currentIdx int) DataStrArgs) (rowstr string, args []any, idxNext int) {
	datas, lastIdx := lambda.MapListWithDependentSideEffectState(
		values, idxStart,
		func(val T, currentIdx int) (DataStrArgs, int) {
			res := genFn(val, currentIdx)
			return res, len(res.Args)
		}, func(currentIdx int, paramNum int) (nextIdx int) {
			nextIdx = currentIdx + paramNum
			return nextIdx
		},
	)

	const valueDelim = ",\n"
	rowstr = strings.Join(lambda.MapList(datas, func(data DataStrArgs) string {
		return data.RowStr
	}), valueDelim)

	args = make([]any, 0)
	args = lambda.FoldList(datas, args, func(data DataStrArgs, currentArgs []any) []any {
		return append(currentArgs, data.Args...)
	})

	idxNext = lastIdx
	return rowstr, args, idxNext
}

func BulkValuesFromBeginning[T any](values []T, genFn func(value T, currentIdx int) DataStrArgs) (rowstr string, args []any, idxNext int) {
	paramIdxStart := 1
	rowstr, args, idxNext = BulkValues(values, paramIdxStart, genFn)
	return rowstr, args, idxNext
}
