package timewheel

import "time"

type TimeWheel struct {
	ticker    time.Ticker
	interval  time.Duration
	tolerance time.Duration
	scale     int
	current   int
	slots     []map[uint64]*Task
	next      *TimeWheel

	addTaskC    chan *Task
	removeTaskC chan uint64
	stopC       chan bool
}

func NewTimeWheel(interval, tolerance time.Duration, scale int) *TimeWheel {
	return &TimeWheel{
		interval:    interval,
		tolerance:   tolerance,
		scale:       scale,
		addTaskC:    make(chan *Task),
		removeTaskC: make(chan uint64),
		stopC:       make(chan bool),
	}
}
