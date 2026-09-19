package typing

import "frascati/lambda"

type Enum struct {
	val  int
	name string
}

func (e Enum) Val() int {
	return e.val
}

func (e Enum) Name() string {
	return e.name
}

const enumInitIdx int = 1

type EnumCollection struct {
	enums []Enum
}

func NewEnumCollection(names ...string) EnumCollection {
	padding := []Enum{{}}
	enums := lambda.MapListWithSerialState(
		names, enumInitIdx,
		func(name string, currentIdx int) Enum {
			return Enum{name: name, val: currentIdx}
		},
		func(currentState int) (nextState int) {
			nextState = currentState + 1
			return nextState
		},
	)

	res := append(padding, enums...)

	return EnumCollection{enums: res}
}

func (c EnumCollection) GetByVal(val int) Enum {
	res := c.enums[0]
	size := len(c.enums)
	if val >= 1 && val < size {
		res = c.enums[val]
	}

	return res
}

func (c EnumCollection) GetByName(name string) Enum {
	res := c.enums[0]
	found := false
	for i := 1; !found && i < len(c.enums); i++ {
		if c.enums[i].Name() == name {
			res = c.enums[i]
			found = true
		}
	}

	return res
}
