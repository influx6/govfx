package govfx

import (
	"regexp"
	"strings"

	"github.com/go-humble/detect"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

//==============================================================================

// window stores the current dom window object.
var window dom.Window
var doc dom.Document

// Root returns the global js.Object for the current js context.
func Root() *js.Object {
	return js.Global
}

// Window returns the root window for the current dom.
func Window() dom.Window {
	if window == nil {
		window = dom.GetWindow()
	}
	return window
}

// Document returns the current document attached to the window
func Document() dom.Document {
	if doc == nil {
		doc = Window().Document()
	}
	return doc
}

// QuerySelectorAll returns a lists of elementals that maches the selector
// provided else returns an empty lists.
func QuerySelectorAll(selector string) Elementals {
	var eml Elementals

	items := Document().QuerySelectorAll(selector)

	for _, item := range items {
		eml = append(eml, NewElement(item, ""))
	}

	return eml
}

// QuerySelector returns the elemental that maches the selector else returns
// nil.
func QuerySelector(selector string) Elemental {
	node := Document().QuerySelector(selector)
	if node == nil {
		return nil
	}

	return NewElement(node, "")
}

// TransformElements returns a lists of Elementals if the argument provided
// is either a dom.Element, a list of dom.Element or Elementals, transforming
// them accordingly else returns nil.
func TransformElements(elem interface{}) Elementals {

	switch elem.(type) {
	case Elementals:
		return elem.(Elementals)
	case Element:
		return Elementals{elem.(Elemental)}
	case dom.Element:
		return Elementals{NewElement(elem.(dom.Element), "")}
	case []dom.Element:
		var m Elementals

		for _, item := range elem.([]dom.Element) {
			if item != nil {
				m = append(m, NewElement(item, ""))
			}
		}

		return m
	}

	return nil
}

// RootElement returns the root parent element that for the provided element.
func RootElement(elem dom.Node) dom.Node {
	var parent dom.Node

	parent = elem

	for parent.ParentNode() != nil {
		parent = parent.ParentNode()
	}

	return parent
}

// HasShadowRoot returns true/false if the element type matches
// the [object shadowRoot].
func HasShadowRoot(elem dom.Node) bool {
	rs := RootElement(elem)
	return rs.Underlying().Call("toString").String() == "[object ShadowRoot]"
}

// ShadowRootDocument returns the DocumentFragment interface for the provided
// shadowRoot.
func ShadowRootDocument(elem dom.Node) dom.DocumentFragment {
	rs := RootElement(elem)
	return dom.WrapDocumentFragment(rs.Underlying())
}

// GetShadowRoot retrieves the shadowRoot connected to the pass dom.Node, else
// returns false as the second argument if the node has no shadowRoot.
func GetShadowRoot(elem dom.Node) (dom.DocumentFragment, bool) {
	if elem == nil {
		return nil, false
	}

	var root *js.Object

	if root = elem.Underlying().Get("shadowRoot"); root == nil {
		if root = elem.Underlying().Get("root"); root == nil {
			return nil, false
		}
	}

	return dom.WrapDocumentFragment(root), true
}

//==============================================================================

// topScrollAttr defines the apppropriate property to retrieve the top scroll
// value.
var topScrollAttr string

// leftScrollAttr defines the apppropriate property to retrieve the left scroll
// value.
var leftScrollAttr string

var useDocForOffset bool

// initScrollProperties initializes the necessary scroll property names.
func initScrollProperties() {
	if Root().Get("pageYOffset") != nil {
		topScrollAttr = "pageYOffset"
	} else {
		topScrollAttr = "scrollTop"
		useDocForOffset = true
	}

	if Root().Get("pageXOffset") != nil {
		leftScrollAttr = "pageXOffset"
	} else {
		leftScrollAttr = "scrollLeft"
		useDocForOffset = true
	}
}

//==============================================================================

// PageBox returns the offset of the current page bounding box.
func PageBox() (float64, float64) {
	var cursor *js.Object

	if useDocForOffset {
		cursor = Document().Underlying()
	} else {
		cursor = Root()
	}

	top := cursor.Get(topScrollAttr)
	left := cursor.Get(leftScrollAttr)
	return ParseFloat(top.String()), ParseFloat(left.String())
}

// ClientBox returns the offset of the current page client box.
func ClientBox() (float64, float64) {
	top := Document().Underlying().Get("clientTop")
	left := Document().Underlying().Get("clientLeft")

	if top == nil || left == nil {
		return 0, 0
	}

	return ParseFloat(top.String()), ParseFloat(left.String())
}

// rootName defines a regexp for matching the string to either be body/html.
var rootName = regexp.MustCompile("^(?:body|html)$")

// Position returns the current position of the dom.Element.
func Position(elem dom.Element) (float64, float64) {
	parent := OffsetParent(elem)

	var parentTop, parentLeft float64
	var marginTop, marginLeft float64
	var pBorderTop, pBorderLeft float64
	var pBorderTopObject *js.Object
	var pBorderLeftObject *js.Object

	nodeNameObject, err := GetProp(parent, "nodeName")
	if err == nil && !rootName.MatchString(strings.ToLower(nodeNameObject.String())) {
		parentElem := dom.WrapElement(parent)
		parentTop, parentLeft = Offset(parentElem)
	}

	if parent.Get("style") != nil {

		pBorderTopObject, err = GetProp(parent, "style.borderTopWidth")
		if err == nil {
			pBorderTop = ParseFloat(pBorderTopObject.String())
		}

		pBorderLeftObject, err = GetProp(parent, "style.borderLeftWidth")
		if err == nil {
			pBorderLeft = ParseFloat(pBorderLeftObject.String())
		}

		parentTop += pBorderTop
		parentLeft += pBorderLeft
	}

	css, _ := GetComputedStyle(elem, "")

	marginTopObject, err := GetComputedStyleValueWith(css, "margin-top")
	if err == nil {
		marginTop = ParseFloat(marginTopObject.String())
	}

	marginLeftObject, err := GetComputedStyleValueWith(css, "margin-left")
	if err == nil {
		marginLeft = ParseFloat(marginLeftObject.String())
	}

	elemTop, elemLeft := Offset(elem)

	elemTop -= marginTop
	elemLeft -= marginLeft

	return elemTop - parentTop, elemLeft - parentLeft
}

// Offset returns the top,left offset of a dom.Element.
func Offset(elem dom.Element) (float64, float64) {
	boxTop, _, _, boxLeft := BoundingBox(elem)
	clientTop, clientLeft := ClientBox()
	pageTop, pageLeft := PageBox()

	top := (boxTop + pageTop) - clientTop
	left := (boxLeft + pageLeft) - clientLeft

	return top, left
}

// BoundingBox returns the top,right,down,left corners of a dom.Element.
func BoundingBox(elem dom.Element) (float64, float64, float64, float64) {
	rect := elem.GetBoundingClientRect()
	return rect.Top, rect.Right, rect.Bottom, rect.Left
}

//==============================================================================

// OffsetParent returns the offset parent element for a specific element.
func OffsetParent(elem dom.Element) *js.Object {
	und := elem.Underlying()

	osp, err := GetProp(und, "offsetParent")
	if err != nil {
		osp = Document().Underlying()
	}

	for osp != nil && !MatchProp(osp, "nodeType", "html") && MatchProp(osp, "style.position", "static") {
		val, err := GetProp(osp, "offsetParent")
		if err != nil {
			break
		}

		osp = val
	}

	return osp
}

//==============================================================================

// init initalizes properties and functions necessary for package wide varaibles.
func init() {
	if detect.IsBrowser() {
		initScrollProperties()
	}
}

//==============================================================================
