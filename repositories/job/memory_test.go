package repository

import (
	"testing"

	"github.com/rcliao/scheduler"
)

func TestPutJob(t *testing.T) {
	repository := NewJobMemoryRepository()
	job, _ := repository.PutJob("test", scheduler.Job{})
	if job.ID != "test" {
		t.Errorf("Failed to upsert a new job with with custom id")
	}
}

func TestGetJob(t *testing.T) {
	repository := NewJobMemoryRepository()
	repository.PutJob("test", scheduler.Job{
		Data: "test-data",
	})
	repository.GetJob("test")
	if job.Data != "test-data" {
		t.Errorf("Failed to insert and get the job with same id")
	}
}
