package _struct

type Resp struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}