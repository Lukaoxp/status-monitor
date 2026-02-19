package health

import (
	"testing"
)

func TestNewService(t *testing.T) {
	version := "1.0.0"
	svc := NewService(version)

	if svc == nil {
		t.Fatal("NewService retornou nil")
	}

	if svc.version != version {
		t.Errorf("esperava version %q, got %q", version, svc.version)
	}
}

func TestGetStatus(t *testing.T) {
	version := "1.0.0"
	svc := NewService(version)

	result := svc.GetStatus()
	if result.Version != version {
		t.Errorf("esperava version %q, got %q", version, result.Version)
	}
	if result.Status != "Up" {
		t.Errorf("Status incorreto, deveria ser Up, mas veio %v", result.Status)
	}
	if result.Uptime < 0 {
		t.Errorf("Uptime nao deveria ser negativo, got %d", result.Uptime)
	}
}
