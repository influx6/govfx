package govfx_test

import (
	"testing"

	"github.com/ardanlabs/kit/tests"
	"github.com/influx6/govfx"
)

// TestAnimatorsRegistery validates the operation and write safetiness of
func TestAnimatorRegistery(t *testing.T) {
	tests.ResetLog()
	defer tests.DisplayLog()

	t.Logf("Given the need to use use the Animator registery")
	{
		t.Logf("\tWhen giving a sequence to register")
		{

			govfx.RegisterSequence("wix", width{})

			wix, err := govfx.NewSequence("wix", map[string]interface{}{
				"value": 20,
			})

			if err != nil {
				t.Fatalf("\t%s\tShould have successfully created a new sequence: %s", tests.Failed, err)
			}
			t.Logf("\t%s\tShould have successfully created a new sequence", tests.Success)

			wx, ok := wix.(*width)
			if !ok {
				t.Fatalf("\t%s\tShould have a sequence with its udnerline type as 'width'", tests.Failed)
			}
			t.Logf("\t%s\tShould have a sequence with its udnerline type as 'width'", tests.Success)

			if wx.Value != 20 {
				t.Fatalf("\t%s\tShould expect to find attribute Value equal to 20: %d", tests.Failed, wx.Value)
			}
			t.Logf("\t%s\tShould expect to find attribute Value equal to 20", tests.Success)

			if wx.Init(nil, nil) != nil {
				t.Fatalf("\t%s\tShould expect empty writers for Init()", tests.Failed)
			}
			t.Logf("\t%s\tShould expect empty writers for Init()", tests.Success)

			if wx.Next(nil, nil) != nil {
				t.Fatalf("\t%s\tShould expect empty writers for Next()", tests.Failed)
			}
			t.Logf("\t%s\tShould expect empty writers for Next()", tests.Success)
		}
	}
}

//==============================================================================

type width struct {
	Value int `govfx:"value"`
}

func (w width) Init(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	return nil
}

func (w width) Next(stats govfx.Stats, elems govfx.Elementals) govfx.DeferWriters {
	return nil
}
