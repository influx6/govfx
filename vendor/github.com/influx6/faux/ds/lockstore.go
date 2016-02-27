package ds

import "sync"

// LockStore provides a mutex controlled map
type LockStore struct {
	rw sync.RWMutex
	c  Stores
}

// NewLockStore returns a new collector instance
func NewLockStore(m Stores) Stores {
	so := LockStore{c: m}
	return &so
}

// Clone makes a new clone of this collector
func (c *LockStore) Clone() Stores {
	var co Stores

	c.rw.RLock()
	co = c.c.Clone()
	c.rw.RUnlock()

	return NewLockStore(co)
}

// Remove deletes a key:value pair
func (c *LockStore) Remove(k string) {
	c.rw.Lock()
	c.c.Remove(k)
	c.rw.Unlock()
}

// Set puts a specific key:value into the collector
func (c *LockStore) Set(k string, v string) {
	c.rw.Lock()
	c.c.Set(k, v)
	c.rw.Unlock()
}

// Copy copies the map into the collector
func (c *LockStore) Copy(m map[string]string) {
	c.rw.Lock()
	for v, k := range m {
		c.c.Set(k, v)
	}
	c.rw.Unlock()
}

// Each iterates through all items in the collector
func (c *LockStore) Each(fx StoreFunc) {
	c.rw.RLock()
	c.c.Each(fx)
	c.rw.RUnlock()
}

// Keys return the keys of the Collector
func (c *LockStore) Keys() []string {
	var keys []string
	c.rw.RLock()
	keys = c.c.Keys()
	c.rw.RUnlock()
	return keys
}

// Get returns the value with the key
func (c *LockStore) Get(k string) string {
	var v string
	c.rw.RLock()
	v = c.c.Get(k)
	c.rw.RUnlock()
	return v
}

// Has returns if a key exists
func (c *LockStore) Has(k string) bool {
	var ok bool
	c.rw.RLock()
	ok = c.c.Has(k)
	c.rw.RUnlock()
	return ok
}

// HasMatch checks if key and value exists and are matching
func (c *LockStore) HasMatch(k string, v string) bool {
	var ok bool
	c.rw.RLock()
	ok = c.c.HasMatch(k, v)
	c.rw.RUnlock()
	return ok
}

// Clear clears the collector
func (c *LockStore) Clear() {
	c.rw.Lock()
	c.c.Clear()
	c.rw.Unlock()
}
