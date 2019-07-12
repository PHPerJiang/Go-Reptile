package scheduler

import "engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

//向请求通道发送请求
func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

//向工作通道发送准备好的worker
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

//单纯的为了实现接口
func (s *QueuedScheduler) ConfigureMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}

//启动
func (s *QueuedScheduler) Run() {
	//创建请求、工作通道
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	//并发
	go func() {
		//创建请求、工作队列
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		//循环调度
		for {
			//声明当前请求及work
			var activeRequest engine.Request
			var acticeWorker chan engine.Request
			//如果队列都不为空，说明至少存在一个请求及一个准备好的work
			if len(requestQ) > 0 && len(workerQ) > 0 {
				acticeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			//通信选择器
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case acticeWorker <- activeRequest:
				//若进请求进入所选的工作通信成功，则抛出当前选线
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
