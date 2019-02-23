package timewheel

import "time"

const (
	HOUR_SCALE   = 24
	MINUTE_SCALE = 60
	SECOND_SCALE = 60

	DIVISOR = 10 // 允许误差 time.Second / 10 = 100ms
)

type Schedule struct {
	*TimeWheel
}

func NewSchedule() *Schedule {
	seconds := NewTimeWheel(time.Second, SECOND_SCALE, DIVISOR)

	minutes := NewTimeWheel(time.Minute, MINUTE_SCALE, DIVISOR)
	minutes.AddNext(seconds)

	hours := NewTimeWheel(time.Hour, HOUR_SCALE, DIVISOR)
	hours.AddNext(minutes)

	return &Schedule{hours}
}

func (s *Schedule) Run() {
	s.Start()
	next := s.next
	for next != nil {
		next.Start()
		next = next.next
	}
}
