package tsqueue

import "sync"

const MAX = 10

type BBQueue struct {
	// Synchronization variables
	lock        sync.Mutex
	itemAdded   sync.Cond
	itemRemoved sync.Cond
	// State variables
	items     [MAX]int
	front     int
	nextEmpty int
}

func (q *BBQueue) insert(item int) {
	q.lock.Lock()
	for q.nextEmpty-q.front == MAX {
		q.itemRemoved.Wait()
	}
	q.items[q.nextEmpty%MAX] = item
	q.nextEmpty++
	q.itemAdded.Signal()
	q.lock.Unlock()
}

func (q *BBQueue) remove() (item int) {
	q.lock.Lock()
	for q.front == q.nextEmpty {
		q.itemAdded.Wait()
	}
	item = q.items[q.front%MAX]
	q.front++
	q.itemRemoved.Signal()
	q.lock.Unlock()
	return
}
