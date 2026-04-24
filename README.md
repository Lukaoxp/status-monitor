# status-monitor

[![Go Version](https://img.shields.io/badge/Go-1.26.2-blue)](https://go.dev/)
[![CI](https://github.com/Lukaoxp/status-monitor/actions/workflows/ci.yml/badge.svg)](https://github.com/Lukaoxp/status-monitor/actions/workflows/ci.yml)
[![Project Status](https://img.shields.io/badge/status-WIP-yellow)](https://github.com/Lukaoxp/status-monitor)

Lightweight HTTP health monitoring service in Go. Built around production readiness: multi-stage Docker builds, graceful shutdown, structured logging, and Prometheus metrics.

---

## Architecture

```
status-monitor/
├── cmd/server/
│   ├── main.go      # Entry point, graceful shutdown
│   ├── server.go    # HTTP handlers
│   └── env.go       # Environment helpers
├── internal/health/
│   ├── health.go      # Domain logic, Service struct
│   └── health_test.go # Unit tests
└── Dockerfile         # Multi-stage build (~15MB final image)
```

**Key decisions:**

* **No external HTTP framework** — `net/http` stdlib only. Zero unnecessary dependencies.
* **Manual dependency injection** — `Server` struct holds dependencies. No globals, no init() side effects.
* **Multi-stage Dockerfile** — builder stage with `golang:1.26-alpine`, final stage with plain `alpine`. Result: ~15MB image.
* **Graceful shutdown** — captures `SIGTERM`/`SIGINT`, waits for in-flight requests (5s timeout) before stopping.

---

## Features

| Feature | Status |
|---|---|
| Health check endpoint (`/status`) | ✅ Implemented |
| JSON response (status, version, uptime) | ✅ Implemented |
| Graceful shutdown | ✅ Implemented |
| Multi-stage Docker build | ✅ Implemented |
| Unit tests | ✅ Implemented |
| CI/CD (GitHub Actions) | ✅ Implemented |
| Prometheus metrics (`/metrics`) | Planned |
| Structured logging (log/slog) | Planned |
| Liveness/readiness probes | Planned |
| Cloud deployment | Planned |

---

## Running locally

**With Go:**

```bash
go run ./cmd/server
```

**With Docker:**

```bash
docker build -t status-monitor .
docker run -p 8080:8080 status-monitor
```

**Custom port:**

```bash
docker run -e PORT=9090 -p 9090:9090 status-monitor
```

**Access:** `http://localhost:8080/status`

**Response format:**

```json
{
  "status": "Up",
  "version": "1.0.0",
  "uptime": 3600
}
```

---

## Tests

```bash
go test ./...

# With coverage
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

---

## CI/CD

GitHub Actions runs on every pull request:

1. `go test -race ./...`
2. `go vet ./...`
3. `gosec ./...` — static security analysis
4. `govulncheck ./...` — known vulnerability scan

---

## Roadmap

- [x] Core service with uptime tracking
- [x] Graceful shutdown (SIGTERM/SIGINT, 5s timeout)
- [x] Multi-stage Dockerfile (~15MB final image)
- [x] Unit tests for domain logic
- [x] CI/CD pipeline (test + vet + gosec + govulncheck)
- [ ] Structured logging with `log/slog` (JSON, request IDs)
- [ ] Prometheus metrics endpoint (`/metrics`)
- [ ] Integration tests with testcontainers
- [ ] Cloud deployment (Google Cloud Run)
- [ ] Kubernetes liveness/readiness probes

---

Built by [Lucas Carturani](https://github.com/Lukaoxp)
