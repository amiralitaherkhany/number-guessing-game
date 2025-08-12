package timer

import (
	"fmt"
	"time"
)

type Timer struct {
	startTime time.Time
	endTime   time.Time
}

func (timer *Timer) GetTime() string {
	elapsed := timer.endTime.Sub(timer.startTime)
	hours := int(elapsed.Hours())
	minutes := int(elapsed.Minutes()) % 60
	seconds := int(elapsed.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func (timer *Timer) StartTimer() {
	timer.startTime = time.Now()
}
func (timer *Timer) StopTimer() {
	timer.endTime = time.Now()
}

func New() *Timer {
	return new(Timer)
}
