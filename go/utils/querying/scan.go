package querying

import (
	"frascati/comp/queryexec"
	"frascati/exception"
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
