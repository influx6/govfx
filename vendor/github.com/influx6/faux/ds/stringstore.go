package ds

// StringStore defines a typ of map string
type StringStore map[string]string

// NewStringStore returns a new StringStore instance
func NewStringStore() StringStore {
	return make(StringStore)
}

// Clone makes a new clone of this StringStore
func (c StringStore) Clone() Stores {
	col := make(StringStore)
	col.Copy(c)
	return col
}

// Remove deletes a key:value pair
func (c StringStore) Remove(k string) {
	if c.Has(k) {
		delete(c, k)
	}
}

// Keys return the keys of the StringStore
func (c StringStore) Keys() []string {
	var keys []string
	c.Each(func(_ string, k string, _ func()) {
		keys = append(keys, k)
	})
	return keys
}

// Get returns the value with the key
func (c StringStore) Get(k string) string {
	return c[k]
}

// Has returns if a key exists
func (c StringStore) Has(k string) bool {
	_, ok := c[k]
	return ok
}

// HasMatch checks if key and value exists and are matching
func (c StringStore) HasMatch(k string, v string) bool {
	if c.Has(k) {
		return c.Get(k) == v
	}
	return false
}

// Set puts a specific key:value into the StringStore
func (c StringStore) Set(k string, v string) {
	c[k] = v
}

// Copy copies the map into the StringStore
func (c StringStore) Copy(m map[string]string) {
	for v, k := range m {
		c.Set(v, k)
	}
}

// StoreFunc defines the type of the Mappable.Each rule
type StoreFunc func(string, string, func())

// Each iterates through all items in the StringStore
func (c StringStore) Each(fx StoreFunc) {
	var state bool
	for k, v := range c {
		if state {
			break
		}

		fx(v, k, func() {
			state = true
		})
	}
}

// Clear clears the StringStore
func (c StringStore) Clear() {
	for k := range c {
		delete(c, k)
	}
}
