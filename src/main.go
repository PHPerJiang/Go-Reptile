package main

import (
	"engine"
	"scheduler"
	"zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		//开启n个协程
		WorkerCount:100,
	}

	//e.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//})

	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun/shanghai",
		ParseFunc:parser.ParseCity,
	})
}
