package godebouncer

import (
	"time"
)

type Debouncer struct {
	timeDuration  time.Duration
	timer         *time.Timer
	triggeredFunc func()
}

func New(duration time.Duration) *Debouncer {
	return &Debouncer{timeDuration: duration, triggeredFunc: func() {}}
}

func (d *Debouncer) WithTriggered(triggeredFunc func()) *Debouncer {
	d.triggeredFunc = triggeredFunc
	return d
}

func (d *Debouncer) Do(signalFunc func()) {
	signalFunc()

	if d.timer != nil {
		d.timer.Stop()
	}

	d.timer = time.AfterFunc(d.timeDuration, d.triggeredFunc)
}
