package main

import (
	"context"
	"fmt"
	"time"
)

// Future interface defines methods for asynchronous operation
type Future interface {
	Get() (interface{}, error)
	GetWithTimeout(time.Duration) (interface{}, error)
	IsComplete() bool
	IsCancelled() bool
	Cancel()
}

// futureImpl implements the Future interface
type futureImpl struct {
	result      chan interface{}
	isComplete  bool
	isCancelled bool
	ctx         context.Context
	cancel      context.CancelFunc
}

// Get returns the result of the operation or an error if it was cancelled
func (f *futureImpl) Get() (interface{}, error) {
	select {
	case res := <-f.result:
		return res, nil
	case <-f.ctx.Done():
		return nil, fmt.Errorf("execution cancelled")
	}
}

// GetWithTimeout returns the result of the operation or an error if it was cancelled or if it took too long
func (f *futureImpl) GetWithTimeout(duration time.Duration) (interface{}, error) {
	select {
	case res := <-f.result:
		return res, nil
	case <-time.After(duration):
		return nil, fmt.Errorf("execution timed out")
	case <-f.ctx.Done():
		return nil, fmt.Errorf("execution cancelled")
	}
}

// IsComplete returns whether the operation is complete
func (f *futureImpl) IsComplete() bool {
	return f.isComplete
}

// IsCancelled returns whether the operation was cancelled
func (f *futureImpl) IsCancelled() bool {
	return f.isCancelled
}

// Cancel cancels the operation
func (f *futureImpl) Cancel() {
	f.cancel()
	f.isCancelled = true
}

// Run starts the operation and returns a Future interface
func Run(fn func() interface{}) Future {
	result := make(chan interface{})
	f := &futureImpl{
		result: result,
		ctx:    context.Background(),
	}
	f.ctx, f.cancel = context.WithCancel(f.ctx)
	go func() {
		select {
		case <-f.ctx.Done():
			return
		default:
			f.result <- fn()
			f.isComplete = true
			close(f.result)
		}
	}()
	return f
}

func main() {
	future := Run(func() interface{} {
		time.Sleep(2 * time.Second)
		return "Hello World!"
	})

	//future.Cancel()
	//fmt.Println(future.Get())

	res, err := future.GetWithTimeout(3 * time.Second)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
