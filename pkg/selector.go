package lb

import "sync"

type Selector struct {
	mu  sync.Mutex
	idx uint64
}

func (s *Selector) Next() uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.idx += 1
	return s.idx
}

func NewSelector() *Selector {
	return &Selector{
		idx: 0,
	}
}
