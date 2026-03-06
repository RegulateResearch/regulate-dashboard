package background

import (
	"frascati/comp/logging"
	"frascati/exception"
	"sync"
	"time"
)

type Processor interface {
	IsOpen() bool
	Open()
	Close()
	Wait()
	AddTask(taskName string, taskFunction func() (any, exception.Exception))
}

type processor struct {
	isOpen    bool
	waitGroup sync.WaitGroup
	logger    logging.ExceptionSupportLogger
}

func NewProcessor(logger logging.ExceptionSupportLogger) Processor {
	return &processor{
		isOpen:    false,
		waitGroup: sync.WaitGroup{},
		logger:    logger,
	}
}

func (p *processor) IsOpen() bool {
	return p.isOpen
}

func (p *processor) Open() {
	p.isOpen = true
}

func (p *processor) Close() {
	p.isOpen = false
}

func (p *processor) Wait() {
	p.waitGroup.Wait()
}

func (p *processor) AddTask(taskName string, taskFunction func() (any, exception.Exception)) {
	task := newTask(taskName, taskFunction)

	p.waitGroup.Add(1)
	go func() {
		defer p.waitGroup.Done()

		startTime := time.Now()
		taskReport := task.Exec()
		endTime := time.Now()

		elapsedTime := endTime.Sub(startTime)
		fields := map[string]any{
			"method":    "background task",
			"latency":   elapsedTime.String(),
			"task name": taskReport.Name(),
			"result":    taskReport.Result(),
		}

		logger := p.logger.WithFields(fields)
		err := taskReport.Err()

		logger.LogException(err)
	}()
}
