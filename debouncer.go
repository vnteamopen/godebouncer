package godebouncer

import (
	"sync"
	"time"
)

type Debouncer struct {
	timeDuration           time.Duration
	timer                  *time.Timer
	triggeredFunc          func()
	signalCalledAt         time.Time
	signalCalledInDuration int
	mu                     sync.Mutex
	done                   chan struct{}
	options                Options
}
type Options struct {
	Leading, Trailing bool
}

type DebouncerOptions func(*Debouncer)

type DebouncerType string

const (
	TRAILING   DebouncerType = "trailing"
	LEADING    DebouncerType = "leading"
	OVERLAPPED DebouncerType = "overlapped"
	INACTIVE   DebouncerType = "inactive"
)

// New creates a new instance of debouncer. Each instance of debouncer works independent, concurrency with different wait duration.
func New(duration time.Duration) *Debouncer {
	return &Debouncer{timeDuration: duration, triggeredFunc: func() {}, options: Options{false, true}}
}

// NewWithOptions takes a slice of option as the rest arguments and return a new instance of debouncer
func NewWithOptions(opts ...DebouncerOptions) *Debouncer {
	var (
		defaultDuration      = 1 * time.Minute
		defaultOptions       = Options{false, true}
		defaultTriggeredFunc = func() {}
	)

	d := &Debouncer{
		timeDuration:  defaultDuration,
		triggeredFunc: defaultTriggeredFunc,
		options:       defaultOptions,
		done:          make(chan struct{}),
	}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

// WithOptions sets the options of debouncer instance.
func WithOptions(options Options) DebouncerOptions {
	return func(d *Debouncer) {
		d.options = options
	}
}

// WithTriggered sets the triggered function of debouncer instance.
func WithTriggered(triggeredFunc func()) DebouncerOptions {
	return func(d *Debouncer) {
		d.triggeredFunc = triggeredFunc
	}
}

// WithTimeDuration sets the time duration of debouncer instance.
func WithTimeDuration(timeDuration time.Duration) DebouncerOptions {
	return func(d *Debouncer) {
		d.timeDuration = timeDuration
	}
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
	now := time.Now()

	switch d.getDebounceType() {
	case TRAILING:
		d.Cancel()
		d.timer = d.invokeTriggeredFunc()
	case LEADING:
		if d.signalCalledAt.IsZero() || now.Sub(d.signalCalledAt) > d.timeDuration {
			d.timer = d.invokeTriggeredFunc()
		}
	case OVERLAPPED:
		if d.signalCalledAt.IsZero() || now.Sub(d.signalCalledAt) >= d.timeDuration {
			d.timer = d.invokeTriggeredFunc()
			break
		}
		if d.signalCalledInDuration > 0 {
			if d.signalCalledInDuration > 1 {
				d.Cancel()
			}
			d.timer = d.invokeTriggeredFunc()
		}
		d.signalCalledInDuration += 1
	default:
	}

}

func (d *Debouncer) invokeTriggeredFunc() *time.Timer {
	d.signalCalledAt = time.Now()

	return time.AfterFunc(d.timeDuration, func() {
		d.triggeredFunc()
		if d.done != nil {
			close(d.done)
		}
		d.done = make(chan struct{})
		d.signalCalledInDuration = 0
	})
}

// getDebounceType get the debouncer time based on Debouncer's options
func (d *Debouncer) getDebounceType() DebouncerType {
	if !d.options.Leading && d.options.Trailing {
		return TRAILING
	}

	if d.options.Leading && !d.options.Trailing {
		return LEADING
	}

	if d.options.Leading && d.options.Trailing {
		return OVERLAPPED
	}

	return INACTIVE
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
}

// UpdateTriggeredFunc replaces triggered function.
func (d *Debouncer) UpdateTriggeredFunc(newTriggeredFunc func()) {
	d.triggeredFunc = newTriggeredFunc
}

// UpdateTimeDuration replaces the waiting time duration. You need to call a SendSignal() again to trigger a new timer with a new waiting time duration.
func (d *Debouncer) UpdateTimeDuration(newTimeDuration time.Duration) {
	d.timeDuration = newTimeDuration
}

// Done returns a receive-only channel to notify the caller when the triggered func has been executed.
func (d *Debouncer) Done() <-chan struct{} {
	if d.done == nil {
		d.done = make(chan struct{})
	}
	return d.done
}
