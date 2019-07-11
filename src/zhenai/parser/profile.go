package parser

import (
	"engine"
	"model"
	"regexp"
	"strconv"
)

//编译正则
var ageReg  = regexp.MustCompile(`<td width="180"><span class="grayL">年龄：</span>([\d]+)</td>`)
//const marriageReg = `<td width="180"><span class="grayL">婚况：</span>([^<]+)</td> `

//简介解析器
func ParseProfile(contents []byte)engine.ParseResult  {
	//初始化一个简介结构体
	profile := model.Profile{}
	//将年纪解析为一个数字
	age, err := strconv.Atoi(extractString(contents,ageReg))
	//若年纪存在则赋值
	if err != nil {
		profile.Age = age
	}
	//声明一个解析返回结构体
	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return result
}

//公共方法=>提取内容中的与正则匹配的信息
func extractString(contents []byte, re *regexp.Regexp) string {
	//匹配信息，只找一条并提取关键词
	matches := re.FindSubmatch(contents)
	//提取到关键词的内容做返回信息
	if len(matches) >= 2 {
		return string(matches[1])
	}else {
		return ""
	}
}
