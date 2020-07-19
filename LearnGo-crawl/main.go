package main

import (
	"github.com/crawler/LearnGo-crawl/engine"
	"github.com/crawler/LearnGo-crawl/parse"
)

func main() {
	e:=engine.ConcurrentEngine{
		&engine.SimpleScheduler{},
		10,
	}
	e.Run(engine.Request{
		Url:"http://www.huhudm.com",
		ParseFunc:parse.ParseContent,
	})

}


