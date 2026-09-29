package main

// AnalyzeRequest represents the incoming JSON request containing raw log lines.
type AnalyzeRequest struct {
	Lines []string `json:"lines"`
}

// AnalyzeResponse represents the structured JSON threat analysis report.
type AnalyzeResponse struct {
	TotalRequests       int            `json:"total_requests"`
	BruteForceDetected bool           `json:"brute_force_detected"`
	SuspiciousIPs      []string       `json:"suspicious_ips"`
	FailedAttemptsPerIP map[string]int `json:"failed_attempts_per_ip"`
	Severity           string         `json:"severity"`
}

// HealthResponse represents the server status response.
type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

// LogEntry holds internal parsed data from a single log line.
type LogEntry struct {
	IP         string
	StatusCode int
	IsValid    bool
}