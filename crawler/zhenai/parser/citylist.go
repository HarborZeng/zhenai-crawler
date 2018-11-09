package parser

import (
	"regexp"
	"zhenai-crawler/crawler/common"
	"zhenai-crawler/crawler/engine"
)

const CityListURL = "http://www.zhenai.com/zhenghun"

const cityListRegExp = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]*>([^<]+)</a>`

var regExpCom = regexp.MustCompile(cityListRegExp)

func ParseCityList(contents []byte) engine.ParseResult {
	matches := regExpCom.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 3
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]), ParserFunc: ParseCity,
		})
		if common.DebugMode {
			limit--
		}
		if limit <= 0 {
			break
		}
	}
	return result
}
