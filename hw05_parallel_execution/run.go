package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

var ErrorCount int32

func worker(id int, tasks <-chan Task, errorChannel chan error, stop chan bool) {
	for {
		task, ok := <-tasks
		if !ok {
			return
		}
		fmt.Printf("Working, %v\n", id)
		err := task()
		if err != nil {
			//errorChannel <- err
			atomic.AddInt32(&ErrorCount, 1)

		}
	}
}

func errorChecker(tasks chan Task, errorChannel chan error, maxErrorCount int, stop chan bool, stop2 chan bool) {
	for {
		//fmt.Println("test34")
		if ErrorCount == int32(maxErrorCount) {
			stop <- true
			return
		}
		select {
		case <-stop2:
			return
		default:

		}
	}
	//var errCount int32
	//for {
	//	select {
	//	case _ = <-errorChannel:
	//		atomic.AddInt32(&errCount, 1)
	//		if errCount == int32(maxErrorCount) {
	//			fmt.Println("toMuch")
	//			stop <- true
	//			//close(errorChannel)
	//			//return
	//		}
	//	case <-stop2:
	//		close(errorChannel)
	//		return
	//	}
	//}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := sync.WaitGroup{}
	workersCount := n

	taskChannel := make(chan Task)
	errorChannel := make(chan error)
	stopChannel := make(chan bool, 1)
	stopChannel2 := make(chan bool, 1)

	for w := 1; w <= workersCount; w++ {
		wg.Add(1)

		w := w
		go func() {
			defer wg.Done()
			worker(w, taskChannel, errorChannel, stopChannel2)
			println("test")
		}()
	}
	go errorChecker(taskChannel, errorChannel, m, stopChannel, stopChannel2)

	for _, t := range tasks {
		taskChannel <- t
		select {
		case <-stopChannel:
			close(taskChannel)
			//stopChannel2 <- true
			return ErrErrorsLimitExceeded
		default:

		}
	}

	close(taskChannel)

	//err := <-stopChannel
	//close(taskChannel)
	wg.Wait()
	//defer func() { stopChannel2 <- true }()
	stopChannel2 <- true

	//close(stopChannel)

	return nil
}
