package service

import (
	"errors"
	"fmt"
	"frascati/comp/background"
	"frascati/exception"
	"time"
)

type TryService interface {
	TryBackgroundSuccess(taskName string) string
	TryBackgroundFail(taskName string) string
}

type tryServiceImpl struct {
	backgroundProcessor background.Processor
}

func NewTryService(backgroundProcessor background.Processor) TryService {
	return tryServiceImpl{
		backgroundProcessor: backgroundProcessor,
	}
}

func (s tryServiceImpl) TryBackgroundSuccess(taskName string) string {
	s.backgroundProcessor.AddTask(taskName, func() (any, exception.Exception) {
		time.Sleep(35 * time.Second)
		return fmt.Sprintf("%s: executed successfully", taskName), nil
	})

	return fmt.Sprintf("%s: executed in background", taskName)
}

func (s tryServiceImpl) TryBackgroundFail(taskName string) string {
	s.backgroundProcessor.AddTask(taskName, func() (any, exception.Exception) {
		time.Sleep(35 * time.Second)
		exc := exception.NewBaseException(exception.CAUSE_USER, "try", "background fail", errors.New("error coba2"))
		return fmt.Sprintf("%s: resulted in error: %v", taskName, exc), exc
	})

	return fmt.Sprintf("%s: executed in background", taskName)
}
