package godebounce

import (
	"time"
)

type Debounce struct {
	timeDuration time.Duration
	timer *time.Timer
	trigger func()
}

func New(duration time.Duration) *Debounce {
	return &Debounce{
		timeDuration: duration,
		trigger: func() {},
	}
}

func (d *Debounce) WithTrigger(triggerFunc func()) *Debounce {
	d.trigger = triggerFunc
	return d
}

func (d *Debounce) Do(action func()) {
	action()

	if d.timer != nil {
		d.timer.Stop()
	}
	
	d.timer = time.AfterFunc(d.timeDuration, d.trigger)
}
