package keing

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type K map[string]interface{}
type HandlerFunc func(*Context)

type ErrorMsg struct {
	Message string 		`json:"msg"`
	Meta	interface{} `json:"meta"`
}


type Engine struct {
	Handlers	[]HandlerFunc
	router 		*httprouter.Router
	path 		*Paths
	index 		int
}

func New() *Engine{
	engine := &Engine{}
	engine.router = httprouter.New()
	engine.index = 0
	return engine
}

func Init() *Engine{
	engine := New()
	engine.Use(Logger())
	engine.path = &Paths{}
	engine.path.maxNum = 0
	engine.path.paths = make(map[int]*Path)
	return engine
}


func (engine *Engine)Use(middlewares ...HandlerFunc){
	engine.Handlers = append(engine.Handlers,middlewares...)
}
func (engine *Engine)Run(addr string){
	engine.path.ShowAllPathString()
	fmt.Printf("[Keing] Listening and serving HTTP on %s\n",addr)
	err := http.ListenAndServe(addr, engine.router)
	if err != nil {
		panic(err)
		return
	}
}

func (engine *Engine)createContext(w http.ResponseWriter,r *http.Request,params httprouter.Params,handlers []HandlerFunc)*Context{
	return &Context{
		Request: r,
		Writer:  w,
		Params: params,
		handlers: handlers,
		index: -1,
	}
}
func (engine *Engine)combineHandlers(handlers []HandlerFunc)[]HandlerFunc{
	s := len(engine.Handlers) + len(handlers)
	hs := make([]HandlerFunc,0,s)
	hs = append(hs,engine.Handlers...)
	hs = append(hs,handlers...)
	return hs
}

func (engine *Engine)Handle(method,p string,handlers []HandlerFunc){
	engine.index++
	engine.path.AddPath(method,p,engine.index)
	handlers = engine.combineHandlers(handlers)
	engine.router.Handle(method,p, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		engine.createContext(w,r,params,handlers).Next()
	})
}

func (engine *Engine)GET(path string,handlers ...HandlerFunc){
	engine.Handle("GET",path,handlers)
}
func (engine *Engine)POST(path string,handlers ...HandlerFunc){
	engine.Handle("POST",path,handlers)
}
func (engine *Engine)DELETE(path string,handlers ...HandlerFunc){
	engine.Handle("DELETE",path,handlers)
}

func (engine *Engine)PATCH(path string,handlers ...HandlerFunc){
	engine.Handle("PATCH",path,handlers)
}

func (engine *Engine)PUT(path string,handlers ...HandlerFunc){
	engine.Handle("PUT",path,handlers)
}

