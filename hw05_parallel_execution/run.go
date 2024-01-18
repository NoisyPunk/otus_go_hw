package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrErrorsNoWorkers     = errors.New("no workers for tasks")
)

type Task func() error

func worker(id int, tasks <-chan Task, maxErrorCount int, errorChannel chan<- int, stopChannel <-chan int32) {
	for {
		task, ok := <-tasks
		if !ok {
			return
		}
		fmt.Println("task chanel is open")
		err := task()
		fmt.Printf("worker %v now at work\n", id)
		if err != nil {
			errorChannel <- 1
		}
		select {
		case data := <-stopChannel:
			if data == int32(maxErrorCount) {
				fmt.Printf("to much errors worker %v - stoped", id)
				return
			}
		default:
		}
	}
}

func errorChecker(workers int, maxErrorCount int, errorChannel <-chan int, stopChannel chan<- int32) {
	var errCount int32
	for {
		select {
		case data := <-errorChannel:
			if data == 0 {
				return
			}
			atomic.AddInt32(&errCount, 1)
			if errCount == int32(maxErrorCount) {
				for i := 1; i <= workers; i++ {
					fmt.Printf("max count of errors achieved: %v, send signal to workers\n", maxErrorCount)
					stopChannel <- errCount
				}
				return
			}
		default:
			continue
		}
	}
}

func taskScheduler(tasks []Task, taskChannel chan<- Task) {
	for _, t := range tasks {
		taskChannel <- t
		fmt.Println("task added to channel")
	}
	close(taskChannel)
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, workersCount, maxErrorsCount int) error {
	wg := sync.WaitGroup{}
	if workersCount <= 0 {
		return ErrErrorsNoWorkers
	}
	if maxErrorsCount == 0 {
		return ErrErrorsLimitExceeded
	}

	taskChannel := make(chan Task, len(tasks))
	errorChannel := make(chan int, maxErrorsCount)
	stopChannel := make(chan int32)

	for w := 1; w <= workersCount; w++ {
		wg.Add(1)

		w := w
		go func() {
			defer wg.Done()
			worker(w, taskChannel, maxErrorsCount, errorChannel, stopChannel)
		}()
	}

	go errorChecker(workersCount, maxErrorsCount, errorChannel, stopChannel)

	go taskScheduler(tasks, taskChannel)

	wg.Wait()

	close(stopChannel)
	close(errorChannel)

	_, ok := <-taskChannel
	if ok {
		return ErrErrorsLimitExceeded
	}

	return nil
}
