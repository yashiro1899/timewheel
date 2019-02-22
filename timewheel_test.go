package timewheel

import (
	"testing"
	"time"
)

func TestTimeWheel_Start(t *testing.T) {
	tw := NewTimeWheel(10*time.Second, 3600, 10)
	tw.Start()

	time.Sleep(22 * time.Second)
	tw.Stop()
	time.Sleep(time.Minute)
}
