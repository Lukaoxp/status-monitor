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
