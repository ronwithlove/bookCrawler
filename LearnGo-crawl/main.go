package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.99comic.com")
	if err!=nil{
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode!=http.StatusOK{
		fmt.Printf("Error status code:%d",resp.StatusCode)
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		panic(err)
	}

	//fmt.Printf("%s",result)
	parseContent(result)
}


func parseContent(content []byte){
	//str:=`<a href="/comic/class_2.html">搞笑</a>`
	re:=regexp.MustCompile(`<span><a href='/comic/class_([0-9]*).html'>([^"]+?)</a></span>`)
	match:=re.FindAllSubmatch(content,-1)//-1是全部都要
	fmt.Printf("%s",match)
	for _,m:=range match{
		fmt.Printf("m[0]:%s,m[1]:%s,m[2]:%s\n",m[0],m[1],m[2])
	}

}