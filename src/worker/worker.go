package worker

import (
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/usernamenenad/kubelite/task"
)

type Worker struct {
	Name      string
	Db        map[uuid.UUID]*task.Task // TODO: real database like etcd
	Queue     queue.Queue
	TaskCount int
}

func (w *Worker) RunTask()   {}
func (w *Worker) StartTask() {}
func (w *Worker) StopTask()  {}
