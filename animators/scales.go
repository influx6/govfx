package animators

//==============================================================================

// // ScaleY defines a sequence for animating css Scale y-axes properties.
// type ScaleY struct {
// 	Value  float64 `govfx:"value"`
// 	Easing string  `govfx:"easing"`
// }
//
// // Init returns the initial writers for the sequence.
// func (t *ScaleY) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
// 	var writers govfx.DeferWriters
//
// 	for _, elem := range elems {
// 		func(e govfx.Elemental) {
// 			writers = append(writers, govfx.NewWriter(func() {
// 				e.Write("transform", "scaleY(1.0)", false)
// 				e.Sync()
// 			}))
// 		}(elem)
// 	}
//
// 	return writers
// }
//
// // Next returns the writers for the next sequence.
// func (t *ScaleY) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
// 	var writers govfx.DeferWriters
//
// 	easing := govfx.GetEasing(t.Easing)
//
// 	for _, elem := range elems {
// 		transform, priority, _ := elem.Read("transform", "scale")
//
// 		func(e govfx.Elemental) {
//
// 			var y float64
//
// 			if govfx.IsMatrix(transform) {
// 				mx, _ := govfx.ToMatrix2D(transform)
// 				y = mx.ScaleY
// 			} else if govfx.IsScale(transform) {
// 				mx, _ := govfx.ToScale(transform)
// 				y = mx.Y
// 			}
//
// 			yd := float64(t.Value) - y
//
// 			yn := easing.Ease(govfx.EaseConfig{
// 				Stat:         stats,
// 				CurrentValue: y,
// 				DeltaValue:   yd,
// 			})
//
// 			transform = fmt.Sprintf("scaleY(%.2f)", yn)
// 			e.EraseMore("transform", "matrix", false)
//
// 			writers = append(writers, govfx.NewWriter(func() {
// 				e.Write("transform", transform, priority)
// 				e.Sync()
// 			}))
// 		}(elem)
// 	}
//
// 	return writers
// }
//
// //==============================================================================
//
// // ScaleX defines a sequence for animating css Scale x-axes properties.
// type ScaleX struct {
// 	Value  float64 `govfx:"value"`
// 	Easing string  `govfx:"easing"`
// }
//
// // Init returns the initial writers for the sequence.
// func (t *ScaleX) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
// 	var writers govfx.DeferWriters
//
// 	for _, elem := range elems {
// 		func(e govfx.Elemental) {
// 			writers = append(writers, govfx.NewWriter(func() {
// 				e.Write("transform", "scaleX(1.0)", false)
// 				e.Sync()
// 			}))
// 		}(elem)
// 	}
//
// 	return writers
// }
//
// // Next returns the writers for the next sequence.
// func (t *ScaleX) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
// 	var writers govfx.DeferWriters
//
// 	easing := govfx.GetEasing(t.Easing)
//
// 	for _, elem := range elems {
// 		transform, priority, _ := elem.Read("transform", "scale")
//
// 		func(e govfx.Elemental) {
//
// 			var x float64
//
// 			if govfx.IsMatrix(transform) {
// 				mx, _ := govfx.ToMatrix2D(transform)
// 				x = mx.ScaleX
// 			} else if govfx.IsScale(transform) {
// 				mx, _ := govfx.ToScale(transform)
// 				x = mx.X
// 			}
//
// 			xd := float64(t.Value) - x
//
// 			xn := easing.Ease(govfx.EaseConfig{
// 				Stat:         stats,
// 				CurrentValue: x,
// 				DeltaValue:   xd,
// 			})
//
// 			e.EraseMore("transform", "matrix", false)
//
// 			transform = fmt.Sprintf("scaleX(%.2f)", xn)
// 			writers = append(writers, govfx.NewWriter(func() {
// 				e.WriteMore("transform", transform, priority)
// 				e.Sync()
// 			}))
// 		}(elem)
// 	}
//
// 	return writers
// }

//==============================================================================
