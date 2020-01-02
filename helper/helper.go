package helper

import (
	"encoding/json"
	"fmt"
	"github.com/wkekai/goblog/RS"
	"net/http"
)

const (
	WARNING = "warning"
	SUCCESS = "success"
	ALERT 	= "alert"
	INFO	= "info"
)

type Response struct {
	Status int
	Data interface{}
	Err Error
}

type Success struct {
	Level string
	Message string
}

type Error struct {
	Level string
	Message string
}

func NewResponse() *Response {
	return &Response{Status: http.StatusOK}
}

func (resp *Response) Tips(level string, rs int) {
	resp.Err = Error{level, "code=" + fmt.Sprint(rs) + "|" + RS.Desc(rs)}
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
