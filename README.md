## Goblog

Goblog 是基于beego框架开发的博客系统。简约、简单、轻量。数据库使用Mysql、图片静态文件使用七牛存储，Vue后台。

#### 关于版本
``` sh
Go    :go version go1.13.5 darwin/amd64
beego :1.12.0
bee   :v1.10.0
```

#### 获取项目 
```
go get -u github.com/wkekai/goblog
```

##### 修改配置
所有配置文件均在项目目录下的<code>conf</code>下。  
1. <code>app.conf</code>，这里是beego框架的配置文件。beego框架地址:[http://beego.me](http://beego.me) .
runmode选择你要运行的模式，对应下面的dev、prod、test。线上模式一般为prod，因为该模式不会输出beego的调试log。
``` ini
appname = goblog
runmode = dev 

// mysql配置
mysqlurls = "localhost"
mysqldb = "goblog"
mysqlport = "3306"
mysqluser = "root"
mysqlpass = "123456"
mysqlprefix = "wkk_"

// 七牛配置
AK = "*****"
SK = "*****"

[dev]
httpport = 8080
[prod]
enablehttp = true
httpport = 80
# 用于生成feed等其它需要用到域名的地方
mydomain = wangkekai.com
[test]
httpport = 8888
```

2.copy运行根目录下的goblog.sql文件并配置对应的数据库信息，程序就跑起来了

3.管理端目前单独在一个项目后期会整合进来，当前sql文件已更新。

>配置好数据库之后是可以直接跑起来的

#### 基本页面 
可以到我的博客[https://kekai.wang](https://kekai.wang)查看
