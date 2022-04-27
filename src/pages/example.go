package main

import (
	"fmt"
	"time"
	"github.com/vnteamopen/godebouncer"
)

func main() {
	wait := 10 * time.Second
	debouncer := godebouncer.New(wait).WithTriggered(func() {
		fmt.Println("Trigger") // Triggered func will be called after 10 seconds from last SendSignal().
	})

	fmt.Println("Action 1")
	debouncer.SendSignal()

	time.Sleep(3 * time.Second)

	fmt.Println("Action 2")
	debouncer.SendSignal()
	// After 10 seconds, the trigger will be called.
	// Previous `SendSignal()` will be ignored to trigger the triggered function.

	time.Sleep(10 * time.Second)
}
