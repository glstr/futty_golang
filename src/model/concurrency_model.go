package model

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

//SelectFor provides example of usage for select for pattern
func SelectFor(done <-chan interface{}) {
	for {
		select {
		case <-done:
			log.Println("done")
		default:
			//Do non-preemptable work
		}
	}
}

//DoneWork wrap a task and return a channel to tell
//the caller task has done
func DoneWork(task func()) <-chan struct{} {
	complete := make(chan struct{})
	go func() {
		defer close(complete)
		task()
	}()
	return complete
}

//TellChildrenDone show how parent tells children goroutine to
//cancel their work
func TellChildrenDone(taskNum int32) int32 {
	done := make(chan struct{})
	var count int32
	childrenTask := func(done <-chan struct{}) {
		for {
			select {
			case <-done:
				atomic.AddInt32(&count, 1)
				return
			default:
				//do task()
				time.Sleep(1 * time.Second)
			}
		}
	}
	for i := int32(0); i < taskNum; i++ {
		go childrenTask(done)
	}
	time.Sleep(1 * time.Second)
	close(done)
	time.Sleep(5 * time.Second)
	return count
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

func ResultChannel(done chan struct{}, num int32) <-chan interface{} {
	result := make(chan interface{})
	go func() {
		defer close(result)
		for i := int32(0); i < 100; i++ {
			select {
			case <-done:
				return
			default:
				result <- i * i
			}
		}
	}()
	return result
}

//Generator return a num stream
func Generator(done <-chan struct{}, intergers ...int) <-chan int {
	inStream := make(chan int)
	go func() {
		defer close(inStream)
		for val := range intergers {
			select {
			case <-done:
				return
			case inStream <- val:
			}
		}
	}()
	return inStream
}

func AddPipe(done <-chan struct{}, inStream <-chan int, addValue int) <-chan int {
	outStream := make(chan int)
	go func() {
		defer close(outStream)
		for inValue := range inStream {
			select {
			case <-done:
				return
			case outStream <- inValue + addValue:
			}
		}
	}()
	return outStream
}

func MultiplyPipe(done <-chan struct{}, inStream <-chan int,
	multiplyVal int) <-chan int {
	outStream := make(chan int)
	go func() {
		defer close(outStream)
		for inValue := range inStream {
			select {
			case <-done:
				return
			case outStream <- inValue * multiplyVal:
			}
		}
	}()
	return outStream
}

//repeat provides repeat generator
func Repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

//take provides take generator
func Take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

func RepeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	outStream := make(chan interface{})
	go func() {
		defer close(outStream)
		for {
			select {
			case <-done:
				return
			case outStream <- fn():
			}
		}
	}()
	return outStream
}

//FanOutIn provides example of usage of fan out and fan in
func FanIn(done <-chan interface{},
	channels ...chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedStream := make(chan interface{})
	multiplex := func(inStream <-chan interface{}) {
		for {
			defer wg.Done()
			select {
			case <-done:
				return
			case multiplexedStream <- <-inStream:
			}
		}
	}

	wg.Add(len(channels))
	for _, channel := range channels {
		go multiplex(channel)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()
	return multiplexedStream
}

//OrDoneChannel provides example of usage of orDoneChannel
func OrDoneChannel(done, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

//Tee provides ability just like tool tee in unix
func Tee(done <-chan interface{}, in <-chan interface{}) (<-chan interface{}, <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})
	go func() {
		defer close(out1)
		defer close(out2)
		for val := range OrDoneChannel(done, in) {
			var out1, out2 = out1, out2
			for i := 0; i < 2; i++ {
				select {
				case <-done:
				case out1 <- val:
					out1 = nil
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()
	return out1, out2
}

//Bridge destructure the channel of channels into a simple channel
func Bridge(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			var stream <-chan interface{}
			select {
			case maybeStream, ok := <-chanStream:
				if ok == false {
					return
				}
				stream = maybeStream
			case <-done:
				return
			}
			for val := range OrDoneChannel(done, stream) {
				select {
				case valStream <- val:
				case <-done:
				}
			}
		}
	}()
	return valStream
}
