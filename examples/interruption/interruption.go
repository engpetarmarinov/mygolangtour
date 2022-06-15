package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx := context.Background()

	// trap Interrupt and Kill and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	defer func() {
		signal.Stop(c)
		cancel()
	}()

	go func() {
		select {
		case <-c:
			fmt.Println("canceling...")
			cancel()
		case <-ctx.Done():
			fmt.Println("Done.")
		}
	}()

	loopWork(ctx)
}

func loopWork(ctx context.Context) {
	tick := time.Tick(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Exiting loop.")
			return
		case <-tick:
			fmt.Println("working...")
		}
	}
}
