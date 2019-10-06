package worker

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rcliao/scheduler"
	"github.com/rcliao/scheduler/processor"
)

/**
 * worker package defines all possible implementation of the scheduler. First
 * implementation is to implement using Go's built-in time.AfterFunc
 */

// NativeScheduler utilizes the Go's built-in method `time.AfterFunc` to
// schedule delayed jobs in future
type NativeScheduler struct {
	timerMap map[string]*time.Timer
}

// NewNativeScheduler is constructor for creating a new NativeScheduler with empty timerMap
func NewNativeScheduler() NativeScheduler {
	return NativeScheduler{
		timerMap: make(map[string]*time.Timer),
	}
}

// DelayJob creates a future event for the job and will be processed at the duration time
func (s *NativeScheduler) DelayJob(duration time.Duration, data string, jobType scheduler.JobType) (scheduler.Job, error) {
	// TODO: create a factory method caller to use different processor func
	timer := time.AfterFunc(duration, processor.Debug(data))
	id := uuid.New().String()
	job := scheduler.Job{
		ID:                 id,
		ScheduledTimestamp: time.Now().Add(duration),
		Status:             scheduler.Scheduled,
		Type:               jobType,
		Data:               data,
	}
	s.timerMap[id] = timer
	return job, nil
}

// Reschedule allows the consumer to update the delay job delayed time
func (s *NativeScheduler) Reschedule(ID string, duration time.Duration) error {
	timer, ok := s.timerMap[ID]
	if !ok {
		return errors.New("Cannot find job with id " + ID)
	}
	// when scheduled job has not been stopped
	if !timer.Stop() {
		<-timer.C
	}
	timer.Reset(duration)
	return nil
}

// Cancel allows the consumer to cancel a job so that it will not happen in the future
func (s *NativeScheduler) Cancel(ID string) error {
	timer, ok := s.timerMap[ID]
	if !ok {
		return errors.New("Cannot find job with id " + ID)
	}
	if !timer.Stop() {
		<-timer.C
	}
	return nil
}
