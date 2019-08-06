package parser

import (
	"engine"
	"regexp"
)

var profileCityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" data-v-5e16505f>([^<]+)</a>`)

func ParseCityList(content []byte) engine.ParseResult {
	//接卸城市列表
	matches := profileCityListRe.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items,"City " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
	}
	return result
}
