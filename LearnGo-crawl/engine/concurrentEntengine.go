package engine

import (
	"github.com/crawler/LearnGo-crawl/fetcher"
	"log"
)

//调度器接口，可以有很多接口，负载均衡也可以作为调度器的一个功能
type Scheduler interface {
	Submit(Request)
	configureWorkChan(chan Request)//配置通道
}

//调度器实现
type SimpleScheduler struct {
	workerChan chan Request //简单版本就放一个任务通道
}

func (s *SimpleScheduler) Submit(r Request) {
	//使用go程，避免锁死
	go func() {s.workerChan<-r}()//就一个功能,把任务chan 给到workerChan
}

func (s *SimpleScheduler) configureWorkChan(c chan Request) {
	s.workerChan=c//有通道进来就复制给workerChan
}

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)

	e.Scheduler.configureWorkChan(in)//配置这个in,这个时候in就是workerChan，workerChan就是in
	for i := 0; i < e.WorkCount; i++ {
		CreateWork(in, out)//创建工人
	}
	//把request 发到 调度器中的workerChan里
	 for _,r:=range seeds{
	 	e.Scheduler.Submit(r)
	 }

	 //处理out
	itemcount:=0
	for{//这个主线程，不断的把worker处理出的request再提交到workChan里
		result:=<-out
		//处理Item，打印出来
		for _,item:=range result.Items{
			log.Printf("Got item %d:%s\n",itemcount,item)
			itemcount++
		}
		//处理requests,再放入workerChan里，继续去执行
		for _,result:=range result.Requests{
			e.Scheduler.Submit(result)
		}
	}
}

func CreateWork(in chan Request, out chan ParseResult) {
	go func() {
		for { //这里是go程，这里有很多个这样的worker，再不停的工作：等待request，拿到了，就开始工作
			request := <-in //从workerChan拿任务，会通过Submit把任务传进去，传进去了，这里就可以拿到了
			result, err := worker(request) //处理得到结果
			if err != nil { //如果有错误，忽略
				continue
			}
			out <- result //结果给到out
		}
	}()
}

func worker(r Request)(ParseResult,error){
	log.Printf("Fetching url: %s", r.Url) //得到网址打印出来
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetching Error: %s", r.Url)
	}
	return  r.ParseFunc(body),nil
}