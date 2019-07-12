package parser

import (
	"engine"
	"regexp"
)

const userReg  = `<th><a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a></th>`

func ParseCity(contents []byte)engine.ParseResult {
	re := regexp.MustCompile(userReg)
	matches := re.FindAllSubmatch(contents, -1)

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
	return result
}
