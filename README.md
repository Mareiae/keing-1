<h1 align="center">Keing Web Framework</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-v0.1.1-blue.svg?cacheSeconds=2592000" />
  <a href="https://github.com/hXoreyer/keing/blob/master/LICENSE" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>
[![Goproxy.cn](https://goproxy.cn/stats/github.com/hxoreyer/keing/badges/download-count.svg)](https://goproxy.cn)
[![Go Report Card](https://goreportcard.com/badge/github.com/hxoreyer/keing)](https://goreportcard.com/report/github.com/hxoreyer/keing)
[![GoDoc](https://pkg.go.dev/badge/github.com/hxoreyer/keing?status.svg)](https://pkg.go.dev/github.com/hxoreyer/keing?tab=doc)

---

### 🏠 [Homepage](https://github.com/hxoreyer/keing)

### ✨ [Example](https://github.com/hXoreyer/keing/tree/master/example)

## 👨‍🎓 Author

👤 **hxoreyer**

* Github: [@hxoreyer](https://github.com/hxoreyer)

## 🎁 Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2021 [hxoreyer](https://github.com/hxoreyer).

This project is [MIT](https://github.com/hXoreyer/keing/LICENSE) licensed.   

---
## 使用✔
当然，提前得先要有go和git  

下载  
```shell
go get -u github.com/hxoreyer/keing
```

在你的go代码里import
```golang
import "github.com/hxoreyer/keing"
```

### 例子💯   
创建最简单的HTTP端点   
```golang   
pakage main

import "github.com/hxoreyer/keing"

func main(){
	k := keing.Init()
	k.GET("/hello",func(c *keing.Context){
		c.String(200,"hello keing")
	})
	
	k.Run(":8080")
}
```
获取路径中的参数 
```golang   
pakage main

import "github.com/hxoreyer/keing"

func main(){
	k := keing.Init()
	k.GET("/username/:name",func(c *keing.Context){
		c.String(200,"hello " + c.Name("name"))
	})
	
	k.Run(":8080")
}
```
分组路由 
```golang   
pakage main

import "github.com/hxoreyer/keing"

func main(){
	k := keing.Init()
	
	v1 := k.Group("/v1)
	{
		v1.POST("login",LoginHandler)
		v1.GET("read",ReadHandler)
	}
	
	v2 := k.Group("/v2")
	{
		v2.POST("login",LoginHandler)
		v2.GET("read",ReadHandler)
	}
	
	k.Run(":8080")
}
```
### 无默认中间件初始化🚫
使用
```golang   
k := keig.New()
```
代替
```golang
k := keing.Init()
```
### 中间件使用🗝
```golang   
k.Use(Logger())
```

