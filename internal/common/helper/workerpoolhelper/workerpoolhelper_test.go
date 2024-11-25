package workerpoolhelper

import (
	"example/internal/common/helper/loghelper"
	"fmt"
	"testing"
	"time"
)

type (
	workerService struct {
		workerPoolHelper ControlledWorkerPool
	}
)

func NewWorkerService() *workerService {
	workerPoolHelper := NewControlledWorkerPool(10)
	return &workerService{
		workerPoolHelper: workerPoolHelper,
	}
}

func (w *workerService) handleTask(task string) {
	w.workerPoolHelper.AddTask(TaskItem{
		Task: func(workerId int, data interface{}) {
			loghelper.Logger.Infof("Handle data of %s", data)
		},
		Key:  fmt.Sprintf(task),
		Data: task,
	})
}

func (w *workerService) handleDone() {
	w.workerPoolHelper.MaskAllTasksSent()
}

func runInWorker(handleTask func(task string), handleDone func()) {
	loghelper.Logger.Info("Starting consume task ...")

	for i := 0; i < 10; i++ {
		handleTask(fmt.Sprintf("task %d", i))
	}

	handleDone()
}

func (w *workerService) startWorkerPoolFlow() string {
	timeStart := time.Now()

	w.workerPoolHelper.RunAndServe()
	runInWorker(w.handleTask, w.handleDone)
	<-w.workerPoolHelper.WaitToDone()

	loghelper.Logger.Infof("All thing done")
	executionTime := time.Since(timeStart)
	loghelper.Logger.Infof("Execution time: %v", executionTime)
	return "Done"
}

func TestControlledWorkerPoolSuccess(t *testing.T) {
	_ = loghelper.InitZap("testing", "debug")
	workerService := NewWorkerService()
	done := workerService.startWorkerPoolFlow()
	if done != "Done" {
		t.Error("workerpool cannot done")
	}
}
