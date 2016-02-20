package govfx

import (
	"fmt"
	"strings"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

//==============================================================================

// GetComputedStyle returns the dom.Element computed css styles.
func GetComputedStyle(elem dom.Element, ps string) (*dom.CSSStyleDeclaration, error) {
	css := Window().GetComputedStyle(elem, ps)
	if css == nil {
		return nil, ErrNotFound
	}

	return css, nil
}

// RemoveComputedStyleValue removes the value of the property from the computed
// style object.
func RemoveComputedStyleValue(css *dom.CSSStyleDeclaration, prop string) {
	defer func() {
		recover()
	}()

	css.Call("removeProperty", prop)
}

// GetComputedStyleValue retrieves the value of the property from the computed
// style object.
func GetComputedStyleValue(elem dom.Element, psudo string, prop string) (*js.Object, error) {
	vs, err := GetComputedStyle(elem, psudo)
	if err != nil {
		return nil, err
	}

	vcs, err := GetComputedStyleValueWith(vs, prop)
	if err != nil {
		return nil, err
	}

	return vcs, nil
}

// GetComputedStyleValueWith usings the CSSStyleDeclaration to
// retrieves the value of the property from the computed
// style object.
func GetComputedStyleValueWith(css *dom.CSSStyleDeclaration, prop string) (*js.Object, error) {
	vs := css.Call("getPropertyValue", prop)
	if vs == nil {
		return nil, ErrNotFound
	}

	return vs, nil
}

// GetComputedStylePriority retrieves the proritiy of the property from the computed
// style object.
func GetComputedStylePriority(css *dom.CSSStyleDeclaration, prop string) (int, error) {
	vs := css.Call("getPropertyPriority", prop)
	if vs == nil {
		return 0, ErrNotFound
	}

	if strings.TrimSpace(vs.String()) == "" {
		return 0, nil
	}

	return 1, nil
}

//==============================================================================

// ComputedStyle defines a style property item.
type ComputedStyle struct {
	Name     string
	Value    string
	Priority bool // values between [0,1] to indicate use of '!important'
}

// ComputedStyleMap defines a map type of computed style properties and values.
type ComputedStyleMap map[string]*ComputedStyle

// GetComputedStyleMap returns a map of computed style properties and values.
func GetComputedStyleMap(elem dom.Element, ps string) (ComputedStyleMap, error) {
	css, err := GetComputedStyle(elem, ps)
	if err != nil {
		return nil, err
	}

	styleMap := make(ComputedStyleMap)

	// Get the map and pull the necessary property:value and importance facts.
	for key, val := range css.ToMap() {
		priority, _ := GetComputedStylePriority(css, key)
		styleMap[key] = &ComputedStyle{
			Name:     key,
			Value:    val,
			Priority: (priority > 0),
		}
	}

	return styleMap, nil
}

// Add adjusts the stylemap with a new property.
func (c ComputedStyleMap) Add(name string, value string, priority bool) {
	c[name] = &ComputedStyle{Name: name, Value: value, Priority: priority}
}

// Has returns true/false if the property exists.
func (c ComputedStyleMap) Has(name string) bool {
	_, ok := c[name]
	return ok
}

// Get retrieves the specific property if it exists.
func (c ComputedStyleMap) Get(name string) (*ComputedStyle, error) {
	cs, ok := c[name]
	if !ok {
		return nil, ErrNotFound
	}

	return cs, nil
}

//==============================================================================

// ToRGB turns a hexademicmal color into rgba format.
// Returns the read, green and blue values as int.
func ToRGB(hex string) (red, green, blue int) {
	if strings.HasPrefix(hex, "#") {
		hex = strings.TrimPrefix(hex, "#")
	}

	// We are dealing with a 3 string hex.
	if len(hex) < 6 {
		parts := strings.Split(hex, "")
		red = parseIntBase16(doubleString(parts[0]))
		green = parseIntBase16(doubleString(parts[1]))
		blue = parseIntBase16(doubleString(parts[2]))
		return
	}

	red = parseIntBase16(hex[0:2])
	green = parseIntBase16(hex[2:4])
	blue = parseIntBase16(hex[4:6])

	return
}

// RGBA turns a hexademicmal color into rgba format.
// Alpha values ranges from 0-100
func RGBA(hex string, alpha int) string {
	r, g, b := ToRGB(hex)
	return fmt.Sprintf("rgba(%d,%d,%d,%.2f)", r, g, b, float64(alpha)/100)
}

// Unit returns a valid unit type in the browser, if the supplied unit is
// standard then it is return else 'px' is returned as default.
func Unit(u string) string {
	switch u {
	case "rem":
		return u
	case "em":
		return u
	case "px":
		return u
	case "%":
		return u
	case "vw":
		return u
	default:
		return "px"
	}
}

// doubleString doubles the giving string.
func doubleString(c string) string {
	return fmt.Sprintf("%s%s", c, c)
}

//==============================================================================
