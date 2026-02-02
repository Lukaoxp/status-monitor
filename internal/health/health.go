package health

import "time"

// Health represents the system status information
type Health struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	Uptime  int64  `json:"uptime"`
}

var StartTime time.Time

func init() {
	StartTime = time.Now()
}

func GetStatus() Health {
	return Health{
		Status:  "Up",
		Version: "1.0.0",
		Uptime:  int64(time.Since(StartTime).Seconds()),
	}
}
