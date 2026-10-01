package querying

import "fmt"

type Pair struct {
	Col     string
	Val     any
	IsParam bool
}

func (p Pair) updateQueryStrWithArg(currentIdx int) (res string, arg any, nextIdx int) {
	if p.IsParam {
		return fmt.Sprintf("%s = $%d", p.Col, currentIdx), p.Val, currentIdx + 1
	}

	return fmt.Sprintf("%s = %v", p.Col, p.Val), nil, currentIdx
}
