package tsqueue

import "sync"

const MAX = 10

type TSQueue struct {
	// Synchronization variables
	lock *sync.Mutex
	// State variables
	items     [MAX]int
	front     int
	nextEmpty int
}

func NewTSQueue() *TSQueue {
	return &TSQueue{lock: &sync.Mutex{}}
}

func (q *TSQueue) tryInsert(item int) (success bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.nextEmpty-q.front < MAX {
		q.items[q.nextEmpty%MAX] = item
		q.nextEmpty++
		success = true
	}
	return
}

func (q *TSQueue) tryRemove(item *int) (success bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.front < q.nextEmpty {
		*item = q.items[q.front%MAX]
		q.front++
		success = true
	}
	return
}
