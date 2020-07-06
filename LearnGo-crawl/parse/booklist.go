package parse

import (
	"github.com/crawler/LearnGo-crawl/engine"
	"regexp"
)

const bookListReg=`<a title='([\S]+)' href='/huhu([0-9]*).html'>`
func ParseBookList(content []byte) engine.ParseResult {
	re:=regexp.MustCompile(bookListReg)
	match:=re.FindAllSubmatch(content,-1)//-1是全部都要
	result:= engine.ParseResult{}
	for _,m:=range match{
		result.Items=append(result.Items,m[1])
		result.Requests=append(result.Requests,engine.Request{
			Url:"http://www.huhudm.com/huhu"+string(m[2])+".html",
			ParseFunc:engine.NilParse,//下一层用的解析器，目前没有，就这样写，不可以直接写nil
			//ParseFunc:nil,//这样写不符合规则，会报错
		})
	}
	return result
}
