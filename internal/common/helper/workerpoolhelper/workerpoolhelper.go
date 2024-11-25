package workerpoolhelper

import "example/internal/common/helper/loghelper"

type (
	TaskFunc = func(workerId int, data interface{})

	TaskItem struct {
		Task TaskFunc
		Key  string
		Data interface{}
	}
)

type (
	ControlledWorkerPool interface {
		RunAndServe()
		AddTask(taskItem TaskItem)
		GetTotalQueuedTasks() int
		MaskAllTasksSent()
		WaitToDone() <-chan bool
	}

	controlledWorkerPool struct {
		workerNum   int
		workerChan  chan TaskItem
		doneChan    chan bool
		allDoneChan chan bool
	}
)

func NewControlledWorkerPool(poolSize int) ControlledWorkerPool {
	return &controlledWorkerPool{
		workerNum: poolSize,
	}
}

func (c *controlledWorkerPool) initChannels() {
	c.workerChan = make(chan TaskItem, 1)
	c.doneChan = make(chan bool)
	c.allDoneChan = make(chan bool)
}

func (c *controlledWorkerPool) followWorkerDone() {
	for i := 0; i < c.workerNum; i++ {
		<-c.doneChan
	}
	c.allDoneChan <- true
	return
}

func (c *controlledWorkerPool) run() {
	for i := 0; i < c.workerNum; i++ {
		go func(workerId int) {
			for {
				task, notClosed := <-c.workerChan
				if notClosed {
					task.Task(workerId, task.Data)
					loghelper.Logger.Infof("%v - Done for key: %v", workerId, task.Key)
				} else {
					loghelper.Logger.Infof("All thing done: %v", workerId)
					c.doneChan <- true
					return
				}
			}
		}(i)
	}
}

func (c *controlledWorkerPool) RunAndServe() {
	c.initChannels()
	go c.followWorkerDone()
	c.run()
}

func (c *controlledWorkerPool) AddTask(taskItem TaskItem) {
	c.workerChan <- taskItem
}

func (c *controlledWorkerPool) GetTotalQueuedTasks() int {
	return len(c.workerChan)
}

func (c *controlledWorkerPool) MaskAllTasksSent() {
	close(c.workerChan)
}

func (c *controlledWorkerPool) WaitToDone() <-chan bool {
	return c.allDoneChan
}
