package repository

import "github.com/rcliao/scheduler"

// JobMemoryRepository handles the job data in memory
type JobMemoryRepository struct {
	data map[string]scheduler.Job
}

// NewJobMemoryRepository creates a new JobMemoryRepository
func NewJobMemoryRepository() JobMemoryRepository {
	return JobMemoryRepository{
		data: make(map[string]scheduler.Job),
	}
}

// GetJob retrieves a job by id
func (j JobMemoryRepository) GetJob(ID string) (scheduler.Job, error) {
	return j.data[ID], nil
}

// PutJob updates a job by id
func (j JobMemoryRepository) PutJob(ID string, job scheduler.Job) (scheduler.Job, error) {
	job.ID = ID
	j.data[ID] = job
	return job, nil
}
