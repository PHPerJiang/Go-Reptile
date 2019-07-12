package engine

import (
	"fetcher"
	"log"
)

type SimpleEngine struct {

}

func (e SimpleEngine)Run(seeds ...Request)  {
	var requests []Request
	for _,r := range seeds{
		requests = append(requests, r)
	}

	//如果请求列表不为空则继续抓取
	for len(requests) > 0{
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err!=nil {
			continue
		}
		//调用解析器
		requests = append(requests, parseResult.Requests...)

		for _,item := range parseResult.Items{
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error){
	log.Printf("Fetching: %s\n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err!=nil {
		log.Printf("Fetcher: error fetching url %s: %v",r.Url, err)
		return ParseResult{}, err
	}
	return r.ParseFunc(body), nil
}
