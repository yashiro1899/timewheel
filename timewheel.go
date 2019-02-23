package timewheel

import (
	"time"
)

type TimeWheel struct {
	ticker   *time.Ticker
	interval time.Duration
	scale    int
	divisor  int // interval/divisor 控制精度
	counter  int
	current  int
	slots    []map[uint64]*Task
	next     *TimeWheel

	addTaskC    chan *Task
	removeTaskC chan uint64
	stopC       chan struct{}
}

func NewTimeWheel(interval time.Duration, scale, divisor int) *TimeWheel {
	return &TimeWheel{
		interval:    interval,
		divisor:     divisor,
		scale:       scale,
		slots:       make([]map[uint64]*Task, scale),
		addTaskC:    make(chan *Task),
		removeTaskC: make(chan uint64),
		stopC:       make(chan struct{}),
	}
}

func (tw *TimeWheel) Start() {
	if tw.ticker == nil {
		tw.ticker = time.NewTicker(tw.interval / time.Duration(tw.divisor))
		go tw.start()
	}
}

func (tw *TimeWheel) Stop() {
	tw.stopC <- struct{}{}
}

func (tw *TimeWheel) AddTask(t *Task) {
	tw.addTaskC <- t
}

func (tw *TimeWheel) AddNext(next *TimeWheel) {
	tw.next = next
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
		case now := <-tw.ticker.C:
			tw.tickHandler(now)
		case t := <-tw.addTaskC:
			tw.addTask(t)
		// case tw.removeTaskC:
		case <-tw.stopC:
			tw.ticker.Stop()
			return
		}
	}
}

func (tw *TimeWheel) tickHandler(now time.Time) {
	if tw.next == nil {
		tw.runTasks(now)
	}

	tw.counter++
	if tw.counter == tw.divisor {
		tw.counter = 0
		tw.current++
		tw.current = tw.current % tw.scale
	}
}

func (tw *TimeWheel) runTasks(now time.Time) {
	for _, t := range tw.slots[tw.current] {
		if now.After(t.expiredAt) {
			go t.Call() // TODO: 控制数量
			delete(tw.slots[tw.current], t.Id)
		}
	}
}

func (tw *TimeWheel) addTask(t *Task) {
	now := time.Now()
	gap := t.expiredAt.Sub(now)

	if gap < tw.interval {
		if tw.next == nil {
			tw.slots[tw.current][t.Id] = t
		} else {
			tw.next.AddTask(t)
		}
	} else {
		pos := (tw.current + int(gap/tw.interval)) % tw.scale
		tw.slots[pos][t.Id] = t
	}
}
