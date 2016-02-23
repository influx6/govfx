package ds

// BoolStore defines a map of possible truthy keys, if a key does not exists,
// then the fact for that key is false.
type BoolStore map[string]bool

// NewBoolStore returns a new BoolStore instance.
func NewBoolStore() BoolStore {
	return make(BoolStore)
}

// Clone makes a new clone of this BoolStore.
func (c BoolStore) Clone() TruthTable {
	col := make(BoolStore)
	col.Copy(c)
	return col
}

// Remove deletes a key from the truthy table.
func (c BoolStore) Remove(k string) {
	if c.Has(k) {
		delete(c, k)
	}
}

// Keys return the keys of the BoolStore.
func (c BoolStore) Keys() []string {
	var keys []string
	c.Each(func(k string, _ func()) {
		keys = append(keys, k)
	})
	return keys
}

// Has returns if a key exists.
func (c BoolStore) Has(k string) bool {
	_, ok := c[k]
	return ok
}

// Set puts a specific key into the BoolStore
func (c BoolStore) Set(k string) {
	c[k] = true
}

// Copy copies the map into the BoolStore
func (c BoolStore) Copy(m map[string]bool) {
	for v, k := range m {
		if k {
			c.Set(v)
		}
	}
}

// BoolFunc defines the type of the Mappable.Each rule
type BoolFunc func(string, func())

// Each iterates through all items in the BoolStore
func (c BoolStore) Each(fx BoolFunc) {
	var state bool
	for k := range c {
		if state {
			break
		}

		fx(k, func() {
			state = true
		})
	}
}

// Clear clears the BoolStore
func (c BoolStore) Clear() {
	for k := range c {
		delete(c, k)
	}
}
