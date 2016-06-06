# Fque
  Fque provides a simple no argument pubsub with subscriber type features.


## Usage

  ```go

			q := fque.New()

			sub := q.Q(func() {
        // Will be called on every emission.
			})

			q.Q(func() {
        // Will be called on every emission.
			})

			q.Run()

      sub.End()

  ```
