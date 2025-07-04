package store

import "sync"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{store: map[string]int{}, mutex: sync.Mutex{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
	mutex sync.Mutex
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	return i.store[name]
}
