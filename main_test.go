package main

import (
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	// Test case 1: positive scenario
	future := Run(func() interface{} {
		time.Sleep(2 * time.Second)
		return "Hello World!"
	})
	res, err := future.GetWithTimeout(3 * time.Second)
	if err != nil {
		t.Errorf("Test case 1: expected result, but got an error: %v", err)
	}
	if res != "Hello World!" {
		t.Errorf("Test case 1: expected 'Hello World!', but got %v", res)
	}
	if !future.IsComplete() {
		t.Error("Test case 1: expected future to be completed")
	}
	if future.IsCancelled() {
		t.Error("Test case 1: expected future to not be cancelled")
	}

	// Test case 2: future cancelled
	future = Run(func() interface{} {
		time.Sleep(2 * time.Second)
		return "Hello World!"
	})
	future.Cancel()
	res, err = future.Get()
	if err == nil {
		t.Error("Test case 2: expected an error, but got no error")
	}
	if res != nil {
		t.Errorf("Test case 2: expected nil, but got %v", res)
	}
	if future.IsComplete() {
		t.Error("Test case 2: expected future to not be completed")
	}
	if !future.IsCancelled() {
		t.Error("Test case 2: expected future to be cancelled")
	}

	// Test case 3: timeout scenario
	future = Run(func() interface{} {
		time.Sleep(2 * time.Second)
		return "Hello World!"
	})
	res, err = future.GetWithTimeout(1 * time.Second)
	if err == nil {
		t.Error("Test case 3: expected an error, but got no error")
	}
	if res != nil {
		t.Errorf("Test case 3: expected nil, but got %v", res)
	}
	if future.IsComplete() {
		t.Error("Test case 3: expected future to not be completed")
	}
	if future.IsCancelled() {
		t.Error("Test case 3: expected future to not be cancelled")
	}
}
