package main

import "sqlmerger/generator"

func main() {
	filenames := []string{
		"header.sql", "users.sql", "courses.sql", "course_members.sql", "course_items.sql",
	}

	merger := generator.NewMerger(filenames...)
	merger.Execute()
}
