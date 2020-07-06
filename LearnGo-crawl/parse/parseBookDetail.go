package parse

import (
	"github.com/crawler/LearnGo-crawl/engine"
	"github.com/crawler/LearnGo-crawl/model"
	"log"
	"regexp"
	"strconv"
)

var authorReg=regexp.MustCompile(`<li><b>作者:</b>([^<]+?)</li>`)
var statusReg=regexp.MustCompile(`<li><b>状态:</b>([^<]+?)</li>`)
var episodeReg=regexp.MustCompile(`<li><b>集数:</b>([0-9]*)`)
var updteReg=regexp.MustCompile(`<li><b>更新:</b>([^<]+?)</li>`)
var collectedReg=regexp.MustCompile(`<li><b>收藏:</b>([0-9]*)`)
var introReg=regexp.MustCompile(`简介</b>:([^<]+?)</li>`)

func ParseBookDetail(content []byte) engine.ParseResult {
	bookDetail:=model.BookDetail{}
	bookDetail.Author=ExtraString(content,authorReg)
	bookDetail.Status=ExtraString(content,statusReg)
	episode, err := strconv.Atoi(ExtraString(content, episodeReg))
	if err!=nil{
		log.Printf("episode error: %s",err)
	}
	bookDetail.Episode=episode
	bookDetail.Updte=ExtraString(content,updteReg)
	collected, err := strconv.Atoi(ExtraString(content, collectedReg))
	if err!=nil{
		log.Printf("collected error: %s",err)
	}
	bookDetail.Collected=collected
	bookDetail.Introduction=ExtraString(content,introReg)

	result:= engine.ParseResult{
		Items:[]interface{}{bookDetail},
	}

	return result
}

func ExtraString(content []byte, re *regexp.Regexp) string{
	match:=re.FindSubmatch(content)
	if len(match)>=2{//至少2个，本身全部的一个[0]，和()中的一个[1]
		return  string(match[1])
	}
	return ""
}