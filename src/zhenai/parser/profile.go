package parser

import (
	"engine"
	"log"
	"model"
	"regexp"
	"strconv"
)

//编译正则

var reg  = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]+)</div><div class="m-btn purple" data-v-bff6f798>([\d]+)岁</div><div class="m-btn purple" data-v-bff6f798>([^<]+)</div><div class="m-btn purple" data-v-bff6f798>([\d]+)cm</div><div class="m-btn purple" data-v-bff6f798>([\d]+)kg</div><div class="m-btn purple" data-v-bff6f798>工作地:([^<]+)</div><div class="m-btn purple" data-v-bff6f798>月收入:([^<]+)</div><div class="m-btn purple" data-v-bff6f798>([^<]+)</div><div class="m-btn purple" data-v-bff6f798>([^<]+)</div></div>`)
//const marriageReg = `<td width="180"><span class="grayL">婚况：</span>([^<]+)</td> `

//简介解析器
func ParseProfile(contents []byte, name string)engine.ParseResult  {
	//初始化一个简介结构体
	profile := model.Profile{}
	profile.Name = name
	//婚姻
	marriage := extractString(contents,1, reg)
	profile.Marriage = marriage

	//年纪
	age, err := strconv.Atoi(extractString(contents,2, reg))
	if err != nil {
		log.Fatal(err)
	}
	profile.Age = age

	//星座
	xingzuo := extractString(contents,3, reg)
	profile.Xingzuo = xingzuo

	//身高
	height,err := strconv.Atoi(extractString(contents,4, reg))
	if err != nil {
		log.Fatal(err)
	}
	profile.Height = height

	//体重
	weight,err := strconv.Atoi(extractString(contents,5,reg))
	if err != nil {
		log.Fatal(err)
	}
	profile.Weight = weight

	//工作地
	address := extractString(contents,6,reg)
	profile.Address = address

	//收入
	income := extractString(contents,7,reg)
	profile.Income = income

	//专业
	major := extractString(contents,8,reg)
	profile.Major = major

	//学历
	education := extractString(contents,9,reg)
	profile.Education = education

	//声明一个解析返回结构体
	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return result
}

//公共方法=>提取内容中的与正则匹配的信息
func extractString(contents []byte,position int,re *regexp.Regexp) string {
	//匹配信息，只找一条并提取关键词
	matches := re.FindSubmatch(contents)
	//提取到关键词的内容做返回信息
	if len(matches) >= 2 {
		return string(matches[position])
	}else {
		return ""
	}
}
