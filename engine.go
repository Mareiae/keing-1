package keing

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Engine struct {
	*RouterGroup
	router        *httprouter.Router
	path          *Paths
	index         int
	HtmlTemplates *template.Template
}

func Init() *Engine {
	engine := New()
	engine.Use(Logger())
	engine.path = &Paths{}
	engine.path.maxNum = 0
	engine.path.paths = make(map[int]*Path)
	return engine
}

func New() *Engine {
	engine := &Engine{}
	engine.RouterGroup = &RouterGroup{nil, "", nil, engine}
	engine.router = httprouter.New()
	engine.index = 0
	return engine
}

//使用ServeHTTP使路由实现 http.Handler的接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	engine.router.ServeHTTP(w, req)
}

func (engine *Engine) Run(addr string) {
	engine.path.ShowAllPathString()
	fmt.Printf("[Keing] Listening and serving HTTP on %s\n", addr)
	err := http.ListenAndServe(addr, engine)
	if err != nil {
		panic(err)
	}
}

func (engine *Engine) SetHtmlTemplates(pattren string) {
	engine.HtmlTemplates = template.Must(template.ParseGlob(pattren))
}
