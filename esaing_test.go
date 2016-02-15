package govfx_test

import (
	"testing"

	"github.com/ardanlabs/kit/tests"
	"github.com/influx6/govfx"
)

// TestEasingProvider validates the use of easing registery in storing easing
// providers.
func TestEasingProvider(t *testing.T) {
	tests.ResetLog()
	defer tests.DisplayLog()

	t.Logf("Given the need to use a EasingProvider")
	{
		t.Logf("\tWhen giving a easing registery")
		{

			var ed e1
			var es e2

			easings := govfx.NewEasingRegister()

			easings.Add("e1", &ed)
			easings.Add("e2", &es)

			if easings.Get("e1") == nil {
				t.Fatalf("\t%s\tShould have retrieved a provider for e1", tests.Failed)
			}
			t.Logf("\t%s\tShould have retrieved a provider for e1", tests.Success)

			if easings.Get("e2") == nil {
				t.Fatalf("\t%s\tShould have retrieved a provider for e1", tests.Failed)
			}
			t.Logf("\t%s\tShould have retrieved a provider for e1", tests.Success)

		}
	}
}

//==============================================================================

type e1 struct{}

func (e *e1) Ease(c govfx.EaseConfig) float64 {
	return 0
}

type e2 struct{}

func (e *e2) Ease(c govfx.EaseConfig) float64 {
	return 1
}

//==============================================================================
