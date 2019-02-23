package timewheel

import (
	"log"
	"testing"
	"time"
)

func f(v ...interface{}) {
	log.Println(time.Now(), v)
}

func TestTimeWheel_Start(t *testing.T) {
	tw := NewTimeWheel(time.Second, 3600, 20)
	tw.Start()
	log.Println(time.Now())

	task1 := NewTaskAfter([]interface{}{"1s"}, f, 130*time.Millisecond)
	task5 := NewTaskAfter([]interface{}{"5s"}, f, 500*time.Millisecond)
	task13 := NewTaskAfter([]interface{}{"13s"}, f, 1300*time.Millisecond)
	task25 := NewTaskAfter([]interface{}{"25s"}, f, 2500*time.Millisecond)

	tw.AddTask(task1)
	tw.AddTask(task5)
	tw.AddTask(task13)
	tw.AddTask(task25)
	tw.RemoveTask(task5.Id)

	time.Sleep(3 * time.Second)
	tw.Stop()
	time.Sleep(time.Second)
}
