package parser

import (
	"regexp"
	"zhenai-crawler/crawler/engine"
)

const userRegExp = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

var userRegExpCom = regexp.MustCompile(userRegExp)

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
	return result
}
