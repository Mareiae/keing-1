package keing

import (
	"encoding/json"
	"encoding/xml"
	"math"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const abort = math.MaxInt8 / 2

type Context struct {
	Request  *http.Request
	Writer   http.ResponseWriter
	Params   httprouter.Params
	handlers []HandlerFunc
	Errors   []ErrorMsg
	index    int8
	code     int
	engine   *Engine
}

func (c *Context) Next() {
	c.index++
	s := int8(len(c.handlers))
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

//终止处理
//强制系统不调用挂起的程序
func (c *Context) Abort(code int) {
	c.Writer.WriteHeader(code)
	c.index = abort
}

//将错误附加到当前上下文
//将错误推送至错误列表
//中间件会从错误列表获取错误信息
func (c *Context) Error(err error, meta interface{}) {
	c.Errors = append(c.Errors, ErrorMsg{
		Message: err.Error(),
		Meta:    meta,
	})
}

//以上的结合
func (c *Context) Fail(code int, err error) {
	c.Error(err, "Operation aborted")
	c.Abort(code)
}

/* 数据发送块*/
//将字符串写至响应正文中
//更改Content-Type
func (c *Context) String(code int, msg string) {
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.WriteHeader(code)
	c.code = code
	c.Writer.Write([]byte(msg))
}

//将给定的结构转化为Json至响应正文中
func (c *Context) Json(code int, obj interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	if code >= 0 {
		c.Writer.WriteHeader(code)
		c.code = code
	}
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		c.Error(err, obj)
		http.Error(c.Writer, err.Error(), 500)
		c.code = 500
	}
}

//将字节组写至响应正文中
func (c *Context) Data(code int, data []byte) {
	c.Writer.WriteHeader(code)
	c.code = code
	c.Writer.Write(data)
}

//将给定的结构转化为XML至响应正文中
func (c *Context) XML(code int, obj interface{}) {
	c.Writer.Header().Set("Content-Type", "application.xml")
	if code >= 0 {
		c.Writer.WriteHeader(code)
	}
	encoder := xml.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		c.Error(err, obj)
		http.Error(c.Writer, err.Error(), 500)
		c.code = 500
	}
}

func (c *Context) HTML(code int, name string, data interface{}) {
	c.Writer.Header().Set("Content-Type", "text/html")
	if code >= 0 {
		c.Writer.WriteHeader(code)
	}
	if err := c.engine.HtmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Error(err, map[string]interface{}{
			"name": name,
			"data": data,
		})
		http.Error(c.Writer, err.Error(), 500)
		c.code = 500
	}
}

/*数据接受块*/
//将正文内容转化为json
func (c *Context) BindJson(js interface{}) error {
	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(&js); err == nil {
		return nil
	} else {
		return err
	}
}

//表单数据获取
func (c *Context) Form(key string) string {
	return c.Request.FormValue(key)
}

//键值获取
func (c *Context) Query(key string) interface{} {
	r := make(map[string]interface{})
	_ = c.BindJson(&r)
	return r[key]
}

//键值获取(转化为相应数据格式)
func (c *Context) QueryInt64(key string) int64 {
	return int64(c.Query(key).(float64))
}
func (c *Context) QueryInt32(key string) int32 {
	return int32(c.Query(key).(float64))
}
func (c *Context) QueryInt(key string) int {
	return int(c.Query(key).(float64))
}
func (c *Context) QueryFloat64(key string) float64 {
	return c.Query(key).(float64)
}
func (c *Context) QueryString(key string) string {
	return c.Query(key).(string)
}

//Name路由值获取
//例如: /hello/:name  request: localhost:8080/hello/keing
//c.Name("name")将获取keing
func (c *Context) Name(key string) string {
	return c.Params.ByName(key)
}
