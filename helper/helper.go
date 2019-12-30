package helper

import (
	"encoding/json"
	"fmt"
	"github.com/deepzz0/goblog/RS"
	"net/http"
)

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

func (resp *Response) WriteJson(w http.ResponseWriter) {
	b, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("helper.go line:33", err)
		w.Write([]byte(`{Status:-1,Err:Error{Level:"alert",Msg:"code=-1|序列化失败！"}}`))
	} else {
		w.Write(b)
	}
}
