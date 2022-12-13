package shared

import (
	"github.com/industry-netsecurity-solution/ins-security-channel/ins"
	"sync"
)

type ConcurrentMap interface {
	RemoveAll()
	Remove(key interface{}) (interface{}, bool)
	Len() int
	Range(cb func(interface{}, interface{}) bool)
	GetPair(key interface{}) (ins.Pair, bool)
	SetPair(pair ...ins.Pair)
	ToArray() []ins.Pair
	GetKeys() []interface{}
	GetValues() []interface{}
	Set(key interface{}, value interface{}) (interface{}, bool)
	Get(key interface{}) (interface{}, bool)
	Has(key interface{}) bool
	AsString(key interface{}) (string, error)
	AsInt(key interface{}) (int, error)
	AsInt32(key interface{}) (int32, error)
	AsInt64(key interface{}) (int64, error)
	AsUint(key interface{}) (uint, error)
	AsUint32(key interface{}) (uint32, error)
	AsUint64(key interface{}) (uint64, error)
	AsByte(key interface{}) (byte, error)
	AsFloat32(key interface{}) (float32, error)
	AsFloat64(key interface{}) (float64, error)
}

type concurrentMap struct {
	sync.RWMutex
	ins.Map
}

func ReferenceConcurrentMap(d map[interface{}]interface{}) ConcurrentMap {
	o := concurrentMap{}
	o.Map = d
	return &o
}

func NewConcurrentMap() ConcurrentMap {
	o := new(concurrentMap)
	o.Map = make(ins.Map)
	return o
}

func (m *concurrentMap) RemoveAll() {
	m.Lock()
	defer m.Unlock()

	m.Map.RemoveAll()
}

func (m *concurrentMap) Remove(key interface{}) (interface{}, bool) {
	m.Lock()
	defer m.Unlock()

	return m.Map.Remove(key)
}

func (m *concurrentMap) Len() int {
	m.RLock()
	defer m.RUnlock()

	return m.Map.Len()
}

func (m *concurrentMap) Range(cb func(interface{}, interface{}) bool) {
	m.RLock()
	defer m.RUnlock()

	m.Map.Range(cb)
}

func (m *concurrentMap) GetPair(key interface{}) (ins.Pair, bool) {
	m.RLock()
	defer m.RUnlock()

	return m.Map.GetPair(key)
}

func (m *concurrentMap) SetPair(pair ...ins.Pair) {
	m.Lock()
	defer m.Unlock()

	m.Map.SetPair(pair...)
}

func (m *concurrentMap) ToArray() []ins.Pair {
	m.RLock()
	defer m.RUnlock()

	return m.Map.ToArray()
}

func (m *concurrentMap) GetKeys() []interface{} {
	m.RLock()
	defer m.RUnlock()

	return m.Map.GetKeys()
}

func (m *concurrentMap) GetValues() []interface{} {
	m.RLock()
	defer m.RUnlock()

	return m.Map.GetValues()
}

func (m *concurrentMap) Set(key interface{}, value interface{}) (interface{}, bool) {
	m.Lock()
	defer m.Unlock()

	return m.Map.Set(key, value)
}

func (m *concurrentMap) Get(key interface{}) (interface{}, bool) {
	m.RLock()
	defer m.RUnlock()

	value, ok := m.Map[key]
	return value, ok
}

func (m *concurrentMap) Has(key interface{}) bool {
	m.RLock()
	defer m.RUnlock()

	return m.Map.Has(key)
}

func (m *concurrentMap) AsString(key interface{}) (string, error) {
	m.RLock()
	defer m.RUnlock()

	return m.Map.AsString(key)
}

func (m *concurrentMap) AsInt(key interface{}) (int, error) {
	m.RLock()
	defer m.RUnlock()

	return m.Map.AsInt(key)
}

func (m *concurrentMap) AsInt32(key interface{}) (int32, error) {
	m.RLock()
	defer m.RUnlock()

	return m.Map.AsInt32(key)
}

func (m *concurrentMap) AsInt64(key interface{}) (int64, error) {
	m.RLock()
	defer m.RUnlock()

	return m.Map.AsInt64(key)
}

func (m *concurrentMap) AsUint(key interface{}) (uint, error) {
	m.RLock()
	defer m.RUnlock()

	return m.Map.AsUint(key)
}

func (m *concurrentMap) AsUint32(key interface{}) (uint32, error) {
	m.RLock()
	defer m.RUnlock()

	return m.Map.AsUint32(key)
}

func (m *concurrentMap) AsUint64(key interface{}) (uint64, error) {
	m.RLock()
	defer m.RUnlock()

	return m.Map.AsUint64(key)
}

func (m *concurrentMap) AsByte(key interface{}) (byte, error) {
	m.RLock()
	defer m.RUnlock()

	return m.Map.AsByte(key)
}

func (m *concurrentMap) AsFloat32(key interface{}) (float32, error) {
	m.RLock()
	defer m.RUnlock()

	return m.Map.AsFloat32(key)
}

func (m *concurrentMap) AsFloat64(key interface{}) (float64, error) {
	m.RLock()
	defer m.RUnlock()

	return m.Map.AsFloat64(key)
}

func CopyMap(dst *concurrentMap, src *concurrentMap) error {
	dst.RLock()
	defer dst.RUnlock()

	src.RLock()
	defer src.RUnlock()

	return ins.CopyMap(dst.Map, src.Map)
}
