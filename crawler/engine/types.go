package engine

import "zhenai-crawler/crawler/model"

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []model.Profile
}

// do nothing parser
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
