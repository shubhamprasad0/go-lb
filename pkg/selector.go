package lb

import "sync"

// Selector is a thread-safe round-robin selector for choosing the next server index.
type Selector struct {
	mu  sync.Mutex
	idx uint64
}

// Next increments the index and returns the next server index in a thread-safe manner.
func (s *Selector) Next() uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.idx += 1
	return s.idx
}

// NewSelector creates and returns a new Selector instance.
func NewSelector() *Selector {
	return &Selector{
		idx: 0,
	}
}
