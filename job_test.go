package main

import (
	"testing"
	"time"
)

func TestJob_SetStatusError(t *testing.T) {
	job := &Job{
		Pid:         123,
		Description: "Test Job",
		Start:       time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
		End:         time.Date(0, 0, 0, 0, 10, 1, 0, time.UTC),
	}

	job.setStatus()

	expected := "ERROR"
	if job.Status != expected {
		t.Errorf("Expected Status to be '%s', got %s", expected, job.Status)
	}
}
func TestJob_SetStatusWarning(t *testing.T) {
	job := &Job{
		Pid:         123,
		Description: "Test Job",
		Start:       time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
		End:         time.Date(0, 0, 0, 0, 5, 1, 0, time.UTC),
	}

	job.setStatus()

	expected := "WARNING"
	if job.Status != expected {
		t.Errorf("Expected Status to be '%s', got %s", expected, job.Status)
	}
}

func TestJob_SetStatusSuccess(t *testing.T) {
	job := &Job{
		Pid:         123,
		Description: "Test Job",
		Start:       time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
		End:         time.Date(0, 0, 0, 0, 4, 0, 0, time.UTC),
	}

	job.setStatus()

	expected := "SUCCESS"
	if job.Status != expected {
		t.Errorf("Expected Status to be '%s', got %s", expected, job.Status)
	}
}

func TestJob_SetStatusIncomplete(t *testing.T) {
	job := &Job{
		Pid:         123,
		Description: "Test Job",
		Start:       time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
	}

	job.setStatus()

	expected := "INCOMPLETE"
	if job.Status != expected {
		t.Errorf("Expected Status to be '%s', got %s", expected, job.Status)
	}
}
