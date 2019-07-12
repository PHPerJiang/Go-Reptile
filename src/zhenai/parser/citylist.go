package parser

import (
	"engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" data-v-5e16505f>([^<]+)</a>`

func ParseCityList(content []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	limit := 1
	for _, m := range matches {
		result.Items = append(result.Items,"City " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
		limit --
		if limit == 0 {
			break
		}
	}
	return result
}
