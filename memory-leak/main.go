package main

import (
	"context"
	"log"
	"sync/atomic"
	"time"
)

var (
	num atomic.Uint64
)

func ticker(ctx context.Context) {
loop:
	for {
		tick := time.NewTicker(time.Second)
		select {
		case <-ctx.Done():
			break loop
		case <-tick.C:
			go ticker(ctx)
			num.Store(num.Load() + 1)
			log.Println("Goroutine ->", num.Load())
		}
	}
}

func main() {
	stopCh := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create go-routine recursively
	go ticker(ctx)
	num.Store(1)

	<-stopCh
}
