package background

import "frascati/exception"

type Task interface {
	Exec() Report
}

type task struct {
	name string
	fun  func() (any, exception.Exception)
}

func newTask(name string, fun func() (any, exception.Exception)) Task {
	return task{
		name: name,
		fun:  fun,
	}
}

func (t task) Exec() Report {
	res, err := t.fun()
	return report{
		name:   t.name,
		result: res,
		err:    err,
	}
}
