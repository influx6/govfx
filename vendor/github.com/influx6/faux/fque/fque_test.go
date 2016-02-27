package fque_test

import (
	"sync"
	"testing"

	"github.com/ardanlabs/kit/tests"
	"github.com/influx6/faux/fque"
)

func init() {
	tests.Init("")
}

// TestQueue validates the behaviour of que api.
func TestQueue(t *testing.T) {
	tests.ResetLog()
	defer tests.DisplayLog()

	t.Logf("Should be able to use a argument selective queue")
	{

		t.Logf("\tWhen giving a mque.Que and a string only allowed constraint")
		{

			var wq sync.WaitGroup
			wq.Add(1)

			q := fque.New()

			q.Q(func() {
				wq.Done()
			})

			q.Run()

			wq.Wait()

			t.Logf("\t%s\tshould have successfully called callback", tests.Success)
		}
	}
}

// TestQueueEnd validates the behaviour of que subscriber End call.
func TestQueueEnd(t *testing.T) {
	tests.ResetLog()
	defer tests.DisplayLog()

	t.Logf("Should be able to use a argument selective queue")
	{

		t.Logf("\tWhen needing to unsubscribe from a queue")
		{

			var count int

			q := fque.New()

			q.Q(func() {
				count++
			})

			sub := q.Q(func() {
				count++
			})

			q.Run()
			sub.End()
			q.Run()

			if count < 2 {
				t.Fatalf("\t%s\tShould have received a string", tests.Failed)
			}
			t.Logf("\t%s\tshould have received a string", tests.Success)

		}
	}
}
