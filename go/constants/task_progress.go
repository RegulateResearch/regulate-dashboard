package constants

import (
	"frascati/typing"
)

type TaskProgress typing.Enum

func (tp TaskProgress) ToString() string {
	enum := typing.Enum(tp)
	return enum.Name()
}

var taskProgressCollection typing.EnumCollection = typing.NewEnumCollection(
	"to-do", "blocked", "in-progress", "in-review", "done", "graded",
)

var (
	TASK_PROGRESS_TODO        = TaskProgressFromString("to-do")
	TASK_PROGRESS_BLOCKED     = TaskProgressFromString("blocked")
	TASK_PROGRESS_IN_PROGRESS = TaskProgressFromString("in-progress")
	TASK_PROGRESS_IN_REVIEW   = TaskProgressFromString("in-review")
	TASK_PROGRESS_DONE        = TaskProgressFromString("done")
	TASK_PROGRESS_GRADED      = TaskProgressFromString("graded")
)

func TaskProgressFromVal(taskVal int) TaskProgress {
	enum := taskProgressCollection.GetByVal(taskVal)
	return TaskProgress(enum)
}

func TaskProgressFromString(taskStr string) TaskProgress {
	enum := taskProgressCollection.GetByName(taskStr)
	return TaskProgress(enum)
}
