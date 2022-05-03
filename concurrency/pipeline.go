package concurrency

func RepeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)

		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()

	return valueStream
}

func Take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeSteam := make(chan interface{})
	go func() {
		defer close(takeSteam)

		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeSteam <- <-valueStream:

			}
		}
	}()

	return takeSteam
}
