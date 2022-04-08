package godebounce_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/vnteamopen/godebounce"
)


func Example() {
	debounce := godebounce.New(10 * time.Second).WithTrigger(func() {
		fmt.Println("Trigger") // Trigger func will be called after 10 seconds from last action.
	})
	
	debounce.Do(func() {
		fmt.Println("Action 1") // After 10 seconds, the trigger will be called.
	})
	
	time.Sleep(3 * time.Second)
	
	debounce.Do(func() {
		fmt.Println("Action 2")
		// The schedule trigger func of Action 1 will be cleared.
		// After 10 seconds of action 2, trigger will be called.
	})
}

func TestDebounceDoBeforeExpired(t *testing.T) {
	countPointer, incrementCount := createIncrementCount(0)
	debounce := godebounce.New(3 * time.Millisecond).WithTrigger(incrementCount)

	debounce.Do(func() {
		fmt.Println("Action 1")
	})
	
	time.Sleep(2* time.Millisecond)
	
	debounce.Do(func() {
		fmt.Println("Action 2")
	})

	time.Sleep(4 * time.Millisecond)

	if *countPointer != 1 {
		t.Error("Expected count 1, was ", *countPointer)
	}
}

func TestDebounceDoAfterExpired(t *testing.T) {
	countPointer, incrementCount := createIncrementCount(0)
	debounce := godebounce.New(3 * time.Millisecond).WithTrigger(incrementCount)

	debounce.Do(func() {
		fmt.Println("Action 1")
	})
	
	time.Sleep(4 * time.Millisecond)
	
	debounce.Do(func() {
		fmt.Println("Action 2")
	})

	time.Sleep(4 * time.Millisecond)

	if *countPointer != 2 {
		t.Error("Expected count 2, was ", *countPointer)
	}
}


func TestDebounceMixed(t *testing.T) {
	countPointer, incrementCount := createIncrementCount(0)
	debounce := godebounce.New(3 * time.Millisecond).WithTrigger(incrementCount)

	debounce.Do(func() {
		fmt.Println("Action 1")
	})
	
	debounce.Do(func() {
		fmt.Println("Action 2")
	})

	time.Sleep(4 * time.Millisecond)

	debounce.Do(func() {
		fmt.Println("Action 3")
	})

	time.Sleep(4 * time.Millisecond)

	if *countPointer != 2 {
		t.Error("Expected count 2, was ", *countPointer)
	}
}

func TestDebounceWithoutTriggeredFunc(t *testing.T) {
	debounce := godebounce.New(3 * time.Millisecond)

	debounce.Do(func() {
		fmt.Println("Action 1")
	})
	time.Sleep(4 * time.Millisecond)
	fmt.Println("debounce.Do() finished successfully!")
}

func createIncrementCount(counter int) (*int, func()) {
	return &counter, func() {
		fmt.Println("Triggered")
		counter++
	}
}
