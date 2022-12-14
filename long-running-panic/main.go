package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	value := os.Getenv("STOP_AFTER")
	if value == "" {
		value = "30s"
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		panic(err)
	}

	terminate := time.NewTicker(duration)
	go func() {
		for {
			select {
			case <-terminate.C:
				cancel()
				panic("Stop now")
			}
		}
	}()

	sec := time.Second
	n, _ := strconv.Atoi(os.Getenv("TIME_INTERVAL"))
	if n > 0 {
		sec = time.Second * time.Duration(n)
	}
	log.Printf("with interval of %v\n", sec)

	n, _ = strconv.Atoi(os.Getenv("GOROUTINE"))
	if n < 1 {
		n = 100
	}
	log.Printf("spawning %d goroutines ", n)

	wg := new(sync.WaitGroup)
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func(no int) {
			defer wg.Done()
			var u64 uint64
			log.Printf("Goroutine[%d] started\n", no)

		loop:
			for {
				tick := time.NewTicker(sec)

				select {
				case <-ctx.Done():
					tick.Stop()
					break loop
				case <-tick.C:
					u64++
					log.Printf("Goroutine[%d], loop -> %d\n", no, u64)
					tick.Stop()
				}
			}

			log.Printf("Goroutine[%d] ended\n", no)
		}(i)
	}
	wg.Wait()
	terminate.Stop()

}
