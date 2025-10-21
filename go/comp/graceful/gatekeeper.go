package graceful

import (
	"errors"
	"frascati/exception"
	"sync"
)

type Gatekeeper interface {
	IsOpen() bool
	Process(fun func()) exception.Exception
	Close()
}

type gatekeeper struct {
	isOpen    bool
	waitGroup sync.WaitGroup
}

func NewGateKeeper() Gatekeeper {
	return &gatekeeper{
		isOpen:    true,
		waitGroup: sync.WaitGroup{},
	}
}

func (k *gatekeeper) IsOpen() bool {
	return k.isOpen
}

func (k *gatekeeper) Process(fun func()) exception.Exception {
	if !k.isOpen {
		return exception.NewBaseException(exception.CAUSE_CLOSURE, "gatekeeper", "we are not accepting request at this moment", errors.New("gate is closed for graceful shutdown"))
	}

	k.waitGroup.Add(1)
	defer k.waitGroup.Done()

	fun()

	return nil
}

func (k *gatekeeper) Close() {
	k.isOpen = false
}
