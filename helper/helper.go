package helper

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

// -------------------- response ----------------------
const (
	WARNING = "warning"
	SUCCESS = "success"
	ALERT 	= "alert"
	INFO	= "info"
	ERROR 	= "error"
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

func (resp *Response) Tips(level string, rs string) {
	resp.Err = Error{level, rs}
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

// -------------------- password ----------------------
const (
	SALT = "#,.><)(_+-w*k$^*&"
)

func Md5(name, password, salt string) string {
	saltString := "!w@k#k"
	pass := md5.New()
	io.WriteString(pass, saltString)
	io.WriteString(pass, name)
	io.WriteString(pass, salt)
	io.WriteString(pass, password)

	return fmt.Sprintf("%x", pass.Sum(nil))
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
