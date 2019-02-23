package timewheel

import (
	"log"
	"testing"
	"time"
)

func TestSchedule_Run(t *testing.T) {
	return
	s := NewSchedule()
	s.Run()
	log.Println(time.Now())

	task1 := NewTaskAfter([]interface{}{"1s"}, f, time.Second)
	task61 := NewTaskAfter([]interface{}{"61s"}, f, 61*time.Second)
	task3 := NewTaskAfter([]interface{}{"3hours"}, f, 3*time.Hour)

	s.AddTask(task1)
	s.AddTask(task61)
	s.AddTask(task3)
	s.RemoveTask(task3.Id)

	time.Sleep(62 * time.Second)
}
