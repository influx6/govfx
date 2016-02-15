package vfx

//==============================================================================

// Animate provides the central engine for managing all animation calls.
// Animate uses writer batching to reduce layout trashing. Hence  each frame
// assigned for each animation call, will have all their writes batched
// into one call.
func Animate(frame Frame) {

	// Due to the fact we cant reset due to the fact that the next call of the
	// game loop could still have the frame registered to run, we cant reset,
	// within the game loop, hence defer a reset if the frame is to be re-added
	// into the animation loop again.
	if frame.IsOver() {
		frame.Reset()
	}

	// Return this frame subscription ender, initialized and run its writers.
	stopCache.Add(frame, engine.Loop(func(delta float64) {
		var writers DeferWriters

		if frame.IsOver() {
			wcache.Clear(frame)

			// Stop this frame for being executed anymore.
			Stop(frame)

			// Reset the frame for re-use.
			// frame.Reset()

			return
		}

		if !frame.Inited() {

			writers = frame.Init(delta)
			frame.Sync()

		} else {

			writers = frame.Sequence(delta)
			frame.Sync()

		}

		// Incase we end up using delays with our sequence, GopherJS can
		// not block and should not block, other processes, so lunch the
		// writers in a Goroutine. Frames have built in reconciliation system
		// to manage the variances when dealing with delays.
		go func() {
			for _, w := range writers {
				w.Write()
			}
		}()
	}, 0))
}

// Stop stops the frame within the animation step, removing its registered
// loopere.
func Stop(frame Frame) {
	looper := stopCache.Get(frame)
	if looper != nil {
		stopCache.Delete(frame)
		looper.End()
	}
}

// RegisterEasing adds a easing provider into the registery with the specified
// name, we allow replacing a easing provider for a keyed name, if you so wish.
func RegisterEasing(name string, easing Easing) {
	easingProviders.Add(name, easing)
}

//==============================================================================
