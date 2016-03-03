package ds

// AnyMap defines a typ of map string
type AnyMap map[string]interface{}

// NewAnyMap returns a new AnyMap instance
func NewAnyMap() AnyMap {
	return make(AnyMap)
}

// Clone makes a new clone of this AnyMap
func (c AnyMap) Clone() Maps {
	col := make(AnyMap)
	col.Copy(c)
	return col
}

// Remove deletes a key:value pair
func (c AnyMap) Remove(k string) {
	if c.Has(k) {
		delete(c, k)
	}
}

// Keys return the keys of the AnyMap
func (c AnyMap) Keys() []string {
	var keys []string
	c.Each(func(_ interface{}, k string, _ func()) {
		keys = append(keys, k)
	})
	return keys
}

// Get returns the value with the key
func (c AnyMap) Get(k string) interface{} {
	return c[k]
}

// Has returns if a key exists
func (c AnyMap) Has(k string) bool {
	_, ok := c[k]
	return ok
}

// HasMatch checks if key and value exists and are matching
func (c AnyMap) HasMatch(k string, v interface{}) bool {
	if c.Has(k) {
		return c.Get(k) == v
	}
	return false
}

// Set puts a specific key:value into the AnyMap
func (c AnyMap) Set(k string, v interface{}) {
	c[k] = v
}

// Copy copies the map into the AnyMap
func (c AnyMap) Copy(m map[string]interface{}) {
	for v, k := range m {
		c.Set(v, k)
	}
}

// AnyMapFunc defines the type of the Mappable.Each rule
type AnyMapFunc func(interface{}, string, func())

// Each iterates through all items in the AnyMap
func (c AnyMap) Each(fx AnyMapFunc) {
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

// Clear clears the AnyMap
func (c AnyMap) Clear() {
	for k := range c {
		delete(c, k)
	}
}
