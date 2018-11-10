package main

import (
	"zhenai-crawler/crawler/engine"
	"zhenai-crawler/crawler/scheduler"
	"zhenai-crawler/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        parser.CityListURL,
		ParserFunc: parser.ParseCityList,
	})
}
