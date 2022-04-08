package godebounce

import (
	"time"
)

type Debounce struct {
	timeDuration  time.Duration
	timer         *time.Timer
	triggeredFunc func()
}

func New(duration time.Duration) *Debounce {
	return &Debounce{timeDuration: duration, triggeredFunc: func() {}}
}

func (d *Debounce) WithTriggered(triggeredFunc func()) *Debounce {
	d.triggeredFunc = triggeredFunc
	return d
}

func (d *Debounce) Do(signalFunc func()) {
	signalFunc()

	if d.timer != nil {
		d.timer.Stop()
	}

	d.timer = time.AfterFunc(d.timeDuration, d.triggeredFunc)
}
