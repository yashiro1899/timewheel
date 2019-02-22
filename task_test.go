package timewheel

import (
	"log"
	"testing"
	"time"
)

func TestTask_Call(t *testing.T) {
	tests := []struct {
		name     string
		callback func(...interface{})
	}{
		{"normal", func(v ...interface{}) { log.Printf("%v\n", v) }},
		{"error", func(v ...interface{}) { panic("panic") }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := NewTaskAfter([]interface{}{"test"}, tt.callback, time.Second)
			task.Call()
		})
	}
}
