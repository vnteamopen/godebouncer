package godebouncer

import (
	"sync"
	"sync/atomic"
	"time"
)

type Debouncer struct {
	timeDuration  time.Duration
	timer         *time.Timer
	triggeredFunc func()
	mu            sync.Mutex
	done          atomic.Value
}

// New creates a new instance of debouncer. Each instance of debouncer works independent, concurrency with different wait duration.
func New(duration time.Duration) *Debouncer {
	return &Debouncer{timeDuration: duration, triggeredFunc: func() {}}
}

// WithTriggered attached a triggered function to debouncer instance and return the same instance of debouncer to use.
func (d *Debouncer) WithTriggered(triggeredFunc func()) *Debouncer {
	d.triggeredFunc = triggeredFunc
	return d
}

// SendSignal makes an action that notifies to invoke the triggered function after a wait duration.
func (d *Debouncer) SendSignal() {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.Cancel()
	d.closeDone()
	d.done.Store(make(chan struct{}))
	d.timer = time.AfterFunc(d.timeDuration, func() {
		d.mu.Lock()
		d.triggeredFunc()
		d.mu.Unlock()
		d.closeDone()
		d.done.Store(closedchan)
	})
}

// Do run the signalFunc() and call SendSignal() after all. The signalFunc() and SendSignal() function run sequentially.
func (d *Debouncer) Do(signalFunc func()) {
	signalFunc()
	d.SendSignal()
}

// Cancel the timer from the last function SendSignal(). The scheduled triggered function is cancelled and doesn't invoke.
func (d *Debouncer) Cancel() {
	if d.timer != nil {
		d.timer.Stop()
	}
	d.done.CompareAndSwap(nil, closedchan)
}

// UpdateTriggeredFunc replaces triggered function.
func (d *Debouncer) UpdateTriggeredFunc(newTriggeredFunc func()) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.triggeredFunc = newTriggeredFunc
}

// UpdateTimeDuration replaces the waiting time duration. You need to call a SendSignal() again to trigger a new timer with a new waiting time duration.
func (d *Debouncer) UpdateTimeDuration(newTimeDuration time.Duration) {
	d.timeDuration = newTimeDuration
}

func (d *Debouncer) Done() <-chan struct{} {
	done := d.done.Load()
	if done != nil {
		return done.(chan struct{})
	}
	d.done.CompareAndSwap(nil, make(chan struct{}))
	return d.done.Load().(chan struct{})
}

func (d *Debouncer) closeDone() {
	done, _ := d.done.Load().(chan struct{})
	if done != nil && done != closedchan {
		close(done)
	}
}

// closedchan is a reusable closed channel.
var closedchan = make(chan struct{})

func init() {
	close(closedchan)
}
