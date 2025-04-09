package main

import (
	"encoding/csv"
	"io"
	"log/slog"
	"os"
	"strconv"
)

func main() {
	inputFile, reader := getInput()
	defer inputFile.Close()
	outputFile, writer := createOutput()
	defer outputFile.Close()
	defer writer.Flush()

	jobs := make(map[int]*Job)
	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			slog.Error("Failed to read line", "error", err)
			continue
		}
		processRow(row, jobs, writer)
	}

	for _, job := range jobs {
		slog.Info("Writing to csv", "job", job)
		if err := writer.Write(job.ToSlice()); err != nil {
			slog.Error("Failed to write csv", "job", job, "error", err)
		}
	}
}

func processRow(row []string, jobs map[int]*Job, writer *csv.Writer) {
	pid, err := strconv.Atoi(row[3])
	if err != nil {
		slog.Error("Failed to parse pid", "row", row, "error", err)
		return
	}

	job, exists := jobs[pid]
	if exists {
		if err := job.SetTime(row[2], row[0]); err != nil {
			slog.Error("Failed to set job time", "job", job, "row", row, "error", err)
			return
		}

		slog.Info("Writing to csv", "job", job)
		if err := writer.Write(job.ToSlice()); err != nil {
			slog.Error("Failed to write csv", "job", job, "error", err)
			return
		}
		delete(jobs, pid)
	} else {
		job := &Job{
			Pid:         pid,
			Description: row[1],
		}
		if err := job.SetTime(row[2], row[0]); err != nil {
			slog.Error("Failed to set job time", "job", job, "row", row, "error", err)
			return
		}
		jobs[pid] = job
	}
}

func getInput() (*os.File, *csv.Reader) {
	input := "logs.log"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	return file, reader
}

func createOutput() (*os.File, *csv.Writer) {
	output := "output.csv"
	if len(os.Args) > 2 {
		output = os.Args[2]
	}

	file, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	writer.Write(GetHeaders())
	return file, writer
}
