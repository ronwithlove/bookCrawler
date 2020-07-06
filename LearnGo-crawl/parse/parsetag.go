package parse

import (
	"github.com/crawler/LearnGo-crawl/engine"
	"regexp"
)

const regexpStr=`<span><a href='/comic/class_([0-9]*).html'>([\S]+)</a></span>`
func ParseContent(content []byte) engine.ParseResult {
	re:=regexp.MustCompile(regexpStr)
	match:=re.FindAllSubmatch(content,-1)//-1是全部都要
	result:= engine.ParseResult{}
	for _,m:=range match{
		result.Items=append(result.Items,m[2])
		result.Requests=append(result.Requests,engine.Request{
			Url:"http://www.huhudm.com/comic/class_"+string(m[1])+".html",
			ParseFunc:ParseBookList,//下一层用的解析器，目前没有，就这样写，不可以直接写nil
			//ParseFunc:nil,//这样写不符合规则，会报错
		})
	}
	return result
}

