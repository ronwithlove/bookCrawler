package model

import "strconv"

type BookDetail struct{
	BookName string
	Author string
	Status string
	Episode int
	Updte string
	Collected int
	Introduction string
}

//重载String方法，当BookDetail被当string打印的时候就会根据这个方法来打印结果。
func (b BookDetail) String() string{
	return  "书名: "+ b.BookName+",作者: "+ b.Author+", 状态: "+b.Status+",集数: "+strconv.Itoa(b.Episode)+",更新: "+b.Updte+",收藏: "+strconv.Itoa(b.Collected)+",简介: "+b.Introduction
}