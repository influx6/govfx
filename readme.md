# GoVFX
 GoVFX is a idiomatic web animation library which brings the style of [VelocityJS](https://julian.com/research/velocity/) to Go.

## Install

  ```go

    go get -u github.com/influx6/govfx/...

  ```

## Features

  - Dead simple API.
  - Ensures simple and fast execution of animations without hindering performance
  - Provides extendibility in all parts including easing, and property animators
  - Supports animations with Shadow DOM.
  - Batch rendering optimizations for animations.

## Concept

  - Sequence and Writers

  `Sequence` in VFX define a behaviour that changes respectively on every tick of
  the animation clock, they hold the calculations that are necessary to achieve
  the desired behaviour. In a way, they define a means of providing a
  processing of the deferred expected behaviour.

  A `sequence` can be a width or height transition, or an opacity animator that
  produces for every iteration the desired end result.
  `Sequences` can be of any type as defined by the animation creator, which
  provides a powerful but flexible system because multiple sequences can be bundled
  together to produce new ones.

  `Sequences` return `Writers`, which are the calculated end result for a animation step.
  The reason there exists such a concept, is due to the need in reducing the effects of
  massive layout trashing, which is basically the browser re-rendering of the DOM
  due to massive multiple changes of properties of different elements, which create
  high costs in performance.

  Writers are returned from sequences to allow effective batching of these changes
  which reduces and minimizes the update calculation performed by the browser DOM.

  - Animation Frames

  `Animation frames` in VFX are a collection of sequences which according to a
  supplied stat will produce the total necessary sequence writers need to
  achieve the desired animation within a specific frame or time of the animation
  loop. It is the central organizational structure in VFX.

  - Stats

  Stats in VFX are captions of current measurements of the animation loop and the
  properties for `Animation Frames`, using stats VFX calls all sequence to produce
  their writers by using the properties of the stats to produce the necessary change
  and easing behaviours that is desired to be achieved.

## Example
  The way VFX was written makes it easy to build animations quickly with as much
  control as possible, yet with efficient optimization applied in.

  - Animating Width Property

      ```go

      package main

      import (
      	"fmt"
      	"time"

      	"github.com/influx6/govfx"
      	"github.com/influx6/govfx/animators"
      )

      func main() {

      	width := govfx.QuerySequence(".zapps",
      		govfx.NewStat(govfx.StatConfig{
      			Duration: 1 * time.Second,
      			Delay:    2 * time.Second,
      			Easing:   "ease-in",
      			Loop:     4,
      			Reverse:  true,
      			Optimize: true,
      		}),
      		&animators.Width{Value: 500})

      	width.OnBegin(func(stats govfx.Frame) {
      		fmt.Println("Animation Has Begun.")
      	})

      	width.OnEnd(func(stats govfx.Frame) {
      		fmt.Println("Animation Has Ended.")
      	})

      	width.OnProgress(func(stats govfx.Frame) {
      		fmt.Println("Animation is progressing.")
      	})

      	govfx.Animate(width)
      }

      ```

      ```html
        <!DOCTYPE html>
        <html>
          <head>
            <meta charset="utf-8">
            <title>VFX Size Animation</title>
            <style>

              .zapps{
                height: 3px;
                background: red;
                margin-bottom: 10px;
              }

              #zapp1{
                width: 10px;
              }

              #zapp2{
                width: 25px;
              }

              #zapp3{
                width: 5px;
              }

              #zapp4{
                width: 30px;
              }

            </style>
          </head>
          <body>
            <div id="zapp1" class="zapps"></div>
            <div id="zapp2" class="zapps"></div>
            <div id="zapp3" class="zapps"></div>
            <div id="zapp4" class="zapps"></div>
          </body>
          <script type="text/javascript" src="app.js"></script>
        </html>

      ```

  - Animating Width Property With ShadowDOM

      ```go

          package main

          import (
          	"fmt"
          	"time"

          	"github.com/influx6/govfx"
          	"github.com/influx6/govfx/animators"
          )

          func main() {

          	root := govfx.NewShadowRoot(govfx.QuerySelector(".root-shadow"))
          	elems := root.QuerySelectorAll(".zapps")

          	width := govfx.DOMSequence(elems,
          		govfx.NewStat(govfx.StatConfig{
          			Duration: 1 * time.Second,
          			Delay:    2 * time.Second,
          			Easing:   "ease-in",
          			Loop:     4,
          			Reverse:  true,
          			Optimize: true,
          		}),
          		&animators.Width{Value: 500})

          	width.OnBegin(func(stats govfx.Frame) {
          		fmt.Println("Animation Has Begun.")
          	})

          	width.OnEnd(func(stats govfx.Frame) {
          		fmt.Println("Animation Has Ended.")
          	})

          	width.OnProgress(func(stats govfx.Frame) {
          		fmt.Println("Animation is progressing.")
          	})

          	govfx.Animate(width)
          }

      ```

      ```html
          <!DOCTYPE html>
          <html>
            <head>
              <meta charset="utf-8">
              <title>VFX Size Animation</title>
            </head>
            <body>
              <div class="root-shadow"></div>
            </body>
            <script>
              root = document.querySelector(".root-shadow")
              shadowRoot = root.createShadowRoot()
              shadowRoot.innerHTML = '<style>.zapps{ height: 3px; background: red; margin-bottom: 10px; } #zapp1{ width: 10px; } #zapp2{ width: 25px; } #zapp3{ width: 5px; } #zapp4{ width: 30px; } </style><div id="zapp1" class="zapps"></div><div id="zapp2" class="zapps"></div><div id="zapp3" class="zapps"></div><div id="zapp4" class="zapps"></div>'
            </script>
            <script type="text/javascript" src="app.js"></script>
          </html>

      ```
