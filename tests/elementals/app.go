package main

import (
	"fmt"
	"time"

	"github.com/influx6/faux/vfx"
)

func main() {

	div := vfx.Document().QuerySelector(".expandable")
	expand := vfx.NewElement(div, "")

	report("margin-top", one(expand.Read("margin-top")))
	report("margin-left", one(expand.Read("margin-left")))

	expand.Write("margin-left", "10px", false)
	expand.Write("margin-top", "40px", true)

	<-time.After(500 * time.Millisecond)

	expand.Write("background-color", "red", true)

	report("margin-top", one(expand.Read("margin-top")))
	report("margin-left", one(expand.Read("margin-left")))

	<-time.After(500 * time.Millisecond)
	expand.Sync()

	<-time.After(500 * time.Millisecond)
	expand.Write("padding-left", "40px", true)
	expand.Write("padding-bottom", "40px", true)
	expand.Write("padding-right", "40px", true)
	expand.Write("padding-top", "40px", true)

	expand.Sync()

	report("margin-top", one(expand.Read("margin-top")))
	report("margin-left", one(expand.Read("margin-left")))

}

// one returns the first argument.
func one(b ...interface{}) interface{} {
	if len(b) == 0 {
		return nil
	}

	return b[0]
}

func report(key string, val interface{}) {
	fmt.Printf("Report[%s]: %+s\n", key, val)
}
