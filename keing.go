package keing

import (
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
)

//K 方便JSON格式数据发送
type K map[string]interface{}

//HandlerFunc 路由函数类型
type HandlerFunc func(*Context)

//ErrorMsg 错误信息结构
type ErrorMsg struct {
	Message string      `json:"msg"`
	Meta    interface{} `json:"meta"`
}

//RouterGroup 路由组结构
type RouterGroup struct {
	Handlers []HandlerFunc
	prefix   string
	parent   *RouterGroup
	engine   *Engine
}

//Group 创建一个路由组，并初始化
func (rt *RouterGroup) Group(component string, handlers ...HandlerFunc) *RouterGroup {
	prefix := path.Join(rt.prefix, component)
	return &RouterGroup{
		Handlers: rt.combineHandlers(handlers),
		parent:   rt,
		prefix:   prefix,
		engine:   rt.engine,
	}
}

//Use 使用中间件
func (rt *RouterGroup) Use(middlewares ...HandlerFunc) {
	rt.Handlers = append(rt.Handlers, middlewares...)
}

//创建响应上下文
func (rt *RouterGroup) createContext(w http.ResponseWriter, r *http.Request, params httprouter.Params, handlers []HandlerFunc) *Context {
	return &Context{
		Request:  r,
		Writer:   w,
		Params:   params,
		handlers: handlers,
		index:    -1,
		engine:   rt.engine,
	}
}

//路由函数操作
func (rt *RouterGroup) combineHandlers(handlers []HandlerFunc) []HandlerFunc {
	s := len(rt.Handlers) + len(handlers)
	hs := make([]HandlerFunc, 0, s)
	hs = append(hs, rt.Handlers...)
	hs = append(hs, handlers...)
	return hs
}

//Handle 给路由添加路由函数
func (rt *RouterGroup) Handle(method, p string, handlers []HandlerFunc) {
	p = path.Join(rt.prefix, p)
	rt.engine.index++
	rt.engine.path.AddPath(method, p, rt.engine.index)
	handlers = rt.combineHandlers(handlers)
	rt.engine.router.Handle(method, p, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		rt.createContext(w, r, params, handlers).Next()
	})
}

//GET GET请求
func (rt *RouterGroup) GET(path string, handlers ...HandlerFunc) {
	rt.Handle("GET", path, handlers)
}

//POST POST请求
func (rt *RouterGroup) POST(path string, handlers ...HandlerFunc) {
	rt.Handle("POST", path, handlers)
}

//DELETE DELETE请求
func (rt *RouterGroup) DELETE(path string, handlers ...HandlerFunc) {
	rt.Handle("DELETE", path, handlers)
}

//PATCH PATCH请求
func (rt *RouterGroup) PATCH(path string, handlers ...HandlerFunc) {
	rt.Handle("PATCH", path, handlers)
}

//PUT PUT请求
func (rt *RouterGroup) PUT(path string, handlers ...HandlerFunc) {
	rt.Handle("PUT", path, handlers)
}
