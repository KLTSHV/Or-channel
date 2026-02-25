package or

import (
	"fmt"
	"time"
)

func ExampleOr() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	<-Or(
		sig(2*time.Second),
		sig(1*time.Second),
	)

	fmt.Println("done after", time.Since(start) < 1500*time.Millisecond)
}
