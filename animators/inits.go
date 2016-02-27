package animators

import "github.com/influx6/govfx"

// init registers all the available animators, so users can take advantage of
// the new initialization API.
func init() {

	govfx.RegisterSequence("height", Height{})
	govfx.RegisterSequence("width", Width{})
	govfx.RegisterSequence("translate-x", TranslateX{})
	govfx.RegisterSequence("translate-y", TranslateY{})

}
