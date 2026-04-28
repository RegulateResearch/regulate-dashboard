package main

import "sqlmerger/generator"

func main() {
	filenames := []string{
		"header.sql", "users.sql", "records.sql", "samples.sql",
	}

	merger := generator.NewMerger(filenames...)
	merger.Execute()
}
