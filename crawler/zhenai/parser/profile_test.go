package parser

import (
	"testing"
	"zhenai-crawler/crawler/fetcher"
)

func TestParseProfile(t *testing.T) {
	bytes, e := fetcher.Fetch("http://album.zhenai.com/u/1813331607")
	if e != nil {
		panic(e)
	}
	ParseProfile(bytes)
}
