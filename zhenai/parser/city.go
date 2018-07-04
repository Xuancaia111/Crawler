package parser

import (
	"crawler/engine"
	"regexp"
)
var(
	profileRe  = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityRe =regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParserResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		name :=string(m[2])
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: func(bytes []byte) engine.ParserResult {
					return ParseProfile(bytes, name)
				},
			})
	}

	cities:=cityRe.FindAllSubmatch(contents,-1)
	for _,c:=range cities{
		result.Requests=append(result.Requests,
			engine.Request{
				Url:string(c[1]),
				ParserFunc: ParseCity,
			})
	}
	return result
}
