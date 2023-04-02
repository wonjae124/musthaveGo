package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "number", 9)
	go square(ctx)
	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
}

func square(ctx context.Context) {
	for {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			wg.Done()
			return

		default:
			v := ctx.Value("number")
			if v != nil {
				n := v.(int)
				fmt.Printf("Square:%d\n", n*n)
			}

		}
	}
}
