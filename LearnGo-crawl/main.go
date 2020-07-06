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
}


