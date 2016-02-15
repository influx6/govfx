package vfx

import (
	"errors"
	"regexp"
	"strings"

	"github.com/gopherjs/gopherjs/js"
)

//==============================================================================

// ErrNotFound provides a not found error used when a property was not found.
var ErrNotFound = errors.New("Not Found")

//==============================================================================

// GetProp retrieves the necessary property for this specific name.
func GetProp(o *js.Object, prop string) (*js.Object, error) {

	// Expand the property for possible period delimited sets.
	props := Expando(prop)

	var jsop *js.Object

	// Loop the property sets and get the next one
	for i := 0; i < len(props); i++ {
		jsop = o.Get(prop)

		if jsop == nil {
			return nil, ErrNotFound
		}
	}

	return jsop, nil
}

//==============================================================================

// MatchProp matches the string value from a property val against a provided
// expected value.
func MatchProp(o *js.Object, prop string, val string) bool {
	op, err := GetProp(o, prop)
	if err != nil {
		return false
	}

	if strings.ToLower(op.String()) != val {
		return false
	}

	return true
}

//==============================================================================
// expandable defines a regexp for matching period delimited strings.
var expandable = regexp.MustCompile("([\\w\\d_-]+\\.[\\w\\d_-]+)+")

// Expando expands a property period delimited string into its component parts.
func Expando(prop string) []string {
	if !expandable.MatchString(prop) {
		return []string{prop}
	}

	return strings.Split(prop, ".")
}

//==============================================================================
