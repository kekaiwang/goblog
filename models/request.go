package models

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"time"
)

type Request struct {
	Id 			int
	Refer    	string              	// 请求来源
	Url        	string              	// 访问页面
	Ip 			string              	// 请求IP
	Proxy	   	string					// 代理地址
	SessionId  	string              	// 请求session
	UserAgent  	string 					// UserAgent
	Created     time.Time           	// 请求时间
}

func NewRequest(r *context.BeegoInput) *Request {
	request := &Request{Created: time.Now()}
	request.Refer = r.Referer()
	request.Url = r.URL()
	request.Ip = r.IP()
	request.Proxy = r.Proxy()[0]
	request.SessionId = r.Cookie("SESSIONWKK")
	request.UserAgent = r.UserAgent()

	return request
}

type RequestManage struct {
	Ch chan *Request
}

var RequestM = NewRequestM()

func NewRequestM() *RequestManage {
	return &RequestManage{Ch: make(chan *Request, 20)}
}

func (m *RequestManage) SaveRequest() {
	for {
		select {
		case request := <-m.Ch:
			//var res Request
			//res := request
			or := orm.NewOrm()
			_, err := or.Insert(request)
			if err != nil {
				fmt.Println("save err " , request)
			}
		}
	}
}
