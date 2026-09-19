package service

import (
	"frascati/comp/txhandler"
	"frascati/exception"
	"frascati/obj/entity"
	"frascati/repository"
	"frascati/typing"
)

type myTaskService interface {
	MyTasks(ctx typing.Context, userID typing.ID, syncFirst bool) ([]entity.UserTask, exception.Exception)
}

type myTaskServiceImpl struct {
	myBaseService
	taskRepo   repository.UserTaskRepository
	recordRepo repository.TaskRecordRepository
	transactor txhandler.Transactor
}

func newMyTaskService(
	myBaseServ myBaseService,
	taskRepo repository.UserTaskRepository,
	recordRepo repository.TaskRecordRepository,
	transactor txhandler.Transactor,
) myTaskService {
	return myTaskServiceImpl{
		myBaseService: myBaseServ,
		taskRepo:      taskRepo,
		recordRepo:    recordRepo,
		transactor:    transactor,
	}
}

func (s myTaskServiceImpl) MyTasks(ctx typing.Context, userID typing.ID, syncFirst bool) ([]entity.UserTask, exception.Exception) {
	var res []entity.UserTask
	err := s.transactor.WithTransaction(ctx, txhandler.TxOptionReadCommitted, false, func(ctx typing.Context) exception.Exception {
		tasks, err := s.getAll(ctx, userID, syncFirst)
		res = tasks
		return err
	})

	return res, err
}

func (s myTaskServiceImpl) getAll(ctx typing.Context, userID typing.ID, syncFirst bool) ([]entity.UserTask, exception.Exception) {
	err := s.checkUserExistByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	user := entity.UserWithID(userID)
	if syncFirst {
		_, err := s.taskRepo.AddInferByNotYetAdded(ctx, user)
		if err != nil {
			return nil, err
		}
	}

	tasks, err := s.taskRepo.FindByUser(ctx, user)
	return tasks, err
}
