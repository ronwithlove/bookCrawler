package main

import (
	"github.com/crawler/LearnGo-crawl/engine"
	"github.com/crawler/LearnGo-crawl/parse"
)

func main() {
	engine.Run(engine.Request{
		Url:"http://www.huhudm.com",
		ParseFunc:parse.ParseContent,
	})

	//engine.Run(engine.Request{
	//	Url:"http://www.huhudm.com/huhu39236.html",
	//	ParseFunc:parse.ParseBookDetail,
	//})

	//body, err := fetcher.Fetch("http://www.huhudm.com/huhu39236.html")
	//if err != nil {
	//	log.Printf("Fetching Error")
	//}
	//fmt.Printf("%s",body)
}


