package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	counter := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Fetched item #%d: %v", counter, item)
			counter++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(out chan ParseResult, scheduler Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			scheduler.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
