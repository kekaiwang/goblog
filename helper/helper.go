package helper

import "github.com/deepzz0/goblog/RS"

type Response struct {
	Status int
	Data interface{}
	Err error
}

type Success struct {
	Level string
	Msg string
}

type Error struct {
	Level string
	Msg string
}

func NewResponse() *Response {
	return &Response{Status: RS.RS_success}
}
