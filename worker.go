package main

import (
	"sync"
)

const MaxWorkers = 50

// ProcessLogs distributes log parsing across a bounded pool of goroutines.
func ProcessLogs(lines []string) []*LogEntry {
	totalJobs := len(lines)
	if totalJobs == 0 {
		return []*LogEntry{}
	}

	numWorkers := MaxWorkers
	if totalJobs < numWorkers {
		numWorkers = totalJobs
	}

	jobs := make(chan string, totalJobs)
	results := make(chan *LogEntry, totalJobs)

	var wg sync.WaitGroup

	// Launch bounded goroutine worker pool
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for line := range jobs {
				results <- ParseLogLine(line)
			}
		}()
	}

	// Queue log lines
	for _, line := range lines {
		jobs <- line
	}
	close(jobs)

	// Wait for processing to complete and close results channel
	wg.Wait()
	close(results)

	// Collect parsed log entries
	entries := make([]*LogEntry, 0, totalJobs)
	for entry := range results {
		if entry.IsValid {
			entries = append(entries, entry)
		}
	}

	return entries
}