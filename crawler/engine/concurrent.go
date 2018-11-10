package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
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

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
