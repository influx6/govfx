package vfx

import "time"

//==============================================================================

// FrameController provides an internal controller for Frames to allow
// controlling of a signal of the beginnning and finishing state of their
// sequence writers. This helps to control the pace of which frames release
// writers for the next sequence.
type frameController interface {
	Frame
	BeginWriting()
	DoneWriting()
}

//==============================================================================

// delayedWriter provides a custom writer that forces a delay in its
// write method. This is useful for inserting good delays around
// animation sequences.
type delayedWriter struct {
	ms time.Duration
	f  frameController
}

// Write calls time.After to perform the necessary delay of operation.
func (d *delayedWriter) Write() {
	if !d.f.Stats().IsFirstDone() {
		<-time.After(d.ms)
	}
}

//==============================================================================

// frameBeginWriter is used to indicate to a frame when its writers have
// begun processing.
type frameBeginWriter struct {
	f frameController
}

func (f *frameBeginWriter) Write() {
	f.f.BeginWriting()
}

// frameEndWriter is used to indicate to a frame when its writers have
// ended processing.
type frameEndWriter struct {
	f frameController
}

// Write calls the DoneWriting method to indicate to a frame, that its sequence
// have finished writing.
func (f *frameEndWriter) Write() {
	f.f.DoneWriting()
}

//==============================================================================
