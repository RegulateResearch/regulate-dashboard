package main

import "sqlmerger/generator"

func main() {
	filenames := []string{
		"header.sql", "users.sql", "courses.sql", "course_members.sql",
	}

	merger := generator.NewMerger(filenames...)
	merger.Execute()
}
