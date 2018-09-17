package parser

import (
	"crawler/engine"
	"regexp"
	"crawler/crawler_contributed/config"
)
var(
	profileRe  = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[\d]+)"[^>]*>([^<]+)</a>`)
	cityRe =regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				Parser: NewProfileParser(string(m[2])),
			})
	}

	cities:=cityRe.FindAllSubmatch(contents,-1)
	for _, m :=range cities{
		result.Requests=append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				Parser: engine.NewFuncParser(ParseCity,config.ParseCity),
			})
	}
	return result
}
