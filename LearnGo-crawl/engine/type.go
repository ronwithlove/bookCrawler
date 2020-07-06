package engine

type ParseResult struct {
	Requests []Request //继续执行的request
	Items []interface{}//这个reuslt里获取的一些item,类型可能是任意的，所以用接口
}

type Request struct{
	Url string
	ParseFunc func([]byte) ParseResult //对上面这个url进行解析的方法,返回一个解析结果
}

//签名和Request结构体中的ParseFunc保持一直
func NilParse([]byte) ParseResult  {
	return  ParseResult{}
}