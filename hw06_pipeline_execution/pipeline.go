package hw06pipelineexecution

import (
	"errors"
	"fmt"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

var ErrNilChannel = errors.New("data channel is nil")

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	errorChan := make(Bi)
	if in == nil {
		go func() {
			defer close(errorChan)
			errorChan <- ErrNilChannel
		}()
		return errorChan
	}
	exec := in
	for i, stage := range stages {
		if stage != nil {
			fmt.Printf("Pipeline %v, executing\n", i+1)
			exec = stage(stageExec(done, exec, i+1))
		}
	}
	return exec
}

func stageExec(done In, in In, stageCounter int) Out {
	outcome := make(Bi)
	go func() {
		defer close(outcome)
		for {
			select {
			case data, ok := <-in:
				if !ok {
					return
				}
				fmt.Printf("Stage %v with value %v in progress\n", stageCounter, data)
				outcome <- data
			case <-done:
				return

			}
		}
	}()
	return outcome
}
