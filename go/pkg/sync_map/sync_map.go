package sync_map

import "sync"

type SyncMap struct {
	val  map[string]struct{}
	lock *sync.Mutex
}

func New() *SyncMap {
	m := make(map[string]struct{}, 0)
	return &SyncMap{
		val:  m,
		lock: &sync.Mutex{},
	}
}
func (m *SyncMap) Get(key string) (struct{}, error) {
	m.lock.Lock()
	val := m.val[key]
	m.lock.Unlock()
	return val, nil
}

func (m *SyncMap) Set(key string, val struct{}) error {
	m.lock.Lock()
	m.val[key] = val
	m.lock.Unlock()
	return nil
}

func (m *SyncMap) Size() int {
	m.lock.Lock()
	size := len(m.val)
	m.lock.Unlock()
	return size
}

func (m *SyncMap) Delete(key string) {
	delete(m.val, key)
}

func (m *SyncMap) DeepCopy() map[string]struct{} {
	m.lock.Lock()
	temp := m.val
	m.lock.Unlock()
	return temp
}
