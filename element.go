package govfx

import (
	"io"
	"regexp"

	"honnef.co/go/js/dom"
)

//==============================================================================

// CSSElem defines a element/structure/object which produces css property strings
// as its output.
type CSSElem interface {
	CSS(io.Writer)
}

//==============================================================================

// Elementals defines a lists of elementals,
type Elementals []*Element

// Element defines a structure that holds ehances the dom.Element api.
// Element provides a caching facility that helps to reduce layout checks
// and improve animation by returning last used data. Also it provides
// an appropriate method to update element properties apart from usings
// inlined styles.
type Element struct {
	dom.Element
	props []Sequence
	css   ComputedStyleMap // css holds the map of computed styles.
}

// NewElement returns an instancee of the Element struct.
func NewElement(elem dom.Element, pseudo string) *Element {
	css, err := GetComputedStyleMap(elem, pseudo)
	if err != nil {
		panic(err)
	}

	em := Element{
		css:     css,
		Element: elem,
	}

	return &em
}

// Add adds the given set of CSSElem objects into the element prop list.
func (e *Element) Add(css ...Sequence) {
	e.props = append(e.props, css...)
}

// Init calls the Init() methods on all items in its property list.
func (e *Element) Init() {
	for _, elem := range e.props {
		elem.Init(e)
	}
}

// Reset resets the resetable sequences within the elements prop list.
func (e *Element) Reset() {
	for _, elem := range e.props {
		if bem, ok := elem.(Resetable); ok {
			bem.Reset()
		}
	}
}

// Blend calls the internal Blend functions of the sequence list.
func (e *Element) Blend(d float64) {
	for _, prop := range e.props {
		if bem, ok := prop.(Blending); ok {
			bem.Blend(d)
		}
	}
}

// Update calls the internal Update functions of the sequence list.
func (e *Element) Update(d float64, timeline float64) {
	for _, prop := range e.props {
		prop.Update(d, timeline)
	}
}

// Clear empties the css sequence list for the element.
func (e *Element) Clear() {
	e.props = nil
}

// CSS collects all the internal css data to be writting and writes it out to the
// passed writer.
func (e *Element) CSS(w io.Writer) {
	for _, elem := range e.props {
		elem.CSS(w)
	}
}

var propName = regexp.MustCompile("([\\w\\-0-9]+)\\(?\\)?")

// Read reads out the elements internal css property rule and returns its
// values list and priority(whether it has !important attached).
// If the property does not exists a false value is returned.
func (e *Element) Read(prop string, selector string) (string, bool, bool) {
	cs, err := e.css.Get(prop)
	if err != nil {
		return "", false, false
	}

	for _, val := range cs.Values {
		valName := propName.FindStringSubmatch(val)[1]
		if valName != selector {
			continue
		}

		return val, cs.Priority, true
	}

	// Read the value, return both value and true state.
	return cs.Value, cs.Priority, false
}

// ReadInt reads the given property and attempts to convert its value into a
// int type else returns 0 as that value type.
func (e *Element) ReadInt(prop string, sel string) (int, bool, bool) {
	val, po, ok := e.Read(prop, sel)
	return ParseInt(val), po, ok
}

// ReadFloat reads the given property and attempts to convert its value into a
// float64 type else returns 0 as that value type.
func (e *Element) ReadFloat(prop string, sel string) (float64, bool, bool) {
	val, po, ok := e.Read(prop, sel)
	return ParseFloat(val), po, ok
}
