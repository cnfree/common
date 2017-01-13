package common

import (
	"time"
	"errors"
	"sync"
	"reflect"
)

type taskUtil struct {
	mutex sync.Mutex
}

var Task = taskUtil{}

var errTimeout = errors.New("timeout")

func (this taskUtil) ExecuteWithTimeout(f func() error, timeout time.Duration) (err error) {
	done := make(chan bool)
	go func() {
		err = f()
		done <- true
	}()
	select {
	case <-time.After(timeout):
		return errTimeout
	case <-done:
		return
	}
}

func (this taskUtil) MergeChannel(cs []chan reflect.Value) (out chan reflect.Value) {
	out = make(chan reflect.Value)
	this.MergeChannelTo(cs, nil, out)
	return out
}

func (this taskUtil) MergeChannelTo(cs []chan reflect.Value, transformFn func(reflect.Value) reflect.Value, out chan reflect.Value) {
	var wg sync.WaitGroup

	for _, c := range cs {
		wg.Add(1)
		go func(c chan reflect.Value) {
			defer wg.Done()
			for n := range c {
				if transformFn != nil {
					n = transformFn(n)
				}
				out <- n
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return
}
