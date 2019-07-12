package engine

import (
	"fetcher"
	"log"
)

func Run(seeds ...Request)  {
	var requests []Request
	for _,r := range seeds{
		requests = append(requests, r)
	}

	//如果请求列表不为空则继续抓取
	for len(requests) > 0{
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching: %s\n", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err!=nil {
			log.Printf("Fetcher: error fetching url %s: %v",r.Url, err)
			continue
		}
		//调用解析器
		parseRequest := r.ParseFunc(body)
		requests = append(requests, parseRequest.Requests...)

		for _,item := range parseRequest.Items{
			log.Printf("Got item %v", item)
		}
	}
}
