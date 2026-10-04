package version

import (
	"testing"
	"time"
)

func TestMetadata(t *testing.T) {
	Version = "v0.8.0"
	BuildCommit = "abc1234"
	BuildBranch = "main"
	BuildTime = "2026-01-01"

	m := Metadata()
	for key, want := range map[string]string{
		"version":      "v0.8.0",
		"build_commit": "abc1234",
		"build_branch": "main",
		"build_time":   "2026-01-01",
	} {
		if got := m[key]; got != want {
			t.Errorf("m[%q] = %q, want %q", key, got, want)
		}
	}
}

func TestStartTime(t *testing.T) {
	if StartTime.IsZero() {
		t.Error("expected non-zero StartTime")
	}
	if time.Since(StartTime) > time.Minute {
		t.Error("StartTime should be recent")
	}
}
