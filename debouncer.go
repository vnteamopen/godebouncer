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

func (d *Debouncer) SendSignal() {
	d.Cancel()
	d.timer = time.AfterFunc(d.timeDuration, func() {
		d.triggeredFunc()
	})
}

func (d *Debouncer) Do(signalFunc func()) {
	signalFunc()
	d.SendSignal()
}

func (d *Debouncer) Cancel() {
	if d.timer != nil {
		d.timer.Stop()
	}
}

func (d *Debouncer) UpdateTriggeredFunc(newTriggeredFunc func()) {
	d.triggeredFunc = newTriggeredFunc
}
