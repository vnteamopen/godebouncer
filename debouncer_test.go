package godebouncer_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/vnteamopen/godebouncer"
)

func Example() {
	debouce := godebouncer.New(10 * time.Second).WithTrigger(func() {
		fmt.Println("Trigger") // Trigger func will be called after 10 seconds from last action.
	})

	debouce.Do(func() {
		fmt.Println("Action 1") // After 10 seconds, the trigger will be called.
	})

	time.Sleep(3 * time.Second)

	debouce.Do(func() {
		fmt.Println("Action 2")
		// The schedule trigger func of Action 1 will be cleared.
		// After 10 seconds of action 2, trigger will be called.
	})
}

var counter uint64
var trigger = func() {
	fmt.Println("Trigger")
	counter++
}
var resetCounter = func() {
	counter = 0
}

func TestDebounceDoBeforeExpired(t *testing.T) {
	resetCounter()
	debouce := godebouncer.New(200 * time.Millisecond).WithTrigger(trigger)
	expectedCounter := uint64(1)

	debouce.Do(func() {
		fmt.Println("Action 1")
	})

	time.Sleep(100 * time.Millisecond)

	debouce.Do(func() {
		fmt.Println("Action 2")
	})

	time.Sleep(300 * time.Millisecond)

	if counter != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, counter)
	}
}

func TestDebounceDoAfterExpired(t *testing.T) {
	resetCounter()
	debouce := godebouncer.New(200 * time.Millisecond).WithTrigger(trigger)
	expectedCounter := uint64(2)

	debouce.Do(func() {
		fmt.Println("Action 1")
	})

	time.Sleep(300 * time.Millisecond)

	debouce.Do(func() {
		fmt.Println("Action 2")
	})

	time.Sleep(300 * time.Millisecond)

	if counter != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, counter)
	}
}

func TestDeounceMixed(t *testing.T) {
	resetCounter()
	debouce := godebouncer.New(200 * time.Millisecond).WithTrigger(trigger)
	expectedCounter := uint64(2)

	debouce.Do(func() {
		fmt.Println("Action 1")
	})

	debouce.Do(func() {
		fmt.Println("Action 2")
	})

	time.Sleep(300 * time.Millisecond)

	debouce.Do(func() {
		fmt.Println("Action 3")
	})

	time.Sleep(300 * time.Millisecond)

	if counter != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, counter)
	}
}
