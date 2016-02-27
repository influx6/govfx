package govfx

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/influx6/faux/ds"
	"github.com/influx6/faux/utils"
	"honnef.co/go/js/dom"
)

//==============================================================================

// Elemental defines a dom.Element with read-write abilities for css properties.
// Calling th Sync() method of an elemental adjusts the changes as a batch into
// the real browser dom for that specific dom element.
type Elemental interface {
	dom.Element
	Read(string) (string, bool, bool)
	ReadMore(string) ([]string, bool, bool)
	ReadInt(string) (int, bool, bool)
	ReadFloat(string) (float64, bool, bool)
	Write(string, string, bool)
	WriteMore(string, string, bool)
	Sync()
}

// Elementals defines a lists of elementals,
type Elementals []Elemental

// NewElement returns an instancee of the Element struct.
func NewElement(elem dom.Element, pseudo string) Elemental {
	var shadow dom.DocumentFragment

	if HasShadowRoot(elem) {
		shadow = ShadowRootDocument(elem)
	}

	css, err := GetComputedStyleMap(elem, pseudo)
	if err != nil {
		panic(err)
	}

	var id string

	if eid := elem.GetAttribute("id"); eid != "" {
		id = eid
	} else {
		id = fmt.Sprintf("elemental-%s", utils.RandString(10))
	}

	elem.SetAttribute("id", id)

	em := Element{
		id:      id,
		Element: elem,
		css:     css,
		cssDiff: ds.NewTruthMap(ds.NewBoolStore()),
		style:   NewStyleSync(id, shadow),
	}

	return &em
}

//==============================================================================

// Element defines a structure that holds ehances the dom.Element api.
// Element provides a caching facility that helps to reduce layout checks
// and improve animation by returning last used data. Also it provides
// an appropriate method to update element properties apart from usings
// inlined styles.
type Element struct {
	dom.Element
	id      string        // the dom.Element id if it has one
	cssDiff ds.TruthTable // contains lists of properties that have be change.
	style   *StyleSync    // style used for syncing elemental property changes.
	rl      sync.RWMutex
	css     ComputedStyleMap // css holds the map of computed styles.
}

// ReadMore reads out the elements internal css property rule and returns its
// values list and priority(whether it has !important attached).
// If the property does not exists a false value is returned.
func (e *Element) ReadMore(prop string) ([]string, bool, bool) {
	e.rl.RLock()
	defer e.rl.RUnlock()

	cs, err := e.css.Get(prop)
	if err != nil {
		return nil, false, false
	}

	// Read the value, return both value and true state.
	return cs.Values, cs.Priority, true
}

// Read reads out the elements internal css property rule and returns its
// value and priority(wether it has !important attached).
// If the property does not exists a false value is returned.
func (e *Element) Read(prop string) (string, bool, bool) {
	e.rl.RLock()
	defer e.rl.RUnlock()

	cs, err := e.css.Get(prop)
	if err != nil {
		return "", false, false
	}

	// Read the value, return both value and true state.
	return cs.Value, cs.Priority, true
}

// ReadInt reads the given property and attempts to convert its value into a
// int type else returns 0 as that value type.
func (e *Element) ReadInt(prop string) (int, bool, bool) {
	val, po, ok := e.Read(prop)
	return ParseInt(val), po, ok
}

// ReadFloat reads the given property and attempts to convert its value into a
// float64 type else returns 0 as that value type.
func (e *Element) ReadFloat(prop string) (float64, bool, bool) {
	val, po, ok := e.Read(prop)
	return ParseFloat(val), po, ok
}

// Write adds the necessary change of value to the giving property
// with the necessary adjustments. If the property is not found in
// the elements css stylesheet rules, it will be ignored. Write replaces
// both the value and value lists for a property,setting that property
// as the sole only value. Usefully for a first reset of a multivalue
// property.
func (e *Element) Write(prop string, value string, priority bool) {
	e.rl.RLock()
	cs, err := e.css.Get(prop)
	e.rl.RUnlock()

	if err != nil {
		e.rl.Lock()
		e.css.Add(prop, value, priority)
		e.rl.Unlock()
	}

	cs.Value = value
	cs.Values = []string{value}

	if priority {
		cs.Priority = true
	}

	// Add the property into our diff map to ensure we deal with this
	// efficiently without re-writing untouched rules.
	e.cssDiff.Set(cs.VendorName)
}

// WriteMore allows writing a multi value property, eg transform, which can
// take a scale, translate,etc properties, it allows augmenting the
// property lists rather than replacing it.
func (e *Element) WriteMore(prop string, value string, priority bool) {
	e.rl.RLock()
	cs, err := e.css.Get(prop)
	e.rl.RUnlock()

	if err != nil {
		e.rl.Lock()
		e.css.AddMore(prop, value, priority)
		e.rl.Unlock()
	}

	cs.Values = append(cs.Values, value)

	if priority {
		cs.Priority = true
	}

	// Add the property into our diff map to ensure we deal with this
	// efficiently without re-writing untouched rules.
	e.cssDiff.Set(cs.VendorName)
}

// End removes this element styles from the dom.
func (e *Element) End() {
	e.style.Disconnect()
	e.cssDiff.Clear()
}

// Sync adjusts the necessary property changes of the giving element back into
// the dom. Any changes made to any properties will be diffed and added.
// Sync only re-writes change properties, all untouched onces are left alone.
func (e *Element) Sync() {
	var content bytes.Buffer

	// Collect all information regarding changed properties.
	e.cssDiff.Each(func(key string, stop func()) {
		val, _ := e.css.Get(key)

		// Range over the values lists instead incase we are dealing with
		// multiple assignables.
		for _, item := range val.Values {
			var valueContent string

			if val.Priority {
				valueContent = "\t%s:%s !important;\n"
			} else {
				valueContent = "\t%s:%s;\n"
			}

			fmt.Fprint(&content, fmt.Sprintf(valueContent, key, item))
		}
	})

	e.style.Write(fmt.Sprintf(`
  #%s {
%s
  }
`, e.id, content.String()))

}

//==============================================================================

// StyleSync provides a structure for syncing a giving string into a stylesheet,
// preferrable a style tag associated with it, to update properties of a
// dom.Element without using inline-styles.
type StyleSync struct {
	id   string
	elem dom.Element // the corresponding style sheet used for syncing.
	root dom.Node
}

// NewStyleSync returns a new style syncer.
func NewStyleSync(id string, root dom.Node) *StyleSync {
	sync := StyleSync{
		id:   id,
		elem: Document().CreateElement("style"),
		root: root,
	}

	sync.elem.SetAttribute("id", fmt.Sprintf("%s-styles", id))

	sync.Connect()

	return &sync
}

// Disconnect the style from the head node.
func (s *StyleSync) Disconnect() {
	if s.root == nil {
		Document().QuerySelector("head").RemoveChild(s.elem)
		return
	}
	s.root.RemoveChild(s.elem)
}

// Connect adds the giving StyleSync internal style into the dom.
func (s *StyleSync) Connect() {
	if s.root == nil {
		Document().QuerySelector("head").AppendChild(s.elem)
		return
	}
	s.root.AppendChild(s.elem)
}

// Write re-writes the content of the style with the provided data.
func (s *StyleSync) Write(styleContent string) {
	s.elem.SetInnerHTML(styleContent)
}
