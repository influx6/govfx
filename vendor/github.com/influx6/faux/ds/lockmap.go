package ds

import "sync"

// LockMap provides a mutex controlled map
type LockMap struct {
	rw sync.RWMutex
	c  Maps
}

//NewLockMap returns a new collector instance
func NewLockMap(m Maps) Maps {
	so := LockMap{c: m}
	return &so
}

//Clone makes a new clone of this collector
func (c *LockMap) Clone() Maps {
	var co Maps

	c.rw.RLock()
	co = c.c.Clone()
	c.rw.RUnlock()

	return NewLockMap(co)
}

//Remove deletes a key:value pair
func (c *LockMap) Remove(k string) {
	c.rw.Lock()
	c.c.Remove(k)
	c.rw.Unlock()
}

//Set puts a specific key:value into the collector
func (c *LockMap) Set(k string, v interface{}) {
	c.rw.Lock()
	c.c.Set(k, v)
	c.rw.Unlock()
}

//Copy copies the map into the collector
func (c *LockMap) Copy(m map[string]interface{}) {
	c.rw.Lock()
	for k, v := range m {
		c.c.Set(k, v)
	}
	c.rw.Unlock()
}

//Each iterates through all items in the collector
func (c *LockMap) Each(fx AnyMapFunc) {
	c.rw.RLock()
	c.c.Each(fx)
	c.rw.RUnlock()
}

//Keys return the keys of the Collector
func (c *LockMap) Keys() []string {
	var keys []string
	c.rw.RLock()
	keys = c.c.Keys()
	c.rw.RUnlock()
	return keys
}

//Get returns the value with the key
func (c *LockMap) Get(k string) interface{} {
	var v interface{}
	c.rw.RLock()
	v = c.c.Get(k)
	c.rw.RUnlock()
	return v
}

//Has returns if a key exists
func (c *LockMap) Has(k string) bool {
	var ok bool
	c.rw.RLock()
	ok = c.c.Has(k)
	c.rw.RUnlock()
	return ok
}

//HasMatch checks if key and value exists and are matching
func (c *LockMap) HasMatch(k string, v interface{}) bool {
	var ok bool
	c.rw.RLock()
	ok = c.c.HasMatch(k, v)
	c.rw.RUnlock()
	return ok
}

//Clear clears the collector
func (c *LockMap) Clear() {
	c.rw.Lock()
	c.c.Clear()
	c.rw.Unlock()
}
