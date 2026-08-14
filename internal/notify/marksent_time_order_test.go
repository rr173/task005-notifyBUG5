package notify

import (
	"testing"
	"time"
)

func TestMarkSentRejectsTimeBeforeCreatedAt(t *testing.T) {
	s := New()
	now := time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)
	s.Create(CreateInput{ID: "TS1", Recipient: "u", Content: "c"}, now)

	pastTime := now.Add(-1 * time.Hour)
	_, err := s.MarkSent("TS1", pastTime)
	if err == nil {
		t.Errorf("MarkSent should reject sentTime before CreatedAt: sentTime=%v, createdAt=%v",
			pastTime, now)
	}
}

func TestMarkSentAllowsTimeAfterCreatedAt(t *testing.T) {
	s := New()
	now := time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)
	s.Create(CreateInput{ID: "TS2", Recipient: "u", Content: "c"}, now)

	futureTime := now.Add(1 * time.Hour)
	_, err := s.MarkSent("TS2", futureTime)
	if err != nil {
		t.Errorf("MarkSent should allow sentTime after CreatedAt: %v", err)
	}
}
