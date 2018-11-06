package main

import (
	"zhenai-crawler/crawler/engine"
	"zhenai-crawler/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        parser.CityListURL,
		ParserFunc: parser.ParseCityList,
	})
}
