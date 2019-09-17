package mtrf

import (
	"time"
)

type Event struct {
	ch chan struct{}
}

func NewEvent() *Event {
	return &Event{ch: make(chan struct{}, 1)}
}

func (ev *Event) Clear() {
	for {
		select {
		case <-ev.ch:
			// pass
		default:
			return
		}
	}
}

func (ev *Event) Raise() {
	select {
	case ev.ch <- struct{}{}:
		// pass
	default:
		// pass
	}
}

func (ev *Event) Wait(timeout time.Duration) bool {
	select {
	case <-time.After(timeout):
		return false
	case <-ev.ch:
		return true
	}
}
