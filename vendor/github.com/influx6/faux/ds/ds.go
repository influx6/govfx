package ds

// Maps define a set of method rules for maps of the string key types
type Maps interface {
	Clear()
	HasMatch(k string, v interface{}) bool
	Each(f AnyMapFunc)
	Keys() []string
	Copy(map[string]interface{})
	Has(string) bool
	Get(string) interface{}
	Remove(string)
	Set(k string, v interface{})
	Clone() Maps
}

// Stores define a set of method rules for maps of the string key types
type Stores interface {
	Clear()
	HasMatch(k string, v string) bool
	Each(f StoreFunc)
	Keys() []string
	Copy(map[string]string)
	Has(string) bool
	Get(string) string
	Remove(string)
	Set(k string, v string)
	Clone() Stores
}

// TruthTable define a set of method rules truth tables.
type TruthTable interface {
	Clear()
	Each(f BoolFunc)
	Keys() []string
	Copy(map[string]bool)
	Has(string) bool
	Remove(string)
	Set(k string)
	Clone() TruthTable
}

// NewTruthTable returns a new instance of TruthTable.
func NewTruthTable() TruthTable {
	return NewTruthMap(NewBoolStore())
}
