package main

import (
	"fmt"
	"time"

	"github.com/vnteamopen/godebouncer"
)

func main() {
	debouncer := godebouncer.New(5 * time.Second).WithTriggered(func() {
		fmt.Println("Trigger") // Triggered func will be called after 5 seconds from last SendSignal().
	})

	fmt.Println("Action 1")
	debouncer.SendSignal()

	time.Sleep(1 * time.Second)

	fmt.Println("Action 2")
	debouncer.SendSignal()

	// After 5 seconds, the trigger will be called.
	// Previous `SendSignal()` will be ignored to trigger the triggered function.
	<-debouncer.Done()
}
