package querying

import (
	"frascati/comp/queryexec"
	"frascati/exception"
	"frascati/lambda"
)

func ScanForRows[T any](rows queryexec.Rows, newValueInitFn func() T, scanFn func(rows queryexec.Rows, elem T) (T, exception.Exception)) ([]T, exception.Exception) {
	res := make([]T, 0)
	for rows.Next() {
		empty := newValueInitFn()
		elem, err := scanFn(rows, empty)
		if err != nil {
			return nil, err
		}

		res = append(res, elem)
	}

	return res, nil
}

func ScanForRowsThenTransform[Origin any, Transformed any](
	rows queryexec.Rows,
	newValueInitFn func() Origin,
	scanFn func(rows queryexec.Rows, elem Origin) (Origin, exception.Exception),
	transformFn func(data Origin) Transformed,
) ([]Transformed, exception.Exception) {
	data, err := ScanForRows(rows, newValueInitFn, scanFn)
	if err != nil {
		return nil, err
	}

	return lambda.MapList(data, transformFn), nil
}
