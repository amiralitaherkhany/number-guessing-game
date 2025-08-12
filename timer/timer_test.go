package timer

import (
	"fmt"
	"testing"
	"time"
)

type mockTimer struct {
	*Timer
}

func newMockTimer(startTime time.Time, endTime time.Time) *mockTimer {
	return &mockTimer{
		Timer: &Timer{
			startTime: startTime,
			endTime:   endTime,
		},
	}
}

func TestNewTimer(t *testing.T) {
	timer := New()
	if !timer.startTime.IsZero() || !timer.endTime.IsZero() {
		t.Error("start and end times must be zero valued in the New function")
	}
}

func TestGetTime(t *testing.T) {
	tests := []struct {
		startTime time.Time
		endTime   time.Time
		expected  string
	}{
		{
			startTime: time.Date(2025, 8, 12, 10, 0, 0, 0, time.UTC),
			endTime:   time.Date(2025, 8, 12, 11, 0, 0, 0, time.UTC),
			expected:  "01:00:00",
		},
		{
			startTime: time.Date(2025, 8, 12, 10, 0, 0, 0, time.UTC),
			endTime:   time.Date(2025, 8, 12, 10, 45, 30, 0, time.UTC),
			expected:  "00:45:30",
		},
		{
			startTime: time.Date(2025, 8, 12, 8, 30, 0, 0, time.UTC),
			endTime:   time.Date(2025, 8, 12, 10, 45, 5, 0, time.UTC),
			expected:  "02:15:05",
		},
		{
			startTime: time.Date(2025, 8, 12, 10, 0, 0, 0, time.UTC),
			endTime:   time.Date(2025, 8, 12, 10, 0, 0, 0, time.UTC),
			expected:  "00:00:00",
		},
		{
			startTime: time.Date(2025, 8, 11, 9, 0, 0, 0, time.UTC),
			endTime:   time.Date(2025, 8, 12, 10, 0, 0, 0, time.UTC),
			expected:  "25:00:00",
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("tessting %s time duration", test.expected), func(t *testing.T) {
			timer := newMockTimer(test.startTime, test.endTime)
			result := timer.GetTime()
			if result != test.expected {
				t.Fatalf("test failed,expected: %s to be returned", test.expected)
			}
		})
	}
}

func TestStartTimer(t *testing.T) {
	timer := New()

	before := time.Now()
	timer.StartTimer()
	after := time.Now()

	if timer.startTime.IsZero() || timer.startTime.Before(before) || timer.startTime.After(after) {
		t.Errorf("startTime = %v; want between %v and %v", timer.startTime, before, after)
	}
}

func TestStopTimer(t *testing.T) {
	timer := New()

	before := time.Now()
	timer.StopTimer()
	after := time.Now()

	if timer.endTime.IsZero() || timer.endTime.Before(before) || timer.endTime.After(after) {
		t.Errorf("endTime = %v; want between %v and %v", timer.endTime, before, after)
	}
}
