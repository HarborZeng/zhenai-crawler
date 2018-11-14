package parser

import (
	"regexp"
	"zhenai-crawler/crawler/engine"
)

const userRegExp = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
const cityUrlRegExp = `<a href="(http://www.zhenai.com/zhenghun/[^"]+)">下一页</a>`

var userRegExpCom = regexp.MustCompile(userRegExp)
var cityUrlRegExpCom = regexp.MustCompile(cityUrlRegExp)

func ParseCity(contents []byte) engine.ParseResult {
	matches := userRegExpCom.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	for _, match := range matches {
		result.Items = append(result.Items, string(match[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(match[1]),
			ParserFunc: ParseProfile,
		})
	}

	matches = cityUrlRegExpCom.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(match[1]),
				ParserFunc: ParseCity,
			})
	}
	return result
}
