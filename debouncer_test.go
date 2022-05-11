package godebouncer_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/vnteamopen/godebouncer"
)

func Example() {
	wait := 5 * time.Second
	debouncer := godebouncer.New(wait).WithTriggered(func() {
		fmt.Println("Trigger") // Triggered func will be called after 5 seconds from last SendSignal().
	})

	fmt.Println("Action 1")
	debouncer.SendSignal()

	time.Sleep(1 * time.Second)

	fmt.Println("Action 2")
	debouncer.SendSignal()

	// After 5 seconds, the trigger will be called.
	//Previous `SendSignal()` will be ignore to trigger the triggered function.
	<-debouncer.Done()
}

func createIncrementCount(counter int) (*int, func()) {
	return &counter, func() {
		fmt.Println("Triggered")
		counter++
	}
}

func TestDebounceDoBeforeExpired(t *testing.T) {
	countPtr, incrementCount := createIncrementCount(0)
	debouncer := godebouncer.New(200 * time.Millisecond).WithTriggered(incrementCount)
	expectedCounter := int(1)

	debouncer.Do(func() {
		fmt.Println("Action 1")
	})

	time.Sleep(50 * time.Millisecond)

	debouncer.Do(func() {
		fmt.Println("Action 2")
	})

	<-debouncer.Done()

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

	<-debouncer.Done()

	debouncer.Do(func() {
		fmt.Println("Action 2")
	})

	<-debouncer.Done()

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

	<-debouncer.Done()

	debouncer.Do(func() {
		fmt.Println("Action 3")
	})

	<-debouncer.Done()

	if *countPtr != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, *countPtr)
	}
}

func TestDebounceWithoutTriggeredFunc(t *testing.T) {
	debouncer := godebouncer.New(200 * time.Millisecond)

	debouncer.Do(func() {
		fmt.Println("Action 1")
	})
	<-debouncer.Done()

	fmt.Println("debouncer.Do() finished successfully!")
}

func TestDebounceSendSignal(t *testing.T) {
	countPtr, incrementCount := createIncrementCount(0)
	debouncer := godebouncer.New(200 * time.Millisecond).WithTriggered(incrementCount)
	expectedCounter := int(1)

	debouncer.SendSignal()
	<-debouncer.Done()

	if *countPtr != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, *countPtr)
	}
}

func TestDebounceUpdateTriggeredFuncBeforeDuration(t *testing.T) {
	countPtr, incrementCount := createIncrementCount(0)
	debouncer := godebouncer.New(200 * time.Millisecond).WithTriggered(incrementCount)
	expectedCounter := int(2)

	debouncer.SendSignal()
	time.Sleep(50 * time.Millisecond)

	debouncer.UpdateTriggeredFunc(func() {
		*countPtr += 2
	})
	<-debouncer.Done()

	if *countPtr != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, *countPtr)
	}
}

func TestDebounceUpdateTriggeredFuncAfterDuration(t *testing.T) {
	countPtr, incrementCount := createIncrementCount(0)
	debouncer := godebouncer.New(200 * time.Millisecond).WithTriggered(incrementCount)
	expectedCounter := int(3)

	debouncer.SendSignal()
	<-debouncer.Done()

	debouncer.UpdateTriggeredFunc(func() {
		*countPtr += 2
	})
	debouncer.SendSignal()
	<-debouncer.Done()

	if *countPtr != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, *countPtr)
	}
}

func TestDebounceCancel(t *testing.T) {
	countPtr, incrementCount := createIncrementCount(0)
	debouncer := godebouncer.New(200 * time.Millisecond).WithTriggered(incrementCount)
	expectedCounter := int(0)

	debouncer.SendSignal()
	time.Sleep(50 * time.Millisecond)

	debouncer.Cancel()
	time.Sleep(400 * time.Millisecond)

	if *countPtr != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, *countPtr)
	}
}

func TestDebounceUpdateDuration(t *testing.T) {
	countPtr, incrementCount := createIncrementCount(0)
	debouncer := godebouncer.New(600 * time.Millisecond).WithTriggered(incrementCount)
	expectedCounter := int(1)

	debouncer.UpdateTimeDuration(200 * time.Millisecond)
	debouncer.SendSignal()
	<-debouncer.Done()

	if *countPtr != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, *countPtr)
	}
}

func TestDebounceUpdateDurationAfterSendSignal(t *testing.T) {
	countPtr, incrementCount := createIncrementCount(0)
	debouncer := godebouncer.New(400 * time.Millisecond).WithTriggered(incrementCount)
	expectedCounter := int(1)

	debouncer.SendSignal()
	time.Sleep(200 * time.Millisecond)

	debouncer.UpdateTimeDuration(600 * time.Millisecond)
	<-debouncer.Done()

	if *countPtr != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, *countPtr)
	}
}

func TestDone(t *testing.T) {
	debouncer := godebouncer.New(200 * time.Millisecond).WithTriggered(func() {
		fmt.Println("Triggered")
	})

	debouncer.SendSignal()
	<-debouncer.Done()

	fmt.Println("debouncer has done signalling")
}
