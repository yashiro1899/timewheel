package timewheel

import (
	"log"
	"time"
)

type TimeWheel struct {
	ticker   *time.Ticker
	interval time.Duration
	scale    int
	divisor  int
	counter  int
	current  int
	slots    []map[uint64]*Task
	next     *TimeWheel

	addTaskC    chan *Task
	removeTaskC chan uint64
	stopC       chan bool
}

func NewTimeWheel(interval time.Duration, scale, divisor int) *TimeWheel {
	return &TimeWheel{
		interval:    interval,
		divisor:     divisor,
		scale:       scale,
		slots:       make([]map[uint64]*Task, scale),
		addTaskC:    make(chan *Task),
		removeTaskC: make(chan uint64),
		stopC:       make(chan bool),
	}
}

func (tw *TimeWheel) Start() {
	tw.ticker = time.NewTicker(tw.interval / time.Duration(tw.divisor))
	go tw.start()
}

func (tw *TimeWheel) initSlots() {
	for i, _ := range tw.slots {
		tw.slots[i] = make(map[uint64]*Task)
	}
}

func (tw *TimeWheel) start() {
	tw.initSlots()

	for {
		select {
		case <-tw.ticker.C:
			tw.tickHandler()
			// case tw.addTaskC:
			// case tw.removeTaskC:
			// case tw.stopC:
		}
	}
}

func (tw *TimeWheel) tickHandler() {
	tw.counter++

	log.Println(tw.counter, tw.current)
	if tw.counter == tw.divisor {
		tw.counter = 0
		tw.current++
		tw.current = tw.current % tw.scale
	}
}
