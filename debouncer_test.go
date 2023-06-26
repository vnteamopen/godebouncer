package godebouncer_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/vnteamopen/godebouncer"
)

func Example() {
	duration := 5 * time.Second
	debouncer := godebouncer.NewWithOptions(
		godebouncer.WithTimeDuration(duration),
		godebouncer.WithTriggered(func() {
			// Triggered func will be called after 5 seconds from last SendSignal()/Do().
			fmt.Println("Trigger")
		}),
	)
	debouncer.SendSignal()
	time.Sleep(1 * time.Second)
	debouncer.SendSignal()

	// After 5 seconds, the trigger will be called.
	// Previous `SendSignal()` will be ignored to trigger the triggered function.
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
	expectedCounter := 1

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
	expectedCounter := 2

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
	expectedCounter := 2

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
	expectedCounter := 1

	debouncer.SendSignal()
	<-debouncer.Done()

	if *countPtr != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, *countPtr)
	}
}

func TestDebounceUpdateTriggeredFuncBeforeDuration(t *testing.T) {
	countPtr, incrementCount := createIncrementCount(0)
	debouncer := godebouncer.New(200 * time.Millisecond).WithTriggered(incrementCount)
	expectedCounter := 2

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
	expectedCounter := 3

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
	expectedCounter := 0

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
	expectedCounter := 1

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
	expectedCounter := 1

	debouncer.SendSignal()
	time.Sleep(200 * time.Millisecond)

	debouncer.UpdateTimeDuration(600 * time.Millisecond)
	<-debouncer.Done()

	if *countPtr != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, *countPtr)
	}
}

func TestDebounceLeading(t *testing.T) {
	duration := 300 * time.Millisecond
	countPtr, incrementCount := createIncrementCount(0)
	debouncer := godebouncer.NewWithOptions(
		godebouncer.WithTimeDuration(duration),
		godebouncer.WithTriggered(incrementCount),
		godebouncer.WithOptions(godebouncer.Options{Leading: true, Trailing: false}),
	)
	expectedCounter := 2

	debouncer.SendSignal() // Called
	debouncer.SendSignal()
	debouncer.SendSignal()
	<-debouncer.Done()
	time.Sleep(duration)
	debouncer.SendSignal() // Called
	debouncer.SendSignal()
	<-debouncer.Done()

	select {
	case <-debouncer.Done():
		t.Error("Done() must not be hang")
	case <-time.After(time.Second):
		break
	}

	if *countPtr != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, *countPtr)
	}
}

func TestDebounceOverlapped(t *testing.T) {
	duration := 300 * time.Millisecond
	countPtr, incrementCount := createIncrementCount(0)
	debouncer := godebouncer.NewWithOptions(
		godebouncer.WithTimeDuration(duration),
		godebouncer.WithTriggered(incrementCount),
		godebouncer.WithOptions(godebouncer.Options{Leading: true, Trailing: true}),
	)
	expectedCounter := 4

	debouncer.SendSignal() // Called
	debouncer.SendSignal()
	debouncer.SendSignal()
	debouncer.SendSignal()
	debouncer.SendSignal()
	debouncer.SendSignal() // Called
	<-debouncer.Done()
	<-debouncer.Done()
	time.Sleep(duration)
	debouncer.SendSignal() // Called
	debouncer.SendSignal()
	debouncer.SendSignal() // Called
	<-debouncer.Done()
	<-debouncer.Done()

	select {
	case <-debouncer.Done():
		t.Error("Done() must not be hang")
	case <-time.After(time.Second):
		break
	}

	if *countPtr != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, *countPtr)
	}
}

func TestDone(t *testing.T) {
	countPtr, incrementCount := createIncrementCount(0)
	debouncer := godebouncer.New(200 * time.Millisecond).WithTriggered(incrementCount)
	expectedCounter := 2

	debouncer.SendSignal()
	<-debouncer.Done()

	debouncer.SendSignal()
	<-debouncer.Done()

	if *countPtr != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, *countPtr)
	}
}

func TestDoneInGoroutine(t *testing.T) {
	duration := 200 * time.Millisecond
	countPtr, incrementCount := createIncrementCount(0)
	debouncer := godebouncer.New(duration).WithTriggered(incrementCount)
	expectedCounter := 3

	debouncer.SendSignal()
	go func() {
		<-debouncer.Done() // awaits for the second send signal to complete
		*countPtr += 2
	}()

	debouncer.SendSignal() // after 1 milliseconds, unblock done channel in 2 goroutines
	<-debouncer.Done()

	time.Sleep(duration)

	if *countPtr != expectedCounter {
		t.Errorf("Expected count %d, was %d", expectedCounter, *countPtr)
	}
}

func TestDoneHangBeforeSendSignal(t *testing.T) {
	debouncer := godebouncer.New(200 * time.Millisecond).WithTriggered(func() {})
	select {
	case <-debouncer.Done():
		t.Error("Done() must not be hang before calling SendSignal()")
	case <-time.After(time.Second):
	}
}

func TestDoneHangAfterSendSignal(t *testing.T) {
	debouncer := godebouncer.New(200 * time.Millisecond).WithTriggered(func() {})
	debouncer.SendSignal()

	select {
	case <-debouncer.Done():
		break
	case <-time.After(time.Second):
		t.Error("Done() must be hang after calling SendSignal()")
	}
}
