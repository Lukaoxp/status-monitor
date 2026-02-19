package health

import "time"

type Service struct {
	version   string
	startTime time.Time
}

func NewService(v string) *Service {
	return &Service{
		version:   v,
		startTime: time.Now(),
	}
}

// Health represents the system status information
type Health struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	Uptime  int64  `json:"uptime"`
}

func (s *Service) GetStatus() Health {
	return Health{
		Status:  "Up",
		Version: s.version,
		Uptime:  int64(time.Since(s.startTime).Seconds()),
	}
}
