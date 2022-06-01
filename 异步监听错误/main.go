package main

import (
	"errors"
	"sync"
)

func test(i int) (int, error) {
	if i > 2 {
		return 0, errors.New("test error")
	}
	return i + 5, nil
}

func test2(i int) (int, error) {
	if i > 3 {
		return 0, errors.New("test2 error")
	}
	return i + 7, nil
}

func main() {
	results := make(chan int, 2)
	errs := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		result, err := test(3)
		if err != nil {
			errs <- err
			return
		}
		results <- result
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		result, err := test2(3)
		if err != nil {
			errs <- err
			return
		}
		results <- result
	}()

	go func() {
		wg.Wait()
		close(results)
		close(errs)
	}()
	for err := range errs {
		println(err.Error())
		return

	}
	for res := range results {
		println("--------- ", res, " ------------")
	}
}
