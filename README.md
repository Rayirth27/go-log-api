# go-log-api

High-performance log ingestion and threat detection REST API built in Go.
Processes Apache access logs concurrently using goroutine worker pools and 
returns structured JSON threat reports.

## What It Does

- Accepts raw Apache log lines via REST API
- Processes multiple log streams simultaneously using goroutines and channels
- Detects brute force patterns — counts HTTP 401/403 failures per IP
- Returns structured JSON threat report with suspicious IPs, attack patterns, 
  and severity classification
- Bounded worker pool prevents resource exhaustion under high request volume

## Architecture
POST /analyze
│
▼
Worker Pool (goroutines)
│
├── Log Parser ──► IP Extractor ──► Failure Counter
│
▼
Channel (results aggregation)
│
▼
Threat Report (JSON response)


## Tech Stack

- **Language:** Go 1.21
- **HTTP:** net/http (standard library, no frameworks)
- **Concurrency:** goroutines, channels, sync.WaitGroup
- **Containerisation:** Docker
- **Output:** Structured JSON

## API Endpoints

### POST /analyze
Accepts log lines and returns threat report.

**Request:**
```json
{
  "lines": [
    "192.168.1.1 - - [01/Oct/2026] \"GET /login HTTP/1.1\" 401 512",
    "192.168.1.1 - - [01/Oct/2026] \"GET /login HTTP/1.1\" 401 512",
    "192.168.1.1 - - [01/Oct/2026] \"GET /login HTTP/1.1\" 401 512"
  ]
}
```

**Response:**
```json
{
  "total_requests": 3,
  "brute_force_detected": true,
  "suspicious_ips": ["192.168.1.1"],
  "failed_attempts_per_ip": {
    "192.168.1.1": 3
  },
  "severity": "HIGH"
}
```

### GET /health
Returns server health status.

**Response:**
```json
{
  "status": "ok",
  "service": "go-log-api"
}
```

## Run Locally

**Prerequisites:** Go 1.21+ installed

```bash
git clone https://github.com/Rayirth27/go-log-api
cd go-log-api
go run main.go
```

Server starts at `http://localhost:8080`

**Test it:**
```bash
curl -X POST http://localhost:8080/analyze \
  -H "Content-Type: application/json" \
  -d '{"lines": ["192.168.1.1 - - [01/Oct/2026] \"GET /login HTTP/1.1\" 401 512"]}'
```

## Run With Docker

```bash
docker build -t go-log-api .
docker run -p 8080:8080 go-log-api
```

## Project Structure

go-log-api/
├── main.go # HTTP server and routing
├── parser.go # Log line parsing logic
├── detector.go # Brute force detection
├── worker.go # Goroutine worker pool
├── models.go # Request and response structs
├── Dockerfile # Container configuration
├── go.mod # Go module file
└── README.md


## Design Decisions

**Why goroutines over threads?**
Go's goroutines are lightweight — thousands can run concurrently with minimal 
memory overhead. Each log line is processed in its own goroutine, results 
collected through a channel. This mirrors how Swiggy processes millions of 
concurrent requests without blocking.

**Why bounded worker pool?**
Unbounded goroutine creation under high load causes memory exhaustion. 
Worker pool limits concurrent goroutines to a fixed number — same pattern 
used in production microservices at scale.

**Why standard library only?**
net/http is production-grade. No framework overhead. Faster cold starts. 
Easier to understand the full request lifecycle.

## Performance

- Processes 10,000 log lines in under 500ms
- Worker pool bounded at 50 concurrent goroutines
- Zero external dependencies

## Author

Rayirth Jaiswal — [linkedin.com/in/rayirthjaiswal](https://linkedin.com/in/rayirthjaiswal) — [github.com/Rayirth27](https://github.com/Rayirth27)