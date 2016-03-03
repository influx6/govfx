package raf

import (
	"math"

	"github.com/go-humble/detect"
	"github.com/gopherjs/gopherjs/js"
)

//==============================================================================

// Mux defines a handler for using with RAF.
type Mux func(float64)

// RequestAnimationFrame provides a cover for RAF using the js
// api for requestAnimationFrame.
func RequestAnimationFrame(r Mux) int {
	return js.Global.Call("requestAnimationFrame", r).Int()
}

// CancelAnimationFrame provides a cover for RAF using the
// js api cancelAnimationFrame.
func CancelAnimationFrame(id int, f ...func()) {
	js.Global.Call("cancelAnimationFrame", id)
	for _, fx := range f {
		fx()
	}
}

//==============================================================================

func init() {
	if detect.IsBrowser() {
		rafPolyfill()
	}
}

func rafPolyfill() {
	window := js.Global
	vendors := []string{"ms", "moz", "webkit", "o"}
	if window.Get("requestAnimationFrame") == nil {
		for i := 0; i < len(vendors) && window.Get("requestAnimationFrame") == nil; i++ {
			vendor := vendors[i]
			window.Set("requestAnimationFrame", window.Get(vendor+"RequestAnimationFrame"))
			window.Set("cancelAnimationFrame", window.Get(vendor+"CancelAnimationFrame"))
			if window.Get("cancelAnimationFrame") == nil {
				window.Set("cancelAnimationFrame", window.Get(vendor+"CancelRequestAnimationFrame"))
			}
		}
	}

	lastTime := 0.0
	if window.Get("requestAnimationFrame") == nil {
		window.Set("requestAnimationFrame", func(callback func(float64)) int {
			currTime := js.Global.Get("Date").New().Call("getTime").Float()
			timeToCall := math.Max(0, 16-(currTime-lastTime))
			id := window.Call("setTimeout", func() { callback(float64(currTime + timeToCall)) }, timeToCall)
			lastTime = currTime + timeToCall
			return id.Int()
		})
	}

	if window.Get("cancelAnimationFrame") == nil {
		window.Set("cancelAnimationFrame", func(id int) {
			js.Global.Get("clearTimeout").Invoke(id)
		})
	}
}

//==============================================================================
