package animators

import "github.com/influx6/govfx"

func init() {
	govfx.RegisterAnimator("translate-x", func(m govfx.Values) govfx.Sequence {
		return gofx.Merge(TranslateX{}, m)
	})

}
