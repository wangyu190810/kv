package Storage

import (
	"container/list"
	"sync"
)

type Queue struct {
	List  *list.List
	Mutex sync.Mutex
}

var DataQueue = Queue{
	List: list.New(),
}

func (q *Queue) Add(v interface{}) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	q.List.PushBack(v)
}

func (q *Queue) Remove(e *list.Element) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	q.List.Remove(e)
}

func (q *Queue) Get() interface{} {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	return q.List.Front()
}
