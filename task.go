package timewheel

import (
	"log"
	"time"
)

type Task struct {
	args      []interface{}
	callback  func(...interface{})
	expiredAt time.Time
}

func (t *Task) Call() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("task with args(%v): %v\n", err)
		}
	}()

	t.callback(t.args...)
}
