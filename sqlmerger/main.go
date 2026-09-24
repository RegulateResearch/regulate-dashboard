package main

import "sqlmerger/generator"

func main() {
	filenames := []string{
		"header.sql", "users.sql",
		"courses.sql", "course_members.sql", "course_items.sql",
		"user_tasks.sql", "task_record.sql",
	}

	merger := generator.NewMerger(filenames...)
	merger.Execute()
}
