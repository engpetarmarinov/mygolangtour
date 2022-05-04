package concurrency

import "context"

func RepeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)

		for {
			select {
			case <-ctx.Done():
				return
			case valueStream <- fn():
			}
		}
	}()

	return valueStream
}

func Take(ctx context.Context, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeSteam := make(chan interface{})
	go func() {
		defer close(takeSteam)

		for i := 0; i < num; i++ {
			select {
			case <-ctx.Done():
				return
			case takeSteam <- <-valueStream:

			}
		}
	}()

	return takeSteam
}
