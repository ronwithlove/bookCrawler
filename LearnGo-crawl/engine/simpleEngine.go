package engine

import (
	"fmt"
	"github.com/crawler/LearnGo-crawl/fetcher"
	"log"
)

type SimpleEngine struct{

}
func (s *SimpleEngine)Run(seeds ...Request) {
	var requests []Request

	for _, s := range seeds {
		requests = append(requests, s) //接成一个 slice
	}

	for len(requests) > 0 {
		r := requests[0]                      //把request中第一个元素拿出来
		requests = requests[1:]               //从slice把第一个原属剔除
		log.Printf("Fetching url: %s", r.Url) //得到网址打印出来
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetching Error: %s", r.Url)
		}
		//fmt.Printf("%s",body)
		parseResult := r.ParseFunc(body)                     //得到的结果用解析器解析
		requests = append(requests, parseResult.Requests...) //解析后得到的结果再继续放到requests中等待解析

		for _, item := range parseResult.Items { //Item打印出来看下
			fmt.Printf("Got item:%s\n", item)
		}
	}

}
