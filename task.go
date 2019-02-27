package timewheel

import (
	"log"
	"time"

	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func init() {
	st, _ := time.Parse("2016-01-02", "2019-01-01")
	settings := sonyflake.Settings{StartTime: st}
	sf = sonyflake.NewSonyflake(settings)
}

type Task struct {
	Id        uint64
	args      []interface{}
	callback  func(...interface{})
	expiredAt time.Time
}

func NewTask(
	args []interface{},
	callback func(...interface{}),
	expiredAt time.Time,
) *Task {
	t := new(Task)
	t.Id, _ = sf.NextID()
	t.args = args
	t.callback = callback
	t.expiredAt = expiredAt
	return t
}

func NewTaskAfter(
	args []interface{},
	callback func(...interface{}),
	delay time.Duration,
) *Task {
	now := time.Now()
	return NewTask(args, callback, now.Add(delay))
}

func (t *Task) Call() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("task(%x) with args(%v): %v\n", t.Id, t.args, err)
		}
	}()

	t.callback(t.args...)
}
