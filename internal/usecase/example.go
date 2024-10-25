package usecase

import (
	"example/internal/common/helper/loghelper"
	"sync"
	"time"
)

type (
	ExampleUseCase interface {
		GoroutineTest() (any, error)
		CronScheduler()
		MutexTest()
	}

	exampleUseCase struct {
		Mutex *sync.Mutex
	}
)

func NewExampleUseCase(mutex *sync.Mutex) ExampleUseCase {
	return &exampleUseCase{
		Mutex: mutex,
	}
}

var count = 0

func (u *exampleUseCase) GoroutineTest() (any, error) {
	return "Ok", nil
}

func (u *exampleUseCase) MutexTest() {
	u.Mutex.Lock()
	defer u.Mutex.Unlock()
	loghelper.Logger.Info("Lock")
	count = count + 1
	loghelper.Logger.Info(count)
	time.Sleep(5 * time.Second)
	loghelper.Logger.Info("Unlock")
}

func (u *exampleUseCase) CronScheduler() {
	loghelper.Logger.Info("CronScheduler")
}
