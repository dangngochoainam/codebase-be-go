package usecase

import (
	"log"
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
	log.Println("Lock")
	count = count + 1
	log.Println(count)
	time.Sleep(5 * time.Second)
	log.Println("Unlock")
}

func (u *exampleUseCase) CronScheduler() {
	log.Println("CronScheduler")
}
