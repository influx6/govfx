package animators

import "github.com/influx6/govfx"

// init registers all the available animators, so users can take advantage of
// the new initialization API.
func init() {

	govfx.RegisterSequence("height", Height{})
	govfx.RegisterSequence("width", Width{})
	govfx.RegisterSequence("translate-x", TranslateX{})
	govfx.RegisterSequence("translate-y", TranslateY{})
	govfx.RegisterSequence("scale-x", ScaleX{})
	govfx.RegisterSequence("scale-y", ScaleY{})
	govfx.RegisterSequence("skew-x", SkewX{})
	govfx.RegisterSequence("skew-y", SkewY{})
	govfx.RegisterSequence("rotate", Rotate{})
	govfx.RegisterSequence("rotate-x", RotateX{})
	govfx.RegisterSequence("rotate-y", RotateY{})

	govfx.RegisterSequence("color", Color{})
	govfx.RegisterSequence("background-color", BackgroundColor{})
}
