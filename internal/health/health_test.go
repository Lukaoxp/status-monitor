package health

import (
	"testing"
)

func TestNewService(t *testing.T) {
	version := "1.0.0"
	svc := NewService(version)

	if svc == nil {
		t.Fatal("NewService returned nil")
	}

	if svc.version != version {
		t.Errorf("expected version %q, got %q", version, svc.version)
	}
}

func TestGetStatus(t *testing.T) {
	version := "1.0.0"
	svc := NewService(version)

	result := svc.GetStatus()
	if result.Version != version {
		t.Errorf("expected version %q, got %q", version, result.Version)
	}
	if result.Status != "Up" {
		t.Errorf("expected status Up, got %v", result.Status)
	}
	if result.Uptime < 0 {
		t.Errorf("uptime should not be negative, got %d", result.Uptime)
	}
}
