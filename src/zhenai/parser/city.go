package parser

import (
	"engine"
	"regexp"
)

var (
	profileCityRe = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a></th>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
	)

func ParseCity(contents []byte)engine.ParseResult {
	matches := profileCityRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items,"User " + name)
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}

	//解析城市列表页里的其他页面
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches{
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
	}
	return result
}
