package task

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type TaskState int

const (
	Pending = iota
	Scheduled
	Running
	Completed
	Failed
)

type TaskEvent struct {
	Id        uuid.UUID
	State     TaskState
	Timestamp time.Time
	Task      Task
}

type Task struct {
	Id            uuid.UUID
	Name          string
	State         TaskState
	Image         string
	Memory        int
	Disk          int
	ExposedPorts  nat.PortSet
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
}
