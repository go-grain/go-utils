package go_utils

import (
	"sync"
)

type StrQueue struct {
	store sync.Map
	len   int
}

func NewQueue() *StrQueue {
	return &StrQueue{}
}

func (q *StrQueue) Enqueue(val string) {
	q.store.Store(val, val)
	q.len++
}

func (q *StrQueue) Dequeue() (s string) {
	if q.Empty() {
		q.AddVal()
	}
	q.store.Range(func(key, value any) bool {
		s = value.(string)
		return false
	})
	q.store.Delete(s)
	q.len--
	return s
}

func (q *StrQueue) Empty() bool {
	return q.len <= 0
}

func (q *StrQueue) Range() int {
	count := 0
	q.store.Range(func(key, value any) bool {
		count++
		return true
	})
	return count
}

func (q *StrQueue) AddVal() {
	for i := 0; i < 500; i++ {
		q.Enqueue(RandomCharset(6))
	}
}
