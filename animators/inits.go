package animators

import "github.com/influx6/govfx"

func init() {

	govfx.RegisterSequence("translate-x", TranslateX{})
	govfx.RegisterSequence("translate-y", TranslateY{})

}
