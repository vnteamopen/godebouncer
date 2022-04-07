package godebounce

type Debounce struct {
}

func New() *Debounce {
	return Debounce{}
}

func (d *Debouce) WithTrigger(triggerFunc func()) *Debouce {
	return d
}

func (d *Debouce) Do(action func()) {
}
