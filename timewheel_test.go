package timewheel

import (
	"log"
	"testing"
	"time"
)

func f(v ...interface{}) {
	log.Println(v...)
}

func TestTimeWheel_Start(t *testing.T) {
	tw := NewTimeWheel(10*time.Second, 3600, 10)
	tw.Start()
	log.Println("start")

	task1 := NewTaskAfter([]interface{}{"1s"}, f, time.Second)
	task13 := NewTaskAfter([]interface{}{"13s"}, f, 13*time.Second)
	task25 := NewTaskAfter([]interface{}{"25s"}, f, 25*time.Second)

	tw.AddTask(task1)
	tw.AddTask(task13)
	tw.AddTask(task25)

	time.Sleep(30 * time.Second)
	tw.Stop()
	time.Sleep(time.Second)
}
