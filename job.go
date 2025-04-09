package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	timeFormat      = "15:04:05"
	errorDuration   = 600
	warningDuration = 300
)

type Job struct {
	Pid         int
	Description string
	Start       time.Time
	End         time.Time
	Status      string
}

func (j *Job) SetTime(verb string, timestamp string) error {
	time, err := time.Parse(timeFormat, timestamp)
	if err != nil {
		return err
	}

	verb = strings.TrimSpace(verb)
	if verb == "START" {
		if !j.Start.IsZero() {
			return fmt.Errorf("start time is already set")
		}
		j.Start = time
	} else if verb == "END" {
		if !j.End.IsZero() {
			return fmt.Errorf("end time is already set")
		}
		j.End = time
	} else {
		return fmt.Errorf("unknown job verb %s", verb)
	}
	j.setStatus()
	return nil
}

func (j *Job) setStatus() {
	if j.Start.IsZero() || j.End.IsZero() {
		j.Status = "INCOMPLETE"
		return
	}

	duration := j.End.Sub(j.Start)
	if duration.Seconds() > errorDuration {
		j.Status = "ERROR"
	} else if duration.Seconds() > warningDuration {
		j.Status = "WARNING"
	} else {
		j.Status = "SUCCESS"
	}
}

func (j *Job) ToSlice() []string {
	return []string{
		strconv.Itoa(j.Pid),
		j.Description,
		timeToString(j.Start, timeFormat),
		timeToString(j.End, timeFormat),
		j.Status,
	}
}

func timeToString(t time.Time, layout string) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(layout)
}

func GetHeaders() []string {
	return []string{"pid", "description", "start", "end", "status"}
}
