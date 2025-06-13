package values

func PrintDelim(chara byte, iter int) string {
	res := make([]byte, iter)
	for i := 0; i < iter; i++ {
		res[i] = chara
	}

	return string(res)
}
