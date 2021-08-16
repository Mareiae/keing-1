package keing

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"math"
	"net/http"
)

const abort = math.MaxInt8 / 2

type Context struct {
	Request 	*http.Request
	Writer	 	http.ResponseWriter
	Params 		httprouter.Params
	handlers 	[]HandlerFunc
	Errors		[]ErrorMsg
	index 		int8
	code 		int
}
func (c *Context)Next(){
	c.index++
	s := int8(len(c.handlers))
	for ; c.index < s; c.index++{
		c.handlers[c.index](c)
	}
}
func (c *Context)Abort(code int){
	c.Writer.WriteHeader(code)
	c.index = abort
}
func (c *Context)String(code int,msg string){
	c.Writer.Header().Set("Content-Type","text/plain")
	c.Writer.WriteHeader(code)
	c.code = code
	c.Writer.Write([]byte(msg))
}
func (c *Context)Json(code int,obj interface{}){
	c.Writer.Header().Set("Content-Type","application/json")
	if code >= 0{
		c.Writer.WriteHeader(code)
		c.code = code
	}
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj);err != nil{
		c.Error(err,obj)
		http.Error(c.Writer,err.Error(),500)
		c.code = 500
	}
}
func (c *Context)Error(err error,meta interface{}){
	c.Errors = append(c.Errors,ErrorMsg{
		Message: err.Error(),
		Meta:    meta,
	})
}
func (c *Context)Data(code int,data []byte){
	c.Writer.WriteHeader(code)
	c.code = code
	c.Writer.Write(data)
}
func (c *Context)BindJson(js interface{}) error{
	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(&js);err == nil{
		return nil
	}else {
		return err
	}
}
func (c *Context)Form(key string)string{
	return c.Request.FormValue(key)
}
func (c *Context)Query(key string)interface{}{
	r := make(map[string]interface{})
	_ = c.BindJson(&r)
	return r[key]
}
func (c *Context)QueryInt64(key string)int64{
	return int64(c.Query(key).(float64))
}
func (c *Context)QueryInt32(key string)int32{
	return int32(c.Query(key).(float64))
}
func (c *Context)QueryInt(key string)int{
	return int(c.Query(key).(float64))
}
func (c *Context)QueryFloat64(key string)float64{
	return c.Query(key).(float64)
}
func (c *Context)QueryString(key string)string{
	return c.Query(key).(string)
}
func (c *Context)Name(key string)string{
	return c.Params.ByName(key)
}