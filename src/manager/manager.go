package manager

import (
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/usernamenenad/kubelite/task"
)

type Manager struct {
	Pending       queue.Queue
	Tasks         map[string][]*task.Task
	TaskEvents    map[string][]*task.TaskEvent
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {}
func (m *Manager) UpdateTask()   {}
func (m *Manager) SendWork()     {}
