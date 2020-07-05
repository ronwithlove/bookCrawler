package main

import (
	"fmt"
	"regexp"
)

func main() {
	str:="My email is shiyiming@outlook.com"
	//re:=regexp.MustCompile("My email is shiyiming@outlook.com")//My email is shiyiming@outlook.com
	//.表示除换行符之外的任意"一个"字符，只表示一个
	//re:=regexp.MustCompile(`.@outlook.com`)// g@outlook.com
	//re:=regexp.MustCompile(`..@outlook.com`)// ng@outlook.com
	//*用来备注他前面的符号，可以出现0次或者多次
	//.*所以连在一起用，就可以向前一直找，直到换行符位置：
	//这里用的符号不是单引号，是1前面的那个点
	re:=regexp.MustCompile(`.*@outlook.com`)// My email is shiyiming@outlook.com
	result:=re.FindString(str)
	fmt.Println(result)

	fmt.Println(
		//[0-9a-zA-Z]表示任意"一个"0-9或者a-z或者A-Z的字符
		regexp.MustCompile(`[0-9a-zA-Z]*@outlook.com`).FindString(str),//shiyiming@outlook.com
	)



}
