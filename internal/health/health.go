package health

import "time"

type Service struct {
	version string
}

func NewService(v string) *Service {
	return &Service{
		version: v,
	}
}

// Health represents the system status information
type Health struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	Uptime  int64  `json:"uptime"`
}

var startTime time.Time

func init() {
	startTime = time.Now()
}

func (s *Service) GetStatus() Health {
	return Health{
		Status:  "Up",
		Version: s.version,
		Uptime:  int64(time.Since(startTime).Seconds()),
	}
}
