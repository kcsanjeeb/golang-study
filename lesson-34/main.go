package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func SimulateLongRunningTask(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context cancelled, task stopping ...")
			return
		default:
			fmt.Println("Simulating work ...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // ensure cancellation if we exit early

	var wg sync.WaitGroup
	wg.Add(1)
	go SimulateLongRunningTask(ctx, &wg)

	// simulate some main thread work for 2 seconds
	time.Sleep(2 * time.Second)

	// cancel the context after some time
	fmt.Println("Cancelling context")
	cancel()

	// wait for the goroutine to finish
	wg.Wait()
}
