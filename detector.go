package main

const BruteForceThreshold = 3

// DetectThreats evaluates parsed logs for 401/403 status code spikes per IP.
func DetectThreats(entries []*LogEntry, totalRequests int) AnalyzeResponse {
	failedAttempts := make(map[string]int)
	suspiciousIPs := make([]string, 0)

	for _, entry := range entries {
		if entry.StatusCode == 401 || entry.StatusCode == 403 {
			failedAttempts[entry.IP]++
		}
	}

	bruteForceDetected := false
	for ip, count := range failedAttempts {
		if count >= BruteForceThreshold {
			suspiciousIPs = append(suspiciousIPs, ip)
			bruteForceDetected = true
		}
	}

	// Classify threat severity
	severity := "LOW"
	if bruteForceDetected {
		if len(suspiciousIPs) > 3 || maxCount(failedAttempts) >= 10 {
			severity = "CRITICAL"
		} else {
			severity = "HIGH"
		}
	} else if len(failedAttempts) > 0 {
		severity = "MEDIUM"
	}

	return AnalyzeResponse{
		TotalRequests:       totalRequests,
		BruteForceDetected: bruteForceDetected,
		SuspiciousIPs:      suspiciousIPs,
		FailedAttemptsPerIP: failedAttempts,
		Severity:           severity,
	}
}

func maxCount(m map[string]int) int {
	maxVal := 0
	for _, val := range m {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}