package godebouncer

import (
	"sync"
	"time"
)

type Debouncer struct {
	timeDuration  time.Duration
	timer         *time.Timer
	triggeredFunc func()
	mu            sync.Mutex
}

// New creates a new instance of debouncer. Each instance of debouncer works independent, concurrency with different wait duration.
func New(duration time.Duration) *Debouncer {
	return &Debouncer{timeDuration: duration, triggeredFunc: func() {}}
}

// WithTriggered attached a triggered function to debouncer instance and return same instance of debouncer to use
func (d *Debouncer) WithTriggered(triggeredFunc func()) *Debouncer {
	d.triggeredFunc = triggeredFunc
	return d
}

// SendSignal makes an action that notify to invoke triggered function after wait duration.
func (d *Debouncer) SendSignal() {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.Cancel()
	d.timer = time.AfterFunc(d.timeDuration, func() {
		d.triggeredFunc()
	})
}

// Do allows user run a entered a function and call SendSignal after all. Entered function and SendSignal function run sequencely.
func (d *Debouncer) Do(signalFunc func()) {
	signalFunc()
	d.SendSignal()
}

// Cancel the timmer from the last function SendSignal(). The scheduled triggered function is cancelled and doesn't invoke.
func (d *Debouncer) Cancel() {
	if d.timer != nil {
		d.timer.Stop()
	}
}

// UpdateTriggeredFunc replaces triggered function
func (d *Debouncer) UpdateTriggeredFunc(newTriggeredFunc func()) {
	d.triggeredFunc = newTriggeredFunc
}

// UpdateTimeDuratioe replaces the waiting time duration. You need to call a SendSignal() again to trigger a new timmer with a new waiting time duration.
func (d *Debouncer) UpdateTimeDuration(newTimeDuration time.Duration) {
	d.timeDuration = newTimeDuration
}
