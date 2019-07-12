package scheduler

import (
	"engine"
)

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

//设置chan
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.WorkerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request){
	go func() {
		s.WorkerChan <- r
	}()
}

