# Loop
 Loop provides a game loop for different rendering backends(browers RAF,gl,etc).
 It provides a central repository for a central working game loop regardless of
 specific platform/backend. Loop provides a peculiar design in that it allows
 the behaviour and management of the game loop to be powered by a third party
 implementation that meets the basic requirement.

# Install

  ```bash

    > go get -u github.com/influx6/faux/loop/...

  ```

# API

  New(gear loop.EngineGear) => GameEngine
    Loop exposes a single method that returns a new instance of the loop engine,
    powered by the gear which actually implements the real loop mechanism. This
    allows building different looping mechanism or using appropriate ones as
    regards to their respective platforms.
    The `loop.EngineGear` is a simple function type that requires matching the
    below interface set.

      ```go

        type Looper interface {
        	End()
        }

        type Mux func(float64)

        type EngineGear func(Mux) Looper

      ```


# Usage

  ```go

    import "github.com/influx6/faux/loop/web"
    import "github.com/influx6/faux/loop"

    func main(){

        // Create a new engine.
        gameloop := loop.New(web.Loop)

        // Subscribe a function into the gameloop.
        subscriber := gameloop.Loop(func(delta float64){
          //.......
        })

        // End the function recall.
        subscriber.End()
    }

  ```
