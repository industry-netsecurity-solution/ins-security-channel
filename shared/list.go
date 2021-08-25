package shared

import (
	"container/list"
	"sync"
)

type ConcurrentList struct {
	sync.RWMutex
	list *list.List
}

func NewConcurrentList() *ConcurrentList {
	o := new(ConcurrentList)
	o.list = list.New()

	return o
}

func (v *ConcurrentList) Init() {
	//v.Lock()
	//defer v.Unlock()

	v.list.Init()
}

func (v *ConcurrentList) Len() int {
	//v.Lock()
	//defer v.Unlock()

	return v.list.Len()
}

func (v *ConcurrentList) Front() *list.Element {
	//v.Lock()
	//defer v.Unlock()

	return v.list.Front()
}

func (v *ConcurrentList) Back() *list.Element {
	//v.Lock()
	//defer v.Unlock()

	return v.list.Back()
}

func (v *ConcurrentList) Remove(e *list.Element) interface{} {
	//v.Lock()
	//defer v.Unlock()

	return v.list.Remove(e)
}

func (v *ConcurrentList) PushFront(e interface{}) *list.Element {
	//v.Lock()
	//defer v.Unlock()

	return v.list.PushFront(e)
}

func (v *ConcurrentList) PushBack(e interface{}) *list.Element {
	//v.Lock()
	//defer v.Unlock()

	return v.list.PushBack(e)
}

func (v *ConcurrentList) InsertBefore(e interface{}, mark *list.Element) *list.Element {
	//v.Lock()
	//defer v.Unlock()

	return v.list.InsertBefore(e, mark)
}

func (v *ConcurrentList) InsertAfter(e interface{}, mark *list.Element) *list.Element {
	//v.Lock()
	//defer v.Unlock()

	return v.list.InsertAfter(e, mark)
}

func (v *ConcurrentList) MoveToFront(e *list.Element) {
	//v.Lock()
	//defer v.Unlock()

	v.list.MoveToFront(e)
}

func (v *ConcurrentList) MoveToBack(e *list.Element) {
	//v.Lock()
	//defer v.Unlock()

	v.list.MoveToBack(e)
}

func (v *ConcurrentList) MoveBefore(e, mark *list.Element) {
	//v.Lock()
	//defer v.Unlock()

	v.list.MoveBefore(e, mark)
}

func (v *ConcurrentList) MoveAfter(e, mark *list.Element) {
	//v.Lock()
	//defer v.Unlock()

	v.list.MoveBefore(e, mark)
}

func (v *ConcurrentList) PushBackList(other *list.List) {
	//v.Lock()
	//defer v.Unlock()

	v.list.PushBackList(other)
}

func (v *ConcurrentList) PushFrontList(other *list.List) {
	//v.Lock()
	//defer v.Unlock()

	v.list.PushBackList(other)
}
