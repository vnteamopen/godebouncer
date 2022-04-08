package godebounce

type Debounce struct {
}

func New() *Debounce {
	return &Debounce{}
}

func (d *Debounce) WithTrigger(triggerFunc func()) *Debounce {
	return d
}

func (d *Debounce) Do(action func()) {
}
