package values

func ClassDirsTuple() map[string][]string {
	tuple := map[string][]string{
		"coba": {"go", "grpc-client-coba"},
	}

	return tuple
}

func DirServiceTuple() map[string]string {
	tuple := map[string]string{
		"go":               "frascati",
		"grpc-client-coba": "coba",
	}

	return tuple
}
