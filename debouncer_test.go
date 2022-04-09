package godebouncer_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/vnteamopen/godebouncer"
)

func Example() {
	debouncer := godebouncer.New(10 * time.Second).WithTriggered(func() {
		fmt.Println("Trigger") // Triggered func will be called after 10 seconds from last action.
	})

	debouncer.Do(func() {
		fmt.Println("Action 1") // After 10 seconds, the trigger will be called.
	})

	time.Sleep(3 * time.Second)

	debouncer.Do(func() {
		fmt.Println("Action 2")
		// The scheduler of triggered func of Action 1 will be cleared.
		// After 10 seconds of action 2, triggered will be called.
	})
}

func TestDebounceDoBeforeExpired(t *testing.T) {
	countPtr, incrementCount := createIncrementCount(0)
	debouncer := godebouncer.New(200 * time.Millisecond).WithTriggered(incrementCount)
	expectedCounter := int(1)

	debouncer.Do(func() {
		fmt.Println("Action 1")
	})

	time.Sleep(100 * time.Millisecond)

	debouncer.Do(func() {
		fmt.Println("Action 2")
	})

	time.Sleep(400 * time.Millisecond)

	if *countPtr != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, *countPtr)
	}
}

func TestDebounceDoAfterExpired(t *testing.T) {
	countPtr, incrementCount := createIncrementCount(0)
	debouncer := godebouncer.New(200 * time.Millisecond).WithTriggered(incrementCount)
	expectedCounter := int(2)

	debouncer.Do(func() {
		fmt.Println("Action 1")
	})

	time.Sleep(400 * time.Millisecond)

	debouncer.Do(func() {
		fmt.Println("Action 2")
	})

	time.Sleep(400 * time.Millisecond)

	if *countPtr != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, *countPtr)
	}
}

func TestDebounceMixed(t *testing.T) {
	countPtr, incrementCount := createIncrementCount(0)
	debouncer := godebouncer.New(200 * time.Millisecond).WithTriggered(incrementCount)
	expectedCounter := int(2)

	debouncer.Do(func() {
		fmt.Println("Action 1")
	})

	debouncer.Do(func() {
		fmt.Println("Action 2")
	})

	time.Sleep(400 * time.Millisecond)

	debouncer.Do(func() {
		fmt.Println("Action 3")
	})

	time.Sleep(400 * time.Millisecond)

	if *countPtr != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, *countPtr)
	}
}

func TestDebounceWithoutTriggeredFunc(t *testing.T) {
	debouncer := godebouncer.New(200 * time.Millisecond)

	debouncer.Do(func() {
		fmt.Println("Action 1")
	})
	time.Sleep(400 * time.Millisecond)
	fmt.Println("debouncer.Do() finished successfully!")
}

func TestDebounceSendSignal(t *testing.T) {
	countPtr, incrementCount := createIncrementCount(0)
	debouncer := godebouncer.New(200 * time.Millisecond).WithTriggered(incrementCount)
	expectedCounter := int(1)

	debouncer.SendSignal()
	time.Sleep(400 * time.Millisecond)

	if *countPtr != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, *countPtr)
	}
}

func createIncrementCount(counter int) (*int, func()) {
	return &counter, func() {
		fmt.Println("Triggered")
		counter++
	}
}
