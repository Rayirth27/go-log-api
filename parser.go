package main

import (
	"regexp"
	"strconv"
	"strings"
)

// Regex matching Apache log IP and HTTP Status Code
// Example line: 192.168.1.1 - - [01/Oct/2026] "GET /login HTTP/1.1" 401 512
var logPattern = regexp.MustCompile(`^(\S+)\s+\S+\s+\S+\s+\[.*?\]\s+"[^"]*"\s+(\d{3})`)

// ParseLogLine extracts IP address and HTTP status code from a single log line.
func ParseLogLine(line string) *LogEntry {
	line = strings.TrimSpace(line)
	if line == "" {
		return &LogEntry{IsValid: false}
	}

	matches := logPattern.FindStringSubmatch(line)
	if len(matches) < 3 {
		return &LogEntry{IsValid: false}
	}

	ip := matches[1]
	statusCode, err := strconv.Atoi(matches[2])
	if err != nil {
		return &LogEntry{IsValid: false}
	}

	return &LogEntry{
		IP:         ip,
		StatusCode: statusCode,
		IsValid:    true,
	}
}