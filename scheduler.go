package scheduler

import "time"

// JobType defines different kinds of job to open up different ways of processing
// the same job
type JobType int

const (
	// DebugFunc simply logs the message in console for debugging purpose
	DebugFunc JobType = iota
	// WebHook is the most basic job to send the job data to a HTTP endpoint
	WebHook
)

// Job is the base entity to store the scheduling meta data as well as its work
type Job struct {
	ID                 string
	ScheduledTimestamp time.Time
	Status             string
	Type               JobType
	Data               string
}

// JobRepository manages the database interaction to store and retrieve job data
type JobRepository interface {
	GetJob(ID string) (Job, error)
	PutJob(ID string, job Job) (Job, error)
}

// Scheduler is the top level API to schedule a delay job
type Scheduler interface {
	DelayJob(duration time.Duration, jobType JobType, data string) (Job, error)
	Reschedule(ID string, duration time.Duration) (Job, error)
	Cancel(ID string) error
}
