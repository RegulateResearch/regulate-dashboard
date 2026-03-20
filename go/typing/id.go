package typing

import "strconv"

type ID int64

func (id ID) String() string {
	return strconv.FormatInt(int64(id), 10)
}

func IDFromString(idstr string) ID {
	var id int64 = -1
	idnum, err := strconv.ParseInt(idstr, 10, 64)
	if err == nil {
		id = idnum
	}

	return ID(id)
}
