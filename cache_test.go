package vfx_test

import (
	"sync"
	"testing"
	"time"

	"github.com/ardanlabs/kit/tests"
	"github.com/influx6/faux/vfx"
)

// TestDeferWriterCache validates the operation and write safetiness of DeferWriterCache.
func TestDeferWriterCache(t *testing.T) {
	tests.ResetLog()
	defer tests.DisplayLog()

	t.Logf("Given the need to use a DeferWriterCache")
	{

		t.Logf("\tWhen giving two stat keys")
		{

			var ws sync.WaitGroup

			ws.Add(1)

			cache := vfx.NewDeferWriterCache()

			stat := vfx.TimeStat(vfx.StatConfig{
				Duration: 1 * time.Second,
				Delay:    2 * time.Second,
				Easing:   "ease-in",
				Loop:     0,
				Reverse:  false,
				Optimize: false,
			})

			stat2 := vfx.TimeStat(vfx.StatConfig{
				Duration: 1 * time.Second,
				Delay:    2 * time.Second,
				Easing:   "ease-in",
				Loop:     0,
				Reverse:  false,
				Optimize: false,
			})

			frame := vfx.NewAnimationSequence("", stat)
			frame2 := vfx.NewAnimationSequence("", stat2)

			defer cache.Clear(frame)
			defer cache.Clear(frame2)

			go func() {
				defer ws.Done()
				for i := 0; i < stat.TotalIterations(); i++ {
					j := stat.CurrentIteration()
					cache.Store(frame, j, buildNDeferWriters(i+1)...)
					stat.Next(0)
				}
			}()

			ws.Wait()
			ws.Add(2)

			go func() {
				defer ws.Done()
				for i := 0; i < 10; i++ {
					wd := cache.Writers(frame, i)
					if len(wd) != (i + 1) {
						t.Fatalf("\t%s\tShould have writer lists for step: %d of len: %d", tests.Failed, i, i*4)
					}
					t.Logf("\t%s\tShould have writer lists for step: %d of len: %d", tests.Success, i, i*4)
				}
			}()

			go func() {
				defer ws.Done()
				for i := 0; i < stat2.TotalIterations(); i++ {
					j := stat2.CurrentIteration()
					cache.Store(frame2, j, buildNDeferWriters(i+1)...)
					stat2.Next(0)
				}
			}()

			ws.Wait()

			w4 := cache.Writers(frame, 4)
			cache.ClearIteration(frame, 4)
			w42 := cache.Writers(frame, 4)

			if len(w42) >= len(w4) {
				t.Fatalf("\t%s\tShould have cleared writers of 4 iteration", tests.Failed)
			}
			t.Logf("\t%s\tShould have cleared writers of 4 iteration", tests.Success)

		}
	}
}

//==============================================================================

// wr implements vfx.DeferWriter interface.
type wr struct{}

// Write writes out the writers details.
func (w *wr) Write() {}

func buildNDeferWriters(size int) vfx.DeferWriters {
	var ws vfx.DeferWriters

	for i := 0; i < size; i++ {
		ws = append(ws, &wr{})
	}

	return ws
}

//==============================================================================
