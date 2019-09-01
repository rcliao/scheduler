package scheduler

import "time"

type WorkFn func() error

type Job struct {
	ID                 string
	ScheduledTimestamp time.Time
	Status             string
	Work               WorkFn
}

type Scheduler interface {
	DelayJob(milliSeconds int, job WorkFn) (Job, error)
	Reschedule(ID string, newDelayedDuration time.Time) (Job, error)
	Cancel(ID string) error
}
