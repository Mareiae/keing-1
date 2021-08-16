package keing

import (
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
)

type K map[string]interface{}
type HandlerFunc func(*Context)

type ErrorMsg struct {
	Message string      `json:"msg"`
	Meta    interface{} `json:"meta"`
}

type RouterGroup struct {
	Handlers []HandlerFunc
	prefix   string
	parent   *RouterGroup
	engine   *Engine
}

func (rt *RouterGroup) Group(component string, handlers ...HandlerFunc) *RouterGroup {
	prefix := path.Join(rt.prefix, component)
	return &RouterGroup{
		Handlers: rt.combineHandlers(handlers),
		parent:   rt,
		prefix:   prefix,
		engine:   rt.engine,
	}
}

func (rt *RouterGroup) Use(middlewares ...HandlerFunc) {
	rt.Handlers = append(rt.Handlers, middlewares...)
}

func (rt *RouterGroup) createContext(w http.ResponseWriter, r *http.Request, params httprouter.Params, handlers []HandlerFunc) *Context {
	return &Context{
		Request:  r,
		Writer:   w,
		Params:   params,
		handlers: handlers,
		index:    -1,
	}
}
func (rt *RouterGroup) combineHandlers(handlers []HandlerFunc) []HandlerFunc {
	s := len(rt.Handlers) + len(handlers)
	hs := make([]HandlerFunc, 0, s)
	hs = append(hs, rt.Handlers...)
	hs = append(hs, handlers...)
	return hs
}

func (rt *RouterGroup) Handle(method, p string, handlers []HandlerFunc) {
	p = path.Join(rt.prefix, p)
	rt.engine.index++
	rt.engine.path.AddPath(method, p, rt.engine.index)
	handlers = rt.combineHandlers(handlers)
	rt.engine.router.Handle(method, p, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		rt.createContext(w, r, params, handlers).Next()
	})
}

func (rt *RouterGroup) GET(path string, handlers ...HandlerFunc) {
	rt.Handle("GET", path, handlers)
}
func (rt *RouterGroup) POST(path string, handlers ...HandlerFunc) {
	rt.Handle("POST", path, handlers)
}
func (rt *RouterGroup) DELETE(path string, handlers ...HandlerFunc) {
	rt.Handle("DELETE", path, handlers)
}

func (rt *RouterGroup) PATCH(path string, handlers ...HandlerFunc) {
	rt.Handle("PATCH", path, handlers)
}

func (rt *RouterGroup) PUT(path string, handlers ...HandlerFunc) {
	rt.Handle("PUT", path, handlers)
}
