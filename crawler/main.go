package main

import (
	"zhenai-crawler/crawler/engine"
	"zhenai-crawler/crawler/persist"
	"zhenai-crawler/crawler/scheduler"
	"zhenai-crawler/crawler/zhenai/parser"
)

func main() {
	saver, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    saver,
	}
	e.Run(engine.Request{
		Url:        parser.CityListURL,
		ParserFunc: parser.ParseCityList,
	})
}
