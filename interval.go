package setinterval

import (
	"errors"
)

// Interval handler struct
type Interval struct {
	running bool
}

// Pause pauses execution
func (i *Interval) Pause() (ri *Interval, err error) {
	if !i.running {
		return i, errors.New("Interval already paused")
	}

	i.running = false

	return i, nil
}

// Resume resumes execution of paused intervals
func (i *Interval) Resume() (ri *Interval, err error) {
	if i.running {
		return i, errors.New("Interval already running")
	}

	i.running = true

	return i, nil
}
