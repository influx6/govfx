# Discontinued and Deprecated

# GoVFX
 GoVFX is a idiomatic web animation library which brings the style of [VelocityJS](https://julian.com/research/velocity/) to Go.

## Install

  ```bash
go get -u github.com/influx6/govfx/...
  ```


## Building Examples
  To build the sample files in the `examples` directory, navigate into the
  directory you wish to test and execute the giving command as below.
  The folders will contain basic html and javascript files that will be
  executed once the html as being opened up in a browser.

  Note: Any sample that deals with the shadow DOM must be opened in Google chrome/chromium, has the shadow DOM API has no full browser support by default

  ```bash
gopherjs build app.go
  ```

## Features

  - Dead simple API.
  - Ensures simple and fast execution of animations without hindering performance
  - Provides extendibility in all parts including easing, and property animators
  - Supports animations with Shadow DOM.
  - Batch rendering optimizations for animations.


## Example
  The way VFX was written makes it easy to build animations quickly with as much
  control as possible, yet with efficient optimization applied in.

```go
package main

import (
	"fmt"
	"time"

	"github.com/influx6/govfx"
	_ "github.com/influx6/govfx/animators"
)

func main() {

	begin := govfx.NewListener(func(dl float64) {
		fmt.Printf("Animation Has Begun at %.4f .\n", dl)
	})

	end := govfx.NewListener(func(dl float64) {
		fmt.Printf("Animation Has Ended at %.4f .\n", dl)
	})

	progress := govfx.NewListener(func(dl float64) {
		fmt.Printf("Animation Is Progressing at %.4f .\n", dl)
	})

	elems := govfx.QuerySelectorAll(".zapps")
	width := govfx.Animate(govfx.Stat{
		Duration: 4 * time.Second,
		Loop:     2,
		Reverse:  true,
		Begin:    begin,
		End:      end,
		Progress: progress,
	}, govfx.Values{
		{"value": 500, "animate": "width", "easing": "ease-in"},
	}, elems)

	width.Start()

}

```
