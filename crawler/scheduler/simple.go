package scheduler

import "zhenai-crawler/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(re engine.Request) {
	go func() {
		s.workerChan <- re
	}()
}

func (s *SimpleScheduler) ConfigureMasterWorkChan(c chan engine.Request) {
	s.workerChan = c
}
