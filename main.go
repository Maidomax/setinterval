/*
Package setinterval mimics the setInterval() function form JavaScript.
It will call a provided function repeadedly, after a time period provided.
It has a blocking variant, and a non-blocking variant that returns a struct that lets you pause/resume the loop
*/
package setinterval

import (
	"time"
)

// Sync runs a passed function periodically
func Sync(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}

// Async runs a passed function periodically returning an Interval struct that can be used to pause it
func Async(d time.Duration, f func()) (ri *Interval) {
	ri = &Interval{running: true}

	go func() {
		for range time.Tick(d) {
			if ri.running {
				f()
			}
		}
	}()

	return ri
}
