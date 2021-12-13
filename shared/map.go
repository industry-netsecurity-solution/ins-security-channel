package shared

import (
	"sync"
)

type Pair struct {
	Key interface{}
	Value interface{}
}

type ConcurrentMap struct {
	sync.RWMutex
	dic     map[interface{}]interface{}
}

func NewConcurrentMap() *ConcurrentMap {
	o := new(ConcurrentMap)
	o.dic = make(map[interface{}]interface{})
	return o
}

func (m *ConcurrentMap) RemoveAll() {
	m.Lock()
	defer m.Unlock()

	for k := range m.dic {
		delete(m.dic, k)
	}
}

func (m *ConcurrentMap) Remove(key interface{}) (interface{}, bool) {
	m.Lock()
	defer m.Unlock()

	var value interface{} = nil
	var ok bool
	if value, ok = m.dic[key]; ok {
		delete(m.dic, key)
	}

	return value, ok
}

func (m *ConcurrentMap) Has(key interface{}) bool {
	m.RLock()
	defer m.RUnlock()

	_, ok := m.dic[key]
	return  ok
}

func (m *ConcurrentMap) Get(key interface{}) (interface{}, bool) {
	m.RLock()
	defer m.RUnlock()

	value, ok := m.dic[key]
	return  value, ok
}

func (m *ConcurrentMap) Set(key interface{}, value interface{}) (interface{}, bool) {
	m.Lock()
	defer m.Unlock()

	var old interface{} = nil
	var ok bool
	if old, ok = m.dic[key]; ok {
	}

	m.dic[key] = value

	return old, ok
}

func (m *ConcurrentMap) Len() int {
	m.RLock()
	defer m.RUnlock()

	return len(m.dic)
}

func (m *ConcurrentMap) Range(cb func(interface{}, interface{}) bool ) {
	m.RLock()
	defer m.RUnlock()

	for k, v := range m.dic {
		if cb(k, v) == false {
			return;
		}
	}
}

func (m *ConcurrentMap) GetPair(key interface{}) (Pair, bool) {
	m.RLock()
	defer m.RUnlock()

	value, ok := m.dic[key]
	return Pair{key, value}, ok
}

func (m *ConcurrentMap) SetPair(pair *Pair) (interface{}, bool) {
	m.Lock()
	defer m.Unlock()

	var old interface{} = nil
	var ok bool
	if old, ok = m.dic[pair.Key]; ok {
	}

	m.dic[pair.Key] = pair.Value

	return old, ok
}

func (m *ConcurrentMap) ToArray() []Pair {
	m.RLock()
	defer m.RUnlock()

	size := len(m.dic)
	result := make([]Pair, size)

	i := 0
	for k, v := range m.dic {
		result[i] = Pair{k, v}
	}

	return result
}