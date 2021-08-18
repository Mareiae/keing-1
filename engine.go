package keing

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Engine 路由核心
type Engine struct {
	*RouterGroup
	router        *httprouter.Router
	path          *Paths
	index         int
	HTMLTemplates *template.Template
}

//Init 路由初始化(带默认中间件)
func Init() *Engine {
	engine := New()
	engine.Use(Logger())
	engine.path = &Paths{}
	engine.path.maxNum = 0
	engine.path.paths = make(map[int]*Path)
	return engine
}

//New 路由初始化(无默认中间件)
func New() *Engine {
	engine := &Engine{}
	engine.RouterGroup = &RouterGroup{nil, "", nil, engine}
	engine.router = httprouter.New()
	engine.index = 0
	return engine
}

//ServeHTTP 使用ServeHTTP使路由实现 http.Handler的接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	engine.router.ServeHTTP(w, req)
}

//Run 通过addr运行
func (engine *Engine) Run(addr string) {
	engine.path.ShowAllPathString()
	fmt.Printf("[Keing] Listening and serving HTTP on %s\n", addr)
	err := http.ListenAndServe(addr, engine)
	if err != nil {
		panic(err)
	}
}

//SetHTMLTemplates 确保HTML模板无错误，再将其赋予变量中
func (engine *Engine) SetHTMLTemplates(pattren string) {
	engine.HTMLTemplates = template.Must(template.ParseGlob(pattren))
}
