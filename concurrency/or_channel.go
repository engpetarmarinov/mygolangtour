package concurrency

func Or(channels ...<-chan struct{}) <-chan struct{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan struct{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			// Here we recursively create an or-channel from all the channels in our slice after third index,
			// and then select from this. This recurrence relation will destructure the rest of the slice into or-channels
			// to form a tree from which the first signal will return.
			// We also pass in the or-Done channel so that when goroutines up the tree exit, goroutine down the tree also exit.
			case <-Or(append(channels[3:], orDone)...):
			}

		}
	}()

	return orDone
}
