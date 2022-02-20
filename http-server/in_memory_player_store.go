package main

import "sync"

type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.Mutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		sync.Mutex{},
	}
}

func (i *InMemoryPlayerStore) getPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) recordWin(name string) {
	i.mu.Lock()
	i.store[name]++
	i.mu.Unlock()
}
