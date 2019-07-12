package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

//定义调度器接口
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	//创建worker
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}

	//解析参数
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	//统计
	itemOut := 0
	//读取请求并打印数据
	for {
		result := <-out
		//打印获取的数据
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v\n",itemOut, item)
			itemOut++
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

//创建worker
func createWorker(out chan ParseResult, s Scheduler) {
	//并发
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			//从管道中取出一个请求
			request := <-in
			//交给worker进行处理
			result, err := worker(request)
			if err != nil{
				continue
			}
			out <- result
		}
	}()
}
