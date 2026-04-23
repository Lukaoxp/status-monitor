# status-monitor

![Go Version](https://img.shields.io/badge/Go-1.26.2-blue)
![CI](https://github.com/Lukaoxp/status-monitor/actions/workflows/ci.yml/badge.svg)

A lightweight HTTP health monitoring service written in Go. Designed around production readiness: multi-stage Docker builds, graceful shutdown, structured logging, and Prometheus metrics.

---

## Architecture

```
status-monitor/
├── cmd/
│   └── server/
│       ├── env.go       # Environment variable helpers
│       ├── main.go      # Entry point, dependency wiring, graceful shutdown
│       └── server.go    # HTTP handlers, Server struct
├── internal/
│   └── health/
│       ├── health.go      # Domain logic, Service struct
│       └── health_test.go # Unit tests
└── Dockerfile             # Multi-stage build
```

**Key decisions:**

- **No external HTTP framework** — `net/http` stdlib only, zero unnecessary dependencies
- **Manual dependency injection** — `Server` struct holds dependencies; no globals, no init() side effects
- **Multi-stage Dockerfile** — builder stage with `golang:1.26-alpine`, final stage with plain `alpine` — resulting image ~15MB
- **Graceful shutdown** — captures `SIGTERM`/`SIGINT`, waits for in-flight requests (5s timeout) before stopping

---

## Endpoints

| Endpoint | Description | Status |
|----------|-------------|--------|
| `GET /status` | Returns service status, version, and uptime in seconds | ✅ Implemented |
| `GET /metrics` | Prometheus metrics | Planned |
| `GET /health/live` | Liveness probe (Kubernetes) | Planned |
| `GET /health/ready` | Readiness probe (Kubernetes) | Planned |

### Response format

```json
{
  "status": "Up",
  "version": "1.0.0",
  "uptime": 3600
}
```

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

Custom port:

```bash
docker run -e PORT=9090 -p 9090:9090 status-monitor
```

Access: `http://localhost:8080/status`

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

---

## Roadmap

- [x] Service struct with uptime tracking
- [x] Dependency injection via Server struct
- [x] Multi-stage Dockerfile (~15MB final image)
- [x] Graceful shutdown (SIGTERM/SIGINT, 5s timeout)
- [x] Unit tests for domain logic
- [ ] Structured logging with `log/slog` (JSON, request ID)
- [ ] Prometheus metrics endpoint (`/metrics`)
- [ ] Integration tests with testcontainers
- [ ] Cloud deploy (Google Cloud Run)
- [ ] Full CI/CD pipeline (test → build → push → deploy)

---

Built by [Lucas Carturani](https://github.com/Lukaoxp)
