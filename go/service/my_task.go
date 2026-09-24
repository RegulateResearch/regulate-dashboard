package service

import (
	"errors"
	"frascati/comp/txhandler"
	"frascati/constants"
	"frascati/exception"
	"frascati/obj/entity"
	"frascati/repository"
	"frascati/typing"
)

type myTaskService interface {
	MyTasks(ctx typing.Context, userID typing.ID, syncFirst bool) ([]entity.UserTask, exception.Exception)
	MyTaskById(ctx typing.Context, userID typing.ID, taskID typing.ID) (entity.UserTask, exception.Exception)
	UpdateTask(ctx typing.Context, userID typing.ID, updateData entity.UserTask) (entity.UserTask, exception.Exception)
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

func (s myTaskServiceImpl) MyTaskById(ctx typing.Context, userID typing.ID, taskID typing.ID) (entity.UserTask, exception.Exception) {
	err := s.checkUserExistByID(ctx, userID)
	if err != nil {
		return entity.UserTask{}, err
	}

	user := entity.UserWithID(userID)
	task, err := s.taskRepo.FindByUserAndId(ctx, user, taskID)
	return task, err
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

func (s myTaskServiceImpl) UpdateTask(ctx typing.Context, userID typing.ID, updateData entity.UserTask) (entity.UserTask, exception.Exception) {
	var res entity.UserTask
	err := s.transactor.WithTransaction(ctx, txhandler.TxOptionReadCommitted, false, func(ctx typing.Context) exception.Exception {
		task, err := s.updateTask(ctx, userID, updateData)
		res = task
		return err
	})

	return res, err
}

func (s myTaskServiceImpl) updateTask(ctx typing.Context, userID typing.ID, updateData entity.UserTask) (entity.UserTask, exception.Exception) {
	user := entity.UserWithID(userID)
	oldData, err := s.taskRepo.FindByUserAndId(ctx, user, updateData.ID)
	if err != nil {
		return entity.UserTask{}, err
	}

	if oldData.Progress == constants.TASK_PROGRESS_GRADED {
		newErr := errors.New("cannot change task status: task has been graded")
		exc := exception.NewBaseException(exception.CAUSE_USER, "my_task/service", newErr.Error(), newErr)
		return entity.UserTask{}, exc
	}

	updateData.IsStartFlag =
		(oldData.Progress == constants.TASK_PROGRESS_TODO || oldData.Progress == constants.TASK_PROGRESS_BLOCKED) &&
			(updateData.Progress != constants.TASK_PROGRESS_TODO && updateData.Progress != constants.TASK_PROGRESS_BLOCKED)

	updateData.IsDoneFlag = updateData.Progress == constants.TASK_PROGRESS_DONE

	res, err := s.taskRepo.UpdateStatusAndTimes(ctx, updateData)
	return res, err
}
