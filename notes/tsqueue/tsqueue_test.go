package tsqueue

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

const THREADS = 3

var wg sync.WaitGroup

func TestTSQueue(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var q [THREADS]*TSQueue
	for i := 0; i < THREADS; i++ {
		q[i] = NewTSQueue()
		wg.Add(1)
		go putSome(q[i])
	}
	time.Sleep(10 * time.Millisecond)
	// wg.Wait()

	for i := 0; i < THREADS; i++ {
		wg.Add(1)
		go testRemoval(q[i], i)
	}
	wg.Wait()
}

func putSome(q *TSQueue) {
	for i := 0; i < 50; i++ {
		q.tryInsert(i)
	}
	wg.Done()
}

func testRemoval(q *TSQueue, id int) {
	var item int
	for i := 0; i < 20; i++ {
		if q.tryRemove(&item) {
			fmt.Printf("Removed item %d from queue %d\n", item, id)
		} else {
			fmt.Println("Nothing there to remove.")
			runtime.Gosched()
		}
	}
	wg.Done()
}
