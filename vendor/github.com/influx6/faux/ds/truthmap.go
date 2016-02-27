package ds

import "sync"

// TruthMap provides a mutex controlled map
type TruthMap struct {
	rw sync.RWMutex
	c  TruthTable
}

//NewTruthMap returns a new collector instance
func NewTruthMap(m TruthTable) TruthTable {
	so := TruthMap{c: m}
	return &so
}

//Clone makes a new clone of this collector
func (c *TruthMap) Clone() TruthTable {
	var co TruthTable

	c.rw.RLock()
	co = c.c.Clone()
	c.rw.RUnlock()

	return NewTruthMap(co)
}

//Remove deletes a key:value pair
func (c *TruthMap) Remove(k string) {
	c.rw.Lock()
	c.c.Remove(k)
	c.rw.Unlock()
}

//Set puts a specific key:value into the collector
func (c *TruthMap) Set(k string) {
	c.rw.Lock()
	c.c.Set(k)
	c.rw.Unlock()
}

//Copy copies the map into the collector
func (c *TruthMap) Copy(m map[string]bool) {
	c.rw.Lock()
	for v, k := range m {
		if k {
			c.c.Set(v)
		}
	}
	c.rw.Unlock()
}

//Each iterates through all items in the collector
func (c *TruthMap) Each(fx BoolFunc) {
	c.rw.RLock()
	c.c.Each(fx)
	c.rw.RUnlock()
}

//Keys return the keys of the Collector
func (c *TruthMap) Keys() []string {
	var keys []string
	c.rw.RLock()
	keys = c.c.Keys()
	c.rw.RUnlock()
	return keys
}

//Has returns if a key exists
func (c *TruthMap) Has(k string) bool {
	var ok bool
	c.rw.RLock()
	ok = c.c.Has(k)
	c.rw.RUnlock()
	return ok
}

//Clear clears the collector
func (c *TruthMap) Clear() {
	c.rw.Lock()
	c.c.Clear()
	c.rw.Unlock()
}
