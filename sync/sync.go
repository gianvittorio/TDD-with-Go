package sync

import "sync"

type Counter struct {
	value int
	lock sync.Mutex
}

func (counter *Counter) Inc() {
	counter.lock.Lock()
	defer counter.lock.Unlock()

	counter.value++
}

func (counter *Counter) Value() int {
	return counter.value
}

func NewCounter() *Counter {
	return &Counter{}
}
