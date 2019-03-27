package model

import "log"

//SelectFor provides example of usage for select for pattern
func SelectFor(done <-chan interface{}) {
	for {
		select {
		case <-done:
			log.Println("done")
		}
	}
}

//OrChannel provides example of usage for or channel
func OrChannel(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
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
			case <-OrChannel(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}
