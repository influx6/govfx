package govfx

import "honnef.co/go/js/dom"

// ShadowRoot provides a DocumentFragment matching interface for a ShadowRoot
// of an element.
type ShadowRoot struct {
	dom.DocumentFragment
	parent dom.Node
}

// NewShadowRoot will return a struct interfacing the shadowRoot else panics if
// the provided node has no shadowRoot.
func NewShadowRoot(node dom.Node) *ShadowRoot {
	root, ok := GetShadowRoot(node)

	if !ok {
		panic("No shadowRoot")
	}

	sr := ShadowRoot{
		DocumentFragment: root,
		parent:           node,
	}

	return &sr
}

// Parent returns the parent for this shadowRoot.
func (s *ShadowRoot) Parent() dom.Node {
	return s.parent
}
