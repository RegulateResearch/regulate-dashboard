package typing

import "strconv"

type ID int64

func (id ID) String() string {
	return strconv.FormatInt(int64(id), 10)
}
