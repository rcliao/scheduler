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

// JobStatus enum defines all possible job status values
type JobStatus int

const (
	// Scheduled meaning that the job status is pending to be completed in the future
	Scheduled JobStatus = iota
	// Cancelled means the job is cancelled and will not complete any time in future
	Cancelled
	// Completed means that the scheduler confirms the job is completed
	Completed
)

// Job is the base entity to store the scheduling meta data as well as its work
type Job struct {
	ID                 string
	ScheduledTimestamp time.Time
	Status             JobStatus
	Type               JobType
	Data               string
}

// JobRepository manages the database interaction to store and retrieve job data
type JobRepository interface {
	Create(job Job) (Job, error)
	Get(ID string) (Job, error)
	Put(ID string, job Job) (Job, error)
}

// Scheduler is the top level API to schedule a delay job
type Scheduler interface {
	DelayJob(duration time.Duration, data string, jobType JobType) (Job, error)
	Reschedule(ID string, duration time.Duration) error
	Cancel(ID string) error
}
